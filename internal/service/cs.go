package service

import (
	"github.com/voltbras/go-ocpp/messages/v1x/cpreq"
	"github.com/voltbras/go-ocpp/messages/v1x/cpresp"
	"github.com/voltbras/go-ocpp/soap"
	"github.com/voltbras/go-ocpp/ws"
)

type CentralSystem interface {
	Send(cpreq.ChargePointRequest) (cpresp.ChargePointResponse, error)
}

type CentralSystemSOAP struct {
	*SOAP
}

func NewCentralSystemSOAP(csURL string, options *soap.CallOptions) CentralSystem {
	return &CentralSystemSOAP{NewSOAP(csURL, options)}
}

func (service *CentralSystemSOAP) Send(req cpreq.ChargePointRequest) (cpresp.ChargePointResponse, error) {
	rawResp, err := service.SOAP.Send(req)
	if err != nil {
		return nil, err
	}
	resp, ok := rawResp.(cpresp.ChargePointResponse)
	if !ok {
		return nil, cpresp.ErrorNotChargePointResponse
	}
	return resp, nil
}

type CentralSystemJSON struct {
	*JSON
}

func NewCentralSystemJSON(conn *ws.Conn) CentralSystem {
	return &CentralSystemJSON{NewJSON(conn)}
}

func (service *CentralSystemJSON) Send(req cpreq.ChargePointRequest) (cpresp.ChargePointResponse, error) {
	rawResp, err := service.JSON.Send(req)
	if err != nil {
		return nil, err
	}
	resp, ok := rawResp.(cpresp.ChargePointResponse)
	if !ok {
		return nil, cpresp.ErrorNotChargePointResponse
	}
	return resp, nil
}
