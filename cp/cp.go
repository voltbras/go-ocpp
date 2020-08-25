package cp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/internal/log"
	"github.com/eduhenke/go-ocpp/internal/service"
	"github.com/eduhenke/go-ocpp/messages"
	"github.com/eduhenke/go-ocpp/messages/v1x/csreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/csresp"
	"github.com/eduhenke/go-ocpp/soap"
	"github.com/eduhenke/go-ocpp/ws"
)

// CentralSystemMessageHandler handles the OCPP messages coming from the central system
type CentralSystemMessageHandler func(cprequest csreq.CentralSystemRequest) (csresp.CentralSystemResponse, error)

type ChargePoint interface {
	service.CentralSystem
	// Run the charge point on the given port
	// and handles each incoming CentralSystemRequest
	Run(port *string, cphandler CentralSystemMessageHandler) error
	Identity() string
}

const (
	websocketConnectionRetryInterval = 5 * time.Second
)

type chargePoint struct {
	service.CentralSystem
	identity         string
	centralSystemURL string
	version          ocpp.Version
	transport        ocpp.Transport
	conn             *ws.Conn
	ctx              context.Context
}

func New(ctx context.Context, identity, csURL string, version ocpp.Version, transport ocpp.Transport) (*chargePoint, error) {
	cp := &chargePoint{
		identity:         identity,
		centralSystemURL: csURL,
		version:          version,
		transport:        transport,
		ctx:              ctx,
	}
	var csService service.CentralSystem
	if transport == ocpp.JSON {
		err := cp.getNewWebsocketConnection()
		if err != nil {
			return nil, fmt.Errorf("could not dial to central system: %w", err)
		}
		go cp.handleWebsocketConnection()
		csService = service.NewCentralSystemJSON(cp.conn)
	}
	if transport == ocpp.SOAP {
		csService = service.NewCentralSystemSOAP(csURL, &soap.CallOptions{ChargeBoxIdentity: identity})
	}
	cp.CentralSystem = csService
	return cp, nil
}

func (cp *chargePoint) Identity() string {
	return cp.identity
}

func (cp *chargePoint) Run(port *string, cshandler CentralSystemMessageHandler) error {
	if cp.transport == ocpp.JSON {
		if cp.conn == nil {
			return errors.New("no ws connection")
		}
		go cp.handleWebsocket(cshandler)
	}
	if cp.transport == ocpp.SOAP {
		if port == nil {
			return errors.New("port must be set when running a SOAP ChargePoint")
		}
		return handleSoap(cp.ctx, *port, cshandler)
	}
	return nil
}

func (cp *chargePoint) getNewWebsocketConnection() error {
	conn, err := ws.Dial(cp.identity, cp.centralSystemURL, cp.version)
	if err != nil {
		return err
	}
	if cp.conn == nil {
		cp.conn = conn
	} else {
		*cp.conn = *conn
	}
	return nil
}

func (cp *chargePoint) handleFailedWebsocketConnection() {
	for {
		err := cp.getNewWebsocketConnection()
		if err != nil {
			log.Error("On restarting connection with Central System: %w", err)
			<-time.After(websocketConnectionRetryInterval)
		} else {
			break
		}
	}
}
func (cp *chargePoint) handleIncomingMessagesUntilConnectionFail() <-chan struct{} {
	connectionFailCh := make(chan struct{})
	go func() {
		for {
			err := cp.conn.ReadMessage()
			if err != nil {
				if !ws.IsNormalCloseError(err) {
					log.Error("On receiving a message: %w", err)
				}
				_ = cp.conn.Close()
				log.Debug("Closed connection of Central System")
				connectionFailCh <- struct{}{}
				break
			}
		}
	}()
	return connectionFailCh
}

func (cp *chargePoint) handleWebsocket(cshandler CentralSystemMessageHandler) {
	for {
		select {
		case <-cp.ctx.Done():
			return
		case req := <-cp.conn.Requests():
			cprequest, ok := req.Request.(csreq.CentralSystemRequest)
			if !ok {
				log.Error(csreq.ErrorNotCentralSystemRequest.Error())
			}
			cpresponse, err := cshandler(cprequest)
			err = cp.conn.SendResponse(req.MessageID, cpresponse, err)
			if err != nil {
				log.Error(err.Error())
			}
		}
	}
}

func (cp *chargePoint) handleWebsocketConnection() {
	connFailed := cp.handleIncomingMessagesUntilConnectionFail()
	for {
		select {
		case <-cp.ctx.Done():
			cp.conn.Close()
			return
		case <-connFailed:
			cp.handleFailedWebsocketConnection()
		}
	}
}

func handleSoap(ctx context.Context, port string, cshandler CentralSystemMessageHandler) error {
	handler := soapHandler{cshandler}
	server := &http.Server{Addr: port, Handler: handler}

	serverErrorCh := make(chan error)
	go func() {
		serverErrorCh <- server.ListenAndServe()
	}()
	select {
	case <-ctx.Done():
		return nil
	case err := <-serverErrorCh:
		return err
	}
}

type soapHandler struct{ cshandler CentralSystemMessageHandler }

func (s soapHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("New SOAP request")

	err := soap.Handle(w, r, func(request messages.Request, cpID string) (messages.Response, error) {
		req, ok := request.(csreq.CentralSystemRequest)
		if !ok {
			return nil, errors.New("request is not a cprequest")
		}
		return s.cshandler(req)
	})
	if err != nil {
		log.Error("Couldn't handle SOAP request: %w", err)
	}
}
