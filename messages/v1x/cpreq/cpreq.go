package cpreq

import (
	"errors"

	"github.com/voltbras/go-ocpp/messages"
)

// ChargePointRequest is a request coming from the chargepoint to the central system
type ChargePointRequest interface {
	messages.Request
	IsChargePointRequest()
}

type chargepointRequest struct{}

func (cpreq *chargepointRequest) IsChargePointRequest() {}
func (cpreq *chargepointRequest) IsRequest()            {}

var (
	ErrorNotChargePointRequest = errors.New("not a chargepoint request")
)
