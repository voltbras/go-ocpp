package csystem

import (
	"encoding/json"
	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/messages/v16/cpreq"
	"github.com/eduhenke/go-ocpp/messages/v16/cpres"
	"github.com/eduhenke/go-ocpp/wsconn"
	ws "github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr, "xxx: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// CentralSystemHandler handles the OCPP messages coming from the charger
type CentralSystemHandler func(cpreq.ChargePointRequest) (cpres.ChargePointResponse, error)

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

func ListenAndServe(port string, handler CentralSystemHandler) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if ws.IsWebSocketUpgrade(r) {
			conn, err := wsconn.New(w, r, []ocpp.Version{ocpp.V16})
			if err != nil {
				return
			}
			for {
				msg, err := conn.ReadMessage()
				if err != nil {
					panic(err)
				}
				logger.Printf("Received a message: %v\n", msg)
				id := msg.MessageID()
				req, wserr := convert(msg)
				logger.Printf("Converted message to cpreq: %v\n", req)
				if wserr != wsconn.Nil {
					panic(wserr)
				}
				logger.Printf("Sending cpreq to handler\n")
				resp, err := handler(req)
				logger.Printf("Handler responded: %v\n", resp)
				msg = unconvert(id, resp, err)
				logger.Printf("Sending message [struct]: %v\n", msg)
				bts, err := json.Marshal(msg)
				if err != nil {
					panic(err)
				}
				//bts := []byte(`[3,"`+string(msg.MessageID())+`",{"currentTime":"2013-02-01T20:53:32.486Z"}]`)
				logger.Printf("Sending message [raw]: %v\n", string(bts))
				err = conn.Conn.WriteMessage(ws.TextMessage, bts)
				if err != nil {
					panic(err)
				}
				logger.Println("Sent message!")
			}
		}

		// // TODO: check whether JSON/SOAP and their versions
		// soap.Handler(w, r, func(env soap.Envelope) (interface{}, error) {
		// 	return handler.HandleEnvelope((ocpp.Envelope)(env))
		// })
	})
	return http.ListenAndServe(port, nil)
}
