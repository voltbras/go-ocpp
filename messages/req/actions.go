package req

import (
	"github.com/voltbras/go-ocpp/messages"
	"github.com/voltbras/go-ocpp/messages/v1x/cpreq"
	"github.com/voltbras/go-ocpp/messages/v1x/csreq"
)

func FromActionName(action string) messages.Request {
	switch action {
	case "Authorize":
		return &cpreq.Authorize{}
	case "BootNotification":
		return &cpreq.BootNotification{}
	// TODO: DataTransfer comes from both the CS and CP
	// case "DataTransfer":
	// 	return &cpreq.DataTransfer{}
	case "DiagnosticsStatusNotification":
		return &cpreq.DiagnosticsStatusNotification{}
	case "FirmwareStatusNotification":
		return &cpreq.FirmwareStatusNotification{}
	case "Heartbeat":
		return &cpreq.Heartbeat{}
	case "MeterValues":
		return &cpreq.MeterValues{}
	case "StartTransaction":
		return &cpreq.StartTransaction{}
	case "StatusNotification":
		return &cpreq.StatusNotification{}
	case "StopTransaction":
		return &cpreq.StopTransaction{}

	case "CancelReservation":
		return &csreq.CancelReservation{}
	case "ChangeAvailability":
		return &csreq.ChangeAvailability{}
	case "ChangeConfiguration":
		return &csreq.ChangeConfiguration{}
	case "ClearCache":
		return &csreq.ClearCache{}
	case "ClearChargingProfile":
		return &csreq.ClearChargingProfile{}
	// TODO: DataTransfer comes from both the CS and CP
	// case "DataTransfer":
	// 	return &csreq.DataTransfer{}
	case "GetCompositeSchedule":
		return &csreq.GetCompositeSchedule{}
	case "GetConfiguration":
		return &csreq.GetConfiguration{}
	case "GetDiagnostics":
		return &csreq.GetDiagnostics{}
	case "GetLocalListVersion":
		return &csreq.GetLocalListVersion{}
	case "RemoteStartTransaction":
		return &csreq.RemoteStartTransaction{}
	case "RemoteStopTransaction":
		return &csreq.RemoteStopTransaction{}
	case "ReserveNow":
		return &csreq.ReserveNow{}
	case "Reset":
		return &csreq.Reset{}
	case "SendLocalList":
		return &csreq.SendLocalList{}
	case "TriggerMessage":
		return &csreq.TriggerMessage{}
	case "UnlockConnector":
		return &csreq.UnlockConnector{}
	case "UpdateFirmware":
		return &csreq.UpdateFirmware{}
	case "SetChargingProfile":
		return &csreq.SetChargingProfile{}
	}
	return nil
}
