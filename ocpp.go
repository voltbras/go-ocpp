package ocpp

import (
	"github.com/eduhenke/go-ocpp/internal/log"
	"github.com/eduhenke/go-ocpp/messages"
)

type MessageHandler func(request messages.Request, cpID string) (messages.Response, error)

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
