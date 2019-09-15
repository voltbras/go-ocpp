package csres

// CentralSystemResponse is a response coming from the chargepoint to the central system
type CentralSystemResponse interface {
	IsCentralSystemResponse()
}

type centralSystemResponse struct{}

func (csres *centralSystemResponse) IsCentralSystemResponse() {}
