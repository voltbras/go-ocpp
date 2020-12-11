package cp

import (
	"context"
	"errors"
	"fmt"

	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/internal/service"
	"github.com/eduhenke/go-ocpp/messages/v1x/csreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/csresp"
	"github.com/eduhenke/go-ocpp/soap"
	"github.com/eduhenke/go-ocpp/ws"
)

// CentralSystemMessageHandler handles the OCPP messages coming from the central system
type CentralSystemMessageHandler func(cprequest csreq.CentralSystemRequest) (csresp.CentralSystemResponse, error)

type ChargePoint interface {
	service.CentralSystem
	// Run the charge point on the given port
	// and handles each incoming CentralSystemRequest
	Run(port *string, cphandler CentralSystemMessageHandler) error
	Identity() string

	// WS related
	Connection() *ws.Conn
	WaitConnect() <-chan struct{}
	WaitDisconnect() <-chan struct{}
}

type chargePoint struct {
	service.CentralSystem
	identity         string
	centralSystemURL string
	version          ocpp.Version
	transport        ocpp.Transport
	conn             *ws.Conn
	ctx              context.Context
	connListeners    []func()
}

func New(ctx context.Context, identity, csURL string, version ocpp.Version, transport ocpp.Transport) (*chargePoint, error) {
	cp := &chargePoint{
		identity:         identity,
		centralSystemURL: csURL,
		version:          version,
		transport:        transport,
		ctx:              ctx,
		connListeners:    []func(){},
	}
	if transport == ocpp.JSON {
		err := cp.getNewWebsocketConnection()
		if err != nil {
			return nil, fmt.Errorf("could not dial to central system: %w", err)
		}
		go cp.handleWebsocketConnection()
	}
	if transport == ocpp.SOAP {
		cp.CentralSystem = service.NewCentralSystemSOAP(csURL, &soap.CallOptions{ChargeBoxIdentity: identity})
	}
	return cp, nil
}

func (cp *chargePoint) Identity() string {
	return cp.identity
}

func (cp *chargePoint) Run(port *string, cshandler CentralSystemMessageHandler) error {
	if cp.transport == ocpp.JSON {
		if cp.conn == nil {
			return errors.New("no ws connection")
		}
		go cp.handleWebsocketRequests(cshandler)
	}
	if cp.transport == ocpp.SOAP {
		if port == nil {
			return errors.New("port must be set when running a SOAP ChargePoint")
		}
		return handleSoap(cp.ctx, *port, cshandler)
	}
	return nil
}
