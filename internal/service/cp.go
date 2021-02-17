package service

import (
	"github.com/voltbras/go-ocpp/messages/v1x/csreq"
	"github.com/voltbras/go-ocpp/messages/v1x/csresp"
	"github.com/voltbras/go-ocpp/soap"
	"github.com/voltbras/go-ocpp/ws"
)

type ChargePoint interface {
	Send(csreq.CentralSystemRequest) (csresp.CentralSystemResponse, error)
}

type ChargePointSOAP struct {
	*SOAP
}

func NewChargePointSOAP(csURL string, options *soap.CallOptions) ChargePoint {
	return &ChargePointSOAP{NewSOAP(csURL, options)}
}

func (service *ChargePointSOAP) Send(req csreq.CentralSystemRequest) (csresp.CentralSystemResponse, error) {
	rawResp, err := service.SOAP.Send(req)
	if err != nil {
		return nil, err
	}
	resp, ok := rawResp.(csresp.CentralSystemResponse)
	if !ok {
		return nil, csresp.ErrorNotCentralSystemResponse
	}
	return resp, nil
}

type ChargePointJSON struct {
	*JSON
}

func NewChargePointJSON(conn *ws.Conn) ChargePoint {
	return &ChargePointJSON{NewJSON(conn)}
}

func (service *ChargePointJSON) Send(req csreq.CentralSystemRequest) (csresp.CentralSystemResponse, error) {
	rawResp, err := service.JSON.Send(req)
	if err != nil {
		return nil, err
	}
	resp, ok := rawResp.(csresp.CentralSystemResponse)
	if !ok {
		return nil, csresp.ErrorNotCentralSystemResponse
	}
	return resp, nil
}
