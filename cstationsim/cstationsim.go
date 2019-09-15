package cstationsim

import (
	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/internal/log"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpres"
	"github.com/eduhenke/go-ocpp/service/cs"
	"github.com/eduhenke/go-ocpp/soap"
	"time"
)

type Station struct {
	CsService cs.Service
}

func NewStation(port, csURL string, transport ocpp.Transport) *Station {
	var service cs.Service
	if transport == ocpp.JSON {
		//service = cs.NewJsonService()
	}
	if transport == ocpp.SOAP {
		service = cs.NewSoapService(csURL, &soap.CallOptions{ChargeBoxIdentity: "STATION_01"})
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

func (st *Station) heartbeat() (*cpres.Heartbeat, error) {
	rawResp, err := st.CsService.Send(&cpreq.Heartbeat{})
	if err != nil {
		return nil, err
	}
	resp, ok := rawResp.(*cpres.Heartbeat)
	if !ok {
		return nil, cpres.ErrorNotChargePointResponse
	}
	return resp, nil
}