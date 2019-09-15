package service

import (
	"github.com/eduhenke/go-ocpp/messages"
	"github.com/eduhenke/go-ocpp/wsconn"
)

type JsonService struct {
	conn *wsconn.Conn
}

func NewJsonService(conn *wsconn.Conn) *JsonService {
	return &JsonService{
		conn: conn,
	}
}

func (service *JsonService) Send(req messages.Request) (messages.Response, error) {
	//service.conn.
	return nil, nil
}