package cpres

// ChargePointResponse is a response coming from the central system to the chargepoint
type ChargePointResponse interface {
	IsChargePointResponse()
}

type chargepointResponse struct{}

func (cpreq *chargepointResponse) IsChargePointResponse() {}
