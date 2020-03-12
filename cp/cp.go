package cp

import (
	"context"
	"errors"
	"fmt"

	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/internal/log"
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
	Run(ctx context.Context, port *string, cphandler CentralSystemMessageHandler) error
}

type chargePoint struct {
	service.CentralSystem
	identity         string
	centralSystemURL string
	version          ocpp.Version
	transport        ocpp.Transport
	conn             *ws.Conn
}

func New(identity, csURL string, version ocpp.Version, transport ocpp.Transport) (*chargePoint, error) {
	var csService service.CentralSystem
	var conn *ws.Conn
	if transport == ocpp.JSON {
		var err error
		conn, err = ws.Dial(identity, csURL, version)
		if err != nil {
			return nil, fmt.Errorf("could not dial to central system: %w", err)
		}
		csService = service.NewCentralSystemJSON(conn)
	}
	if transport == ocpp.SOAP {
		csService = service.NewCentralSystemSOAP(csURL, &soap.CallOptions{ChargeBoxIdentity: identity})
	}
	return &chargePoint{
		CentralSystem:    csService,
		identity:         identity,
		centralSystemURL: csURL,
		version:          version,
		transport:        transport,
		conn:             conn,
	}, nil
}

func (cp *chargePoint) Run(ctx context.Context, port *string, cshandler CentralSystemMessageHandler) error {
	if cp.transport == ocpp.JSON {
		if cp.conn == nil {
			return errors.New("no ws connection")
		}
		go handleWebsocket(ctx, cp.conn, cshandler)
	}
	if cp.transport == ocpp.SOAP {
		go handleSoap(ctx, cshandler)
	}
	return nil
}

func handleWebsocket(ctx context.Context, conn *ws.Conn, cshandler CentralSystemMessageHandler) {
	go func() {
		for {
			err := conn.ReadMessage()
			if err != nil {
				if !ws.IsNormalCloseError(err) {
					log.Error("On receiving a message: %w", err)
				}
				_ = conn.Close()
				log.Debug("Closed connection of central system")
				break
			}
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case req := <-conn.Requests():
			cprequest, ok := req.Request.(csreq.CentralSystemRequest)
			if !ok {
				log.Error(csreq.ErrorNotCentralSystemRequest.Error())
			}
			cpresponse, err := cshandler(cprequest)
			err = conn.SendResponse(req.MessageID, cpresponse, err)
			if err != nil {
				log.Error(err.Error())
			}
		}
	}
}

func handleSoap(ctx context.Context, cshandler CentralSystemMessageHandler) {
	panic("SOAP yet not supported in ChargePoint")
	// log.Debug("New SOAP request")
	// err := soap.Handle(w, r, func(request messages.Request, cpID string) (messages.Response, error) {
	// 	req, ok := request.(cpreq.ChargePointRequest)
	// 	if !ok {
	// 		return nil, errors.New("request is not a cprequest")
	// 	}
	// 	return cshandler(req, cpID)
	// })
	// if err != nil {
	// 	log.Error("Couldn't handle SOAP request: %w", err)
	// }
}
