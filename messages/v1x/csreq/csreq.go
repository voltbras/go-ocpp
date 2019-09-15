package csreq

import (
	"github.com/eduhenke/go-ocpp/messages"
	"github.com/eduhenke/go-ocpp/messages/v1x/csres"
)

// CentralSystemRequest is a request coming from the central system to the chargepoint
type CentralSystemRequest interface {
	messages.Request
	IsCentralSystemRequest()
	GetCentralSystemResponse() csres.CentralSystemResponse
}

type centralSystemRequest struct{}

func (csreq *centralSystemRequest) IsCentralSystemRequest() {}
func (csreq *centralSystemRequest) IsRequest() {}
