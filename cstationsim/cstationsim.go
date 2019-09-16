package cstationsim

import (
	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/internal/log"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpresp"
	"github.com/eduhenke/go-ocpp/service/cs"
	"github.com/eduhenke/go-ocpp/soap"
	"github.com/eduhenke/go-ocpp/wsconn"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type Station struct {
	CsService cs.Service
}

func NewStation(identity, csURL string, transport ocpp.Transport) *Station {
	var service cs.Service
	if transport == ocpp.JSON {
		dialer := websocket.Dialer{
			Subprotocols: []string{string(ocpp.V16)},
		}
		socket, _, err := dialer.Dial(csURL+"/"+identity, http.Header{})
		if err != nil {
			panic(err)
		}
		conn := wsconn.NewConn(socket)
		service = cs.NewJsonService(conn)
	}
	if transport == ocpp.SOAP {
		service = cs.NewSoapService(csURL, &soap.CallOptions{ChargeBoxIdentity: identity})
	}
	return &Station{
		CsService: service,
	}
}

func (st *Station) Run() {
	for range time.Tick(5 * time.Second) {
		res, err := st.heartbeat()
		if err != nil {
			log.Error("Couldn't send heartbeat to central system: %w", err)
			continue
		}
		log.Debug("Received response from central system: %+v\n", *res)
	}
}

func (st *Station) heartbeat() (*cpresp.Heartbeat, error) {
	rawResp, err := st.CsService.Send(&cpreq.Heartbeat{})
	if err != nil {
		return nil, err
	}
	resp, ok := rawResp.(*cpresp.Heartbeat)
	if !ok {
		return nil, cpresp.ErrorNotChargePointResponse
	}
	return resp, nil
}