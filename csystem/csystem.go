package csystem

import (
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

type CentralSystem struct {
	Conns map[string]*wsconn.Conn
	sentMsgs map[string]*wsconn.Message
}

func New() *CentralSystem {
	return &CentralSystem{
		Conns: make(map[string]*wsconn.Conn, 0),
		sentMsgs: make(map[string]*wsconn.Message, 0),
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
			go func() {
				for {
					err := conn.ReadMessage()
					if err != nil {
						log.Error("On receiving a message: %w\n", err)
						_ = conn.Close()
						delete(csys.Conns, cpID)
						break
					}
				}
			}()
			for {
				req := <-conn.Requests()
				cprequest, ok := req.Request.(cpreq.ChargePointRequest)
				if !ok {
					log.Error(cpreq.ErrorNotChargePointRequest.Error())
				}
				cpresponse, err := cphandler(cprequest, cpID)
				err = conn.SendResponse(req.MessageID, cpresponse, err)
				if err != nil {
					log.Error(err.Error())
				}
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
