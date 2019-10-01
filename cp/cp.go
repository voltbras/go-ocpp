package cp

import (
	"fmt"

	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/cs"
	"github.com/eduhenke/go-ocpp/internal/log"
	"github.com/eduhenke/go-ocpp/soap"
	"github.com/eduhenke/go-ocpp/ws"
)

type ChargePoint struct {
	cs.Service
	identity         string
	centralSystemURL string
	version          ocpp.Version
	transport        ocpp.Transport
}

func NewChargePoint(identity, csURL string, version ocpp.Version, transport ocpp.Transport) (*ChargePoint, error) {
	var service cs.Service
	if transport == ocpp.JSON {
		conn, err := ws.Dial(identity, csURL, version)
		if err != nil {
			return nil, fmt.Errorf("could not dial to central system %w", err)
		}
		go func() {
			for {
				err := conn.ReadMessage()
				if err != nil {
					log.Error("On receiving a message: %w\n", err)
					_ = conn.Close()
					break
				}
			}
		}()
		service = cs.NewJsonService(conn)
	}
	if transport == ocpp.SOAP {
		service = cs.NewSoapService(csURL, &soap.CallOptions{ChargeBoxIdentity: identity})
	}
	return &ChargePoint{
		Service:          service,
		identity:         identity,
		centralSystemURL: csURL,
		version:          version,
		transport:        transport,
	}, nil
}
