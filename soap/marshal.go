package soap

import (
	"encoding/xml"
	"errors"

	"github.com/eduhenke/go-ocpp/messages/v1x/cpreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpresp"
	"github.com/eduhenke/go-ocpp/messages/v1x/csreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/csresp"
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
				msg = new(cpresp.BootNotification)
			case "heartbeatResponse":
				msg = new(cpresp.Heartbeat)
			case "statusNotificationResponse":
				msg = new(cpresp.StatusNotification)
			case "meterValuesResponse":
				msg = new(cpresp.MeterValues)
			case "authorizeResponse":
				msg = new(cpresp.Authorize)
			case "startTransactionResponse":
				msg = new(cpresp.StartTransaction)
			case "stopTransactionResponse":
				msg = new(cpresp.StopTransaction)
			case "firmwareStatusNotificationResponse":
				msg = new(cpresp.FirmwareStatusNotification)
			case "diagnosticsStatusNotificationResponse":
				msg = new(cpresp.DiagnosticsStatusNotification)

			// requests coming from the central system
			case "unlockConnectorRequest":
				msg = new(csreq.UnlockConnector)
			case "resetRequest":
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
				msg = new(csresp.UnlockConnector)
			case "resetResponse":
				msg = new(csresp.Reset)
			case "changeAvailabilityResponse":
				msg = new(csresp.ChangeAvailability)
			case "getDiagnosticsResponse":
				msg = new(csresp.GetDiagnostics)
			case "clearCacheResponse":
				msg = new(csresp.ClearCache)
			case "updateFirmwareResponse":
				msg = new(csresp.UpdateFirmware)
			case "changeConfigurationResponse":
				msg = new(csresp.ChangeConfiguration)
			case "remoteStartTransactionResponse":
				msg = new(csresp.RemoteStartTransaction)
			case "remoteStopTransactionResponse":
				msg = new(csresp.RemoteStopTransaction)
			case "cancelReservationResponse":
				msg = new(csresp.CancelReservation)
			case "dataTransferResponse":
				msg = new(csresp.DataTransfer)
			case "getConfigurationResponse":
				msg = new(csresp.GetConfiguration)
			case "getLocalListVersionResponse":
				msg = new(csresp.GetLocalListVersion)
			case "reserveNowResponse":
				msg = new(csresp.ReserveNow)
			case "sendLocalListResponse":
				msg = new(csresp.SendLocalList)
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
