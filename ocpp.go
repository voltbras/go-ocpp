package ocpp

import (
	"github.com/eduhenke/go-ocpp/soap"
)

type Envelope soap.Envelope
type Version string

const (
	V15 Version = "ocpp1.5"
	V16 Version = "ocpp1.6"
)
