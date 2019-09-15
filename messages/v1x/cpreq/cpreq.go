package cpreq

import (
	"github.com/eduhenke/go-ocpp/messages"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpres"
)

// ChargePointRequest is a request coming from the chargepoint to the central system
type ChargePointRequest interface {
	messages.Request
	IsChargePointRequest()
	GetChargePointResponse() cpres.ChargePointResponse
}

type chargepointRequest struct{}

func (cpreq *chargepointRequest) IsChargePointRequest() {}
func (cpreq *chargepointRequest) IsRequest() {}

func GetRequestFromActionName(action string) ChargePointRequest {
	switch action {
	case "Authorize":
		return &Authorize{}
	case "BootNotification":
		return &BootNotification{}
	case "DataTransfer":
		return &DataTransfer{}
	case "DiagnosticsStatusNotification":
		return &DiagnosticsStatusNotification{}
	case "FirmwareStatusNotification":
		return &FirmwareStatusNotification{}
	case "Heartbeat":
		return &Heartbeat{}
	case "MeterValues":
		return &MeterValues{}
	case "StartTransaction":
		return &StartTransaction{}
	case "StatusNotification":
		return &StatusNotification{}
	case "StopTransaction":
		return &StopTransaction{}
	}
	return nil
}
