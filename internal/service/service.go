package service

import "github.com/eduhenke/go-ocpp/messages"

type Service interface {
	Send(request messages.Request) (messages.Response, error)
}
