package cp

import (
	"fmt"
	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/cs"
	"github.com/eduhenke/go-ocpp/soap"
	"github.com/eduhenke/go-ocpp/ws"
)

type ChargePoint struct {
	cs.Service
	identity string
	centralSystemURL string
	version ocpp.Version
	transport ocpp.Transport
}

func NewChargePoint(identity, csURL string, version ocpp.Version, transport ocpp.Transport) (*ChargePoint, error) {
	var service cs.Service
	if transport == ocpp.JSON {
		conn, err := ws.Dial(identity, csURL, version)
		if err != nil {
			return nil, fmt.Errorf("could not dial to central system %w", err)
		}
		service = cs.NewJsonService(conn)
	}
	if transport == ocpp.SOAP {
		service = cs.NewSoapService(csURL, &soap.CallOptions{ChargeBoxIdentity: identity})
	}
	return &ChargePoint{
		Service:              service,
		identity:             identity,
		centralSystemURL:     csURL,
		version:              version,
		transport:            transport,
	}, nil
}
