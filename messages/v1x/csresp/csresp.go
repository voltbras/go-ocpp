package csresp

import (
	"errors"

	"github.com/voltbras/go-ocpp/messages"
)

// CentralSystemResponse is a response coming from the chargepoint to the central system
type CentralSystemResponse interface {
	messages.Response
	IsCentralSystemResponse()
}

type centralSystemResponse struct{}

func (csresp *centralSystemResponse) IsCentralSystemResponse() {}
func (csresp *centralSystemResponse) IsResponse()              {}

var (
	ErrorNotCentralSystemResponse = errors.New("not a central system response")
)
