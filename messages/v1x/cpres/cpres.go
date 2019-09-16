package cpres

import (
	"errors"
	"github.com/eduhenke/go-ocpp/messages"
)

// ChargePointResponse is a response coming from the central system to the chargepoint
type ChargePointResponse interface {
	messages.Response
	IsChargePointResponse()
}

type chargepointResponse struct{}

func (cpreq *chargepointResponse) IsChargePointResponse() {}
func (cpreq *chargepointResponse) IsResponse() {}

var (
	ErrorNotChargePointResponse = errors.New("not a chargepoint response")
)