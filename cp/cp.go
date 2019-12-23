package cp

import (
	"fmt"

	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/internal/log"
	"github.com/eduhenke/go-ocpp/internal/service"
	"github.com/eduhenke/go-ocpp/soap"
	"github.com/eduhenke/go-ocpp/ws"
)

type ChargePoint struct {
	service.CentralSystem
	identity         string
	centralSystemURL string
	version          ocpp.Version
	transport        ocpp.Transport
}

func NewChargePoint(identity, csURL string, version ocpp.Version, transport ocpp.Transport) (*ChargePoint, error) {
	var csService service.CentralSystem
	if transport == ocpp.JSON {
		conn, err := ws.Dial(identity, csURL, version)
		if err != nil {
			return nil, fmt.Errorf("could not dial to central system: %w", err)
		}
		go func() {
			for {
				err := conn.ReadMessage()
				if err != nil {
					if !ws.IsNormalCloseError(err) {
						log.Error("On receiving a message: %w\n", err)
					}
					_ = conn.Close()
					break
				}
			}
		}()
		csService = service.NewCentralSystemJSON(conn)
	}
	if transport == ocpp.SOAP {
		csService = service.NewCentralSystemSOAP(csURL, &soap.CallOptions{ChargeBoxIdentity: identity})
	}
	return &ChargePoint{
		CentralSystem:    csService,
		identity:         identity,
		centralSystemURL: csURL,
		version:          version,
		transport:        transport,
	}, nil
}
