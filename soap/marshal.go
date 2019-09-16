package soap

import (
	"encoding/xml"
	"errors"

	"github.com/eduhenke/go-ocpp/messages/v1x/cpreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpres"
	"github.com/eduhenke/go-ocpp/messages/v1x/csreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/csres"
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
				msg = new(cpreq.BootNotification)
			case "heartbeatRequest":
				msg = new(cpreq.Heartbeat)
			case "statusNotificationRequest":
				msg = new(cpreq.StatusNotification)
			case "meterValuesRequest":
				msg = new(cpreq.MeterValues)
			case "authorizeRequest":
				msg = new(cpreq.Authorize)
			case "startTransactionRequest":
				msg = new(cpreq.StartTransaction)
			case "stopTransactionRequest":
				msg = new(cpreq.StopTransaction)
			case "firmwareStatusNotificationRequest":
				msg = new(cpreq.FirmwareStatusNotification)
			case "diagnosticsStatusNotificationRequest":
				msg = new(cpreq.DiagnosticsStatusNotification)

			// responses coming from the central system
			case "bootNotificationResponse":
				msg = new(cpres.BootNotification)
			case "heartbeatResponse":
				msg = new(cpres.Heartbeat)
			case "statusNotificationResponse":
				msg = new(cpres.StatusNotification)
			case "meterValuesResponse":
				msg = new(cpres.MeterValues)
			case "authorizeResponse":
				msg = new(cpres.Authorize)
			case "startTransactionResponse":
				msg = new(cpres.StartTransaction)
			case "stopTransactionResponse":
				msg = new(cpres.StopTransaction)
			case "firmwareStatusNotificationResponse":
				msg = new(cpres.FirmwareStatusNotification)
			case "diagnosticsStatusNotificationResponse":
				msg = new(cpres.DiagnosticsStatusNotification)

			// requests coming from the central system
			case "unlockConnectorRequest":
				msg = new(csreq.UnlockConnector)
			case "eesetRequest":
				msg = new(csreq.Reset)
			case "changeAvailabilityRequest":
				msg = new(csreq.ChangeAvailability)
			case "getDiagnosticsRequest":
				msg = new(csreq.GetDiagnostics)
			case "clearCacheRequest":
				msg = new(csreq.ClearCache)
			case "updateFirmwareRequest":
				msg = new(csreq.UpdateFirmware)
			case "changeConfigurationRequest":
				msg = new(csreq.ChangeConfiguration)
			case "remoteStartTransactionRequest":
				msg = new(csreq.RemoteStartTransaction)
			case "remoteStopTransactionRequest":
				msg = new(csreq.RemoteStopTransaction)
			case "cancelReservationRequest":
				msg = new(csreq.CancelReservation)
			case "dataTransferRequest":
				msg = new(csreq.DataTransfer)
			case "getConfigurationRequest":
				msg = new(csreq.GetConfiguration)
			case "getLocalListVersionRequest":
				msg = new(csreq.GetLocalListVersion)
			case "reserveNowRequest":
				msg = new(csreq.ReserveNow)
			case "sendLocalListRequest":
				msg = new(csreq.SendLocalList)

			// responses coming from the charger
			case "unlockConnectorResponse":
				msg = new(csres.UnlockConnector)
			case "eesetResponse":
				msg = new(csres.Reset)
			case "changeAvailabilityResponse":
				msg = new(csres.ChangeAvailability)
			case "getDiagnosticsResponse":
				msg = new(csres.GetDiagnostics)
			case "clearCacheResponse":
				msg = new(csres.ClearCache)
			case "updateFirmwareResponse":
				msg = new(csres.UpdateFirmware)
			case "changeConfigurationResponse":
				msg = new(csres.ChangeConfiguration)
			case "remoteStartTransactionResponse":
				msg = new(csres.RemoteStartTransaction)
			case "remoteStopTransactionResponse":
				msg = new(csres.RemoteStopTransaction)
			case "cancelReservationResponse":
				msg = new(csres.CancelReservation)
			case "dataTransferResponse":
				msg = new(csres.DataTransfer)
			case "getConfigurationResponse":
				msg = new(csres.GetConfiguration)
			case "getLocalListVersionResponse":
				msg = new(csres.GetLocalListVersion)
			case "reserveNowResponse":
				msg = new(csres.ReserveNow)
			case "sendLocalListResponse":
				msg = new(csres.SendLocalList)
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
