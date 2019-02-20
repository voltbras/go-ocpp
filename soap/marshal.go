package soap

import (
	"encoding/xml"
	"errors"

	"github.com/eduhenke/go-ocpp/messages"
)

func Unmarshal(data []byte) (receivedEnvelope, error) {
	parsed := receivedEnvelope{}
	err := xml.Unmarshal(data, &parsed)
	if err != nil {
		return receivedEnvelope{}, err
	}
	return parsed, nil
}

func Marshal(data interface{}, err error, NS string) ([]byte, error) {
	respEnv := toSendEnvelope{XMLNS: NS}
	if err != nil {
		respEnv.Body.Fault = &toSendFault{Reason: err.Error()}
	}
	respEnv.Body.Content = data

	return xml.Marshal(respEnv)
}

func (b *receivedBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// decode inner elements
	for {
		t, err := d.Token()
		if err != nil {
			return err
		}
		var msg interface{}
		switch tt := t.(type) {
		case xml.StartElement:
			switch tt.Name.Local {
			case "Fault":
				fault := receivedFault{}
				err := d.DecodeElement(&fault, &tt)
				if err != nil {
					return err
				}
				b.Fault = &fault
			// requests coming from the charger
			case "bootNotificationRequest":
				msg = new(messages.BootNotificationRequest)
			case "heartbeatRequest":
				msg = new(messages.HeartbeatRequest)
			case "statusNotificationRequest":
				msg = new(messages.StatusNotificationRequest)
			case "meterValuesRequest":
				msg = new(messages.MeterValuesRequest)
			case "authorizeRequest":
				msg = new(messages.AuthorizeRequest)
			case "startTransactionRequest":
				msg = new(messages.StartTransactionRequest)
			case "stopTransactionRequest":
				msg = new(messages.StopTransactionRequest)
			case "firmwareStatusNotificationRequest":
				msg = new(messages.FirmwareStatusNotificationRequest)
			case "diagnosticsStatusNotificationRequest":
				msg = new(messages.DiagnosticsStatusNotificationRequest)

			// responses coming from the central system
			case "bootNotificationResponse":
				msg = new(messages.BootNotificationResponse)
			case "heartbeatResponse":
				msg = new(messages.HeartbeatResponse)
			case "statusNotificationResponse":
				msg = new(messages.StatusNotificationResponse)
			case "meterValuesResponse":
				msg = new(messages.MeterValuesResponse)
			case "authorizeResponse":
				msg = new(messages.AuthorizeResponse)
			case "startTransactionResponse":
				msg = new(messages.StartTransactionResponse)
			case "stopTransactionResponse":
				msg = new(messages.StopTransactionResponse)
			case "firmwareStatusNotificationResponse":
				msg = new(messages.FirmwareStatusNotificationResponse)
			case "diagnosticsStatusNotificationResponse":
				msg = new(messages.DiagnosticsStatusNotificationResponse)

			// requests coming from the central system
			case "unlockConnectorRequest":
				msg = new(messages.UnlockConnectorRequest)
			case "eesetRequest":
				msg = new(messages.ResetRequest)
			case "changeAvailabilityRequest":
				msg = new(messages.ChangeAvailabilityRequest)
			case "getDiagnosticsRequest":
				msg = new(messages.GetDiagnosticsRequest)
			case "clearCacheRequest":
				msg = new(messages.ClearCacheRequest)
			case "updateFirmwareRequest":
				msg = new(messages.UpdateFirmwareRequest)
			case "changeConfigurationRequest":
				msg = new(messages.ChangeConfigurationRequest)
			case "remoteStartTransactionRequest":
				msg = new(messages.RemoteStartTransactionRequest)
			case "remoteStopTransactionRequest":
				msg = new(messages.RemoteStopTransactionRequest)
			case "cancelReservationRequest":
				msg = new(messages.CancelReservationRequest)
			case "dataTransferRequest":
				msg = new(messages.DataTransferRequest)
			case "getConfigurationRequest":
				msg = new(messages.GetConfigurationRequest)
			case "getLocalListVersionRequest":
				msg = new(messages.GetLocalListVersionRequest)
			case "reserveNowRequest":
				msg = new(messages.ReserveNowRequest)
			case "sendLocalListRequest":
				msg = new(messages.SendLocalListRequest)

			// responses coming from the charger
			case "unlockConnectorResponse":
				msg = new(messages.UnlockConnectorResponse)
			case "eesetResponse":
				msg = new(messages.ResetResponse)
			case "changeAvailabilityResponse":
				msg = new(messages.ChangeAvailabilityResponse)
			case "getDiagnosticsResponse":
				msg = new(messages.GetDiagnosticsResponse)
			case "clearCacheResponse":
				msg = new(messages.ClearCacheResponse)
			case "updateFirmwareResponse":
				msg = new(messages.UpdateFirmwareResponse)
			case "changeConfigurationResponse":
				msg = new(messages.ChangeConfigurationResponse)
			case "remoteStartTransactionResponse":
				msg = new(messages.RemoteStartTransactionResponse)
			case "remoteStopTransactionResponse":
				msg = new(messages.RemoteStopTransactionResponse)
			case "cancelReservationResponse":
				msg = new(messages.CancelReservationResponse)
			case "dataTransferResponse":
				msg = new(messages.DataTransferResponse)
			case "getConfigurationResponse":
				msg = new(messages.GetConfigurationResponse)
			case "getLocalListVersionResponse":
				msg = new(messages.GetLocalListVersionResponse)
			case "reserveNowResponse":
				msg = new(messages.ReserveNowResponse)
			case "sendLocalListResponse":
				msg = new(messages.SendLocalListResponse)
			default:
				return errors.New("not implemented unmarshal for this OCPP message:" + tt.Name.Local)
			}
			// known child element found, decode it
			if msg != nil {
				err = d.DecodeElement(msg, &tt)
				if err != nil {
					return err
				}
				b.Content = msg
				msg = nil
			}
		case xml.EndElement:
			if tt == start.End() {
				return nil
			}
		}
	}
}
