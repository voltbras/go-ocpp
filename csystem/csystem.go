package csystem

import (
	"encoding/json"
	"errors"
	"github.com/eduhenke/go-ocpp/internal/log"
	"net/http"
	"strings"

	"github.com/eduhenke/go-ocpp/messages"

	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpres"
	"github.com/eduhenke/go-ocpp/soap"
	"github.com/eduhenke/go-ocpp/wsconn"
	ws "github.com/gorilla/websocket"
)

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
			csys.Conns[cpID] = conn
			log.Debug("Accepted new WS conn: %v. from: %v\n", conn, cpID)
			for {
				msg, err := conn.ReadMessage()
				if err != nil {
					log.Error("On receiving a message: %w\n", err)
					_ = conn.Close()
					delete(csys.Conns, cpID)
					break
				}
				log.Debug("Received a message: %v\n", msg)
				id := msg.MessageID()
				req, wserr := convert(msg)
				log.Debug("Converted message to cpreq: %v\n", req)
				if wserr != wsconn.Nil {
					continue
				}
				log.Debug("Sending cpreq to handler from cpID: %s\n", cpID)
				resp, err := cphandler(req, cpID)
				log.Debug("Handler responded: %v\n", resp)
				msg = unconvert(id, resp, err)
				log.Debug("Sending message [struct]: %v\n", msg)
				bts, err := json.Marshal(msg)
				if err != nil {
					continue
				}
				log.Debug("Sending message [raw]: %v\n", string(bts))
				err = conn.Conn.WriteMessage(ws.TextMessage, bts)
				if err != nil {
					log.Error("On sending message %w\n", err)
					continue
				}
				log.Debug("Sent message!\n")
			}
		} else {
			log.Debug("New SOAP request\n")
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

	})
	log.Debug("Central system running on port: %s\n", port)
	return http.ListenAndServe(port, nil)
}
