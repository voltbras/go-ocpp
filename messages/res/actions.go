package res

import (
	"github.com/voltbras/go-ocpp/messages"
	"github.com/voltbras/go-ocpp/messages/v1x/cpresp"
	"github.com/voltbras/go-ocpp/messages/v1x/csresp"
)

func FromActionName(action string) messages.Response {
	switch action {
	case "CancelReservation":
		return &csresp.CancelReservation{}
	case "ChangeAvailability":
		return &csresp.ChangeAvailability{}
	case "ChangeConfiguration":
		return &csresp.ChangeConfiguration{}
	case "ClearCache":
		return &csresp.ClearCache{}
	case "ClearChargingProfile":
		return &csresp.ClearChargingProfile{}
	// TODO: DataTransfer comes from both the CS and CP
	// case "DataTransfer":
	// 	return &csresp.DataTransfer{}
	case "GetCompositeSchedule":
		return &csresp.GetCompositeSchedule{}
	case "GetConfiguration":
		return &csresp.GetConfiguration{}
	case "GetDiagnostics":
		return &csresp.GetDiagnostics{}
	case "GetLocalListVersion":
		return &csresp.GetLocalListVersion{}
	case "RemoteStartTransaction":
		return &csresp.RemoteStartTransaction{}
	case "RemoteStopTransaction":
		return &csresp.RemoteStopTransaction{}
	case "ReserveNow":
		return &csresp.ReserveNow{}
	case "Reset":
		return &csresp.Reset{}
	case "SendLocalList":
		return &csresp.SendLocalList{}
	case "TriggerMessage":
		return &csresp.TriggerMessage{}
	case "UnlockConnector":
		return &csresp.UnlockConnector{}
	case "UpdateFirmware":
		return &csresp.UpdateFirmware{}
	case "SetChargingProfile":
		return &csresp.SetChargingProfile{}

	case "Authorize":
		return &cpresp.Authorize{}
	case "BootNotification":
		return &cpresp.BootNotification{}
	// TODO: DataTransfer comes from both the CS and CP
	// case "DataTransfer":
	// 	return &cpresp.DataTransfer{}
	case "DiagnosticsStatusNotification":
		return &cpresp.DiagnosticsStatusNotification{}
	case "FirmwareStatusNotification":
		return &cpresp.FirmwareStatusNotification{}
	case "Heartbeat":
		return &cpresp.Heartbeat{}
	case "MeterValues":
		return &cpresp.MeterValues{}
	case "StartTransaction":
		return &cpresp.StartTransaction{}
	case "StatusNotification":
		return &cpresp.StatusNotification{}
	case "StopTransaction":
		return &cpresp.StopTransaction{}
	}
	return nil
}
