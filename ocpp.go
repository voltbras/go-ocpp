package ocpp

import (
	"github.com/eduhenke/go-ocpp/messages"
)

type MessageHandler func(request messages.Request, cpID string) (messages.Response, error)
type Version string

const (
	V15 Version = "ocpp1.5"
	V16 Version = "ocpp1.6"
)
