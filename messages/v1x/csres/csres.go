package csres

import (
	"errors"
	"github.com/eduhenke/go-ocpp/messages"
)

// CentralSystemResponse is a response coming from the chargepoint to the central system
type CentralSystemResponse interface {
	messages.Response
	IsCentralSystemResponse()
}

type centralSystemResponse struct{}

func (csres *centralSystemResponse) IsCentralSystemResponse() {}
func (csres *centralSystemResponse) IsResponse() {}

var (
	ErrorNotCentralSystemResponse = errors.New("not a central system response")
)