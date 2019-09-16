package csreq

import (
	"errors"
	"github.com/eduhenke/go-ocpp/messages"
)

// CentralSystemRequest is a request coming from the central system to the chargepoint
type CentralSystemRequest interface {
	messages.Request
	IsCentralSystemRequest()
}

type centralSystemRequest struct{}

func (csreq *centralSystemRequest) IsCentralSystemRequest() {}
func (csreq *centralSystemRequest) IsRequest()              {}

var (
	ErrorNotCentralSystemRequest = errors.New("not a central system request")
)
