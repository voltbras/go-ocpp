package service

import (
	"github.com/voltbras/go-ocpp/messages"
	"github.com/voltbras/go-ocpp/ws"
)

type JSON struct {
	conn *ws.Conn
}

// NewJSON before calling it, you should
// be reading messages from the connection
// as to get the responses back:
// go func() {
// 	for {
// 		conn.ReadMessage()
// 	}
// }()
func NewJSON(conn *ws.Conn) *JSON {
	return &JSON{
		conn: conn,
	}
}

func (service *JSON) Send(req messages.Request) (messages.Response, error) {
	return service.conn.SendRequest(req)
}
