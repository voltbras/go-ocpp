package cs

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
	"sync"

	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/internal/log"
	"github.com/eduhenke/go-ocpp/internal/service"
	"github.com/eduhenke/go-ocpp/messages"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpresp"
	"github.com/eduhenke/go-ocpp/soap"
	"github.com/eduhenke/go-ocpp/ws"
)

// ChargePointMessageHandler handles the OCPP messages coming from the charger
type ChargePointMessageHandler func(cprequest cpreq.ChargePointRequest, cpID string) (cpresp.ChargePointResponse, error)

type ChargePointConnectionListener func(cpID string)
type CentralSystem interface {
	// Run the central system on the given port
	// and handles each incoming ChargepointRequest
	Run(port string, cphandler ChargePointMessageHandler) error

	// GetServiceOf a chargepoint to enable
	// communication with the chargepoint
	//
	// the url parameter is *NOT* used if
	// the link between the CentralSystem
	// and Chargepoint is via Websocket
	GetServiceOf(cpID string, version ocpp.Version, url string) (service.ChargePoint, error)

	SetChargePointConnectionListener(ChargePointConnectionListener)
	SetChargePointDisconnectionListener(ChargePointConnectionListener)
	WaitConnect(cpID string) <-chan struct{}
	WaitDisconnect(cpID string) <-chan struct{}
}

type centralSystem struct {
	conns map[string]*ws.Conn
	// used to symbolize if the connection is connected
	connChans       map[string]chan struct{}
	connsConnected  map[string]bool
	connsCount      map[string]int
	connMux         sync.Mutex
	connListener    ChargePointConnectionListener
	disconnListener ChargePointConnectionListener
}

func New() CentralSystem {
	return &centralSystem{
		conns:           make(map[string]*ws.Conn, 0),
		connChans:       make(map[string]chan struct{}, 0),
		connsCount:      make(map[string]int, 0),
		connsConnected:  make(map[string]bool, 0),
		connListener:    func(cpID string) {},
		disconnListener: func(cpID string) {},
	}
}

func (csys *centralSystem) Run(port string, cphandler ChargePointMessageHandler) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if ws.IsWebSocketUpgrade(r) {
			csys.handleWebsocket(w, r, cphandler)
		} else if r.Method == http.MethodGet {
			// it's not a SOAP request
			// it's someone lurking around in this URL
			// let's present something nice
			body := fmt.Sprintf(
				`<h1>OCPP Central System</h1>
				<p>currently connected with %d OCPP-J stations, and more OCPP-S stations</p>
				<i>Central System using <a href="https://github.com/voltbras/go-ocpp"/>https://github.com/voltbras/go-ocpp</i>`,
				len(csys.conns),
			)
			w.Write([]byte(body))
		} else {
			csys.handleSoap(w, r, cphandler)
		}

	})
	log.Debug("Central system running on port: %s", port)
	return http.ListenAndServe(port, nil)
}

func (csys *centralSystem) handleWebsocket(w http.ResponseWriter, r *http.Request, cphandler ChargePointMessageHandler) {
	log.Debug("Current WS connections map: %v", csys.conns)
	cpID := strings.TrimPrefix(r.URL.Path, "/")

	rawReq, _ := httputil.DumpRequest(r, true)
	log.Debug("Raw WS request: %s", string(rawReq))

	conn, err := ws.Handshake(w, r, []ocpp.Version{ocpp.V16})
	if err != nil {
		log.Error("Couldn't handshake request %w", err)
		return
	}

	csys.connMux.Lock()
	csys.conns[cpID] = conn
	csys.connsCount[cpID]++
	if csys.connChans[cpID] == nil {
		csys.connChans[cpID] = make(chan struct{})
	}
	if !csys.connsConnected[cpID] {
		close(csys.connChans[cpID])
	}
	csys.connsConnected[cpID] = true
	csys.connMux.Unlock()

	log.Debug("Connected with %s", cpID)
	go csys.connListener(cpID)

	defer func() {
		conn.Close()
		log.Debug("Closed connection of: %s", cpID)
		go csys.disconnListener(cpID)
		csys.connMux.Lock()
		csys.connsCount[cpID]--
		// if the same CP connected more times before we do the
		// connection cleanup, don't remove the connection reference
		if csys.connsCount[cpID] == 0 {
			delete(csys.conns, cpID)
			csys.connChans[cpID] = make(chan struct{})
		}
		csys.connsConnected[cpID] = false
		csys.connMux.Unlock()
	}()

	for {
		select {
		case req := <-conn.Requests():
			cprequest, ok := req.Request.(cpreq.ChargePointRequest)
			if !ok {
				log.Error(cpreq.ErrorNotChargePointRequest.Error())
			}
			cpresponse, err := cphandler(cprequest, cpID)
			err = conn.SendResponse(req.MessageID, cpresponse, err)
			if err != nil {
				log.Error(err.Error())
			}
		case <-conn.WaitClose():
			return
		case err := <-conn.ReadMessageAsync():
			if err != nil {
				if !ws.IsCloseError(err) {
					continue
				}
				return
			}
		}
	}
}

func (csys *centralSystem) handleSoap(w http.ResponseWriter, r *http.Request, cphandler ChargePointMessageHandler) {
	log.Debug("New SOAP request")
	err := soap.Handle(w, r, func(request messages.Request, cpID string) (messages.Response, error) {
		req, ok := request.(cpreq.ChargePointRequest)
		if !ok {
			return nil, errors.New("request is not a cprequest")
		}
		return cphandler(req, cpID)
	})
	if err != nil {
		log.Error("Couldn't handle SOAP request: %w", err)
	}
}

func (csys *centralSystem) GetServiceOf(cpID string, version ocpp.Version, url string) (service.ChargePoint, error) {
	if version == ocpp.V15 {
		return service.NewChargePointSOAP(url, &soap.CallOptions{
			ChargeBoxIdentity: cpID,
			// TODO: insert IP address
			// From: <url>,
		}), nil
	}
	if version == ocpp.V16 {
		csys.connMux.Lock()
		conn := csys.conns[cpID]
		csys.connMux.Unlock()
		if conn == nil { // TODO: or conn closed
			return nil, errors.New("no connection to this charge point")
		}
		return service.NewChargePointJSON(conn), nil
	}
	return nil, errors.New("charge point has no configured OCPP version(1.5/1.6)")
}

func (csys *centralSystem) SetChargePointConnectionListener(f ChargePointConnectionListener) {
	csys.connListener = f
}

func (csys *centralSystem) SetChargePointDisconnectionListener(f ChargePointConnectionListener) {
	csys.disconnListener = f
}

func (csys *centralSystem) WaitDisconnect(cpID string) <-chan struct{} {
	csys.connMux.Lock()
	conn := csys.conns[cpID]
	csys.connMux.Unlock()

	ch := make(chan struct{})
	if conn == nil {
		close(ch)
		return ch
	}
	return conn.WaitClose()
}

func (csys *centralSystem) WaitConnect(cpID string) <-chan struct{} {
	csys.connMux.Lock()
	if csys.connChans[cpID] == nil {
		csys.connChans[cpID] = make(chan struct{})
	}
	connCh := csys.connChans[cpID]
	csys.connMux.Unlock()
	return connCh
}
