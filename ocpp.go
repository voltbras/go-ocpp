package ocpp

import (
	"github.com/voltbras/go-ocpp/internal/log"
	"github.com/voltbras/go-ocpp/messages"
)

type MessageHandler func(request messages.Request, cpID string, url string) (messages.Response, error)

type Version string

const (
	V15 Version = "V15"
	V16 Version = "V16"
)

type Transport string

const (
	SOAP Transport = "soap"
	JSON Transport = "json"
)

func SetErrorLogger(logger log.Logger) { log.SetErrorLogger(logger) }
func SetDebugLogger(logger log.Logger) { log.SetDebugLogger(logger) }
