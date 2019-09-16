package res

import (
	"github.com/eduhenke/go-ocpp/messages"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpres"
	"github.com/eduhenke/go-ocpp/messages/v1x/csres"
)

func FromActionName(action string) messages.Response {
	switch action {
	case "CancelReservation":
		return &csres.CancelReservation{}
	case "ChangeAvailability":
		return &csres.ChangeAvailability{}
	case "ChangeConfiguration":
		return &csres.ChangeConfiguration{}
	case "ClearCache":
		return &csres.ClearCache{}
	case "ClearChargingProfile":
		return &csres.ClearChargingProfile{}
	// TODO: DataTransfer comes from both the CS and CP
	// case "DataTransfer":
	// 	return &csres.DataTransfer{}
	case "GetCompositeSchedule":
		return &csres.GetCompositeSchedule{}
	case "GetConfiguration":
		return &csres.GetConfiguration{}
	case "GetDiagnostics":
		return &csres.GetDiagnostics{}
	case "GetLocalListVersion":
		return &csres.GetLocalListVersion{}
	case "RemoteStartTransaction":
		return &csres.RemoteStartTransaction{}
	case "RemoteStopTransaction":
		return &csres.RemoteStopTransaction{}
	case "ReserveNow":
		return &csres.ReserveNow{}
	case "Reset":
		return &csres.Reset{}
	case "SendLocalList":
		return &csres.SendLocalList{}
	case "TriggerMessage":
		return &csres.TriggerMessage{}
	case "UnlockConnector":
		return &csres.UnlockConnector{}
	case "UpdateFirmware":
		return &csres.UpdateFirmware{}
	case "SetChargingProfile":
		return &csres.SetChargingProfile{}

	case "Authorize":
		return &cpres.Authorize{}
	case "BootNotification":
		return &cpres.BootNotification{}
	// TODO: DataTransfer comes from both the CS and CP
	// case "DataTransfer":
	// 	return &cpres.DataTransfer{}
	case "DiagnosticsStatusNotification":
		return &cpres.DiagnosticsStatusNotification{}
	case "FirmwareStatusNotification":
		return &cpres.FirmwareStatusNotification{}
	case "Heartbeat":
		return &cpres.Heartbeat{}
	case "MeterValues":
		return &cpres.MeterValues{}
	case "StartTransaction":
		return &cpres.StartTransaction{}
	case "StatusNotification":
		return &cpres.StatusNotification{}
	case "StopTransaction":
		return &cpres.StopTransaction{}
	}
	return nil
}
