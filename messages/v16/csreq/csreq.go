package csreq

import "github.com/eduhenke/go-ocpp/messages/v16/csres"

// CentralSystemRequest is a request coming from the central system to the chargepoint
type CentralSystemRequest interface {
	IsCentralSystemRequest()
	Action() string
	GetResponseObject() csres.CentralSystemResponse
}

type centralSystemRequest struct{}

func (csreq *centralSystemRequest) IsCentralSystemRequest() {}
