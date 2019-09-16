package service

import (
	"github.com/eduhenke/go-ocpp/messages"
	"github.com/eduhenke/go-ocpp/ws"
)

type JsonService struct {
	conn *ws.Conn
}

func NewJsonService(conn *ws.Conn) *JsonService {
	return &JsonService{
		conn: conn,
	}
}

func (service *JsonService) Send(req messages.Request) (messages.Response, error) {
	return service.conn.SendRequest(req)
}
