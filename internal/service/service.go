package service

import "github.com/voltbras/go-ocpp/messages"

type Service interface {
	Send(request messages.Request) (messages.Response, error)
}
