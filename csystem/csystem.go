package csystem

import (
	"encoding/json"
	"errors"
	"github.com/eduhenke/go-ocpp/messages"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpres"
	"github.com/eduhenke/go-ocpp/soap"
	"github.com/eduhenke/go-ocpp/wsconn"
	ws "github.com/gorilla/websocket"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr, "xxx: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// ChargePointMessageHandler handles the OCPP messages coming from the charger
type ChargePointMessageHandler func(cprequest cpreq.ChargePointRequest, cpID string) (cpres.ChargePointResponse, error)

func convert(msg wsconn.Message) (cpreq.ChargePointRequest, wsconn.ErrorCode) {
	switch call := msg.(type) {
	case *wsconn.CallMessage:
		req := cpreq.GetRequestFromActionName(string(call.Action))
		if req == nil {
			return nil, wsconn.NotSupported
		}
		originalPayload, err := json.Marshal(call.Payload)
		if err != nil {
			return nil, wsconn.GenericError
		}
		err = json.Unmarshal(originalPayload, req)
		if err != nil {
			return nil, wsconn.FormationViolation
		}
		return req, wsconn.Nil
	}
	return nil, wsconn.GenericError
}

func unconvert(id wsconn.MessageID, resp cpres.ChargePointResponse, err error) wsconn.Message {
	if err != nil {
		return wsconn.NewCallErrorMessage(id, wsconn.InternalError, err.Error())
	}
	return wsconn.NewCallResult(id, resp)
}

type CentralSystem struct {
	Conns map[string]*wsconn.Conn
}

func New() *CentralSystem {
	return &CentralSystem{
		Conns: make(map[string]*wsconn.Conn, 0),
	}
}

func (csys *CentralSystem) Run(port string, cphandler ChargePointMessageHandler) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cpID := strings.TrimPrefix(r.URL.Path, "/")
		if ws.IsWebSocketUpgrade(r) {
			conn, err := wsconn.New(w, r, []ocpp.Version{ocpp.V16})
			if err != nil {
				return
			}
			for {
				msg, err := conn.ReadMessage()
				if err != nil {
					_ = conn.Close()
					delete(csys.Conns, cpID)
					break
				}
				// logger.Printf("Received a message: %v\n", msg)
				id := msg.MessageID()
				req, wserr := convert(msg)
				// logger.Printf("Converted message to cpreq: %v\n", req)
				if wserr != wsconn.Nil {
					continue
				}
				csys.Conns[cpID] = conn
				logger.Printf("Sending cpreq to handler from cpID: %s\n", cpID)
				resp, err := cphandler(req, cpID)
				logger.Printf("Handler responded: %v\n", resp)
				msg = unconvert(id, resp, err)
				logger.Printf("Sending message [struct]: %v\n", msg)
				bts, err := json.Marshal(msg)
				if err != nil {
					continue
				}
				// logger.Printf("Sending message [raw]: %v\n", string(bts))
				err = conn.Conn.WriteMessage(ws.TextMessage, bts)
				if err != nil {
					continue
				}
				// logger.Println("Sent message!")
			}
		}

		// TODO: check whether JSON/SOAP and their versions
		soap.Handler(w, r, func (request messages.Request, cpID string) (messages.Response, error) {
			cprequest, ok := request.(cpreq.ChargePointRequest)
			if !ok {
				return nil, errors.New("request is not a cprequest")
			}
			return cphandler(cprequest, cpID)
		})
	})
	return http.ListenAndServe(port, nil)
}
