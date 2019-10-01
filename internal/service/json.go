package service

import (
	"github.com/eduhenke/go-ocpp/messages"
	"github.com/eduhenke/go-ocpp/ws"
)

type JsonService struct {
	conn *ws.Conn
}

// NewJsonService before calling it, you should
// be reading messages from the connection
// as to get the responses back:
// go func() {
// 	for {
// 		conn.ReadMessage()
// 	}
// }()
func NewJsonService(conn *ws.Conn) *JsonService {
	return &JsonService{
		conn: conn,
	}
}

func (service *JsonService) Send(req messages.Request) (messages.Response, error) {
	return service.conn.SendRequest(req)
}
