package soap

import (
	"testing"

	"github.com/eduhenke/go-ocpp/messages/v15"
)

const bootNotification = `
<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://www.w3.org/2003/05/soap-envelope">
	<S:Header xmlns:wsa5="http://www.w3.org/2005/08/addressing">
		<wsa5:To>http:	env.Body.OCPPMessage.(central_system.BootNotificationRequest)//192.168.2.197:8080</wsa5:To>
		<wsa5:Action>/BootNotification</wsa5:Action>
		<wsa5:MessageID>uuid:543a0a2e-5d97-44ac-a275-6372b8a5c44c</wsa5:MessageID>
		<wsa5:From>
			<wsa5:Address />
		</wsa5:From>
	</S:Header>
	<S:Body>
		<ocpp:bootNotificationRequest xmlns:ocpp="urn://Ocpp/Cs/2012/06/"
			xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
			<ocpp:chargePointVendor>VENDOR</ocpp:chargePointVendor>
			<ocpp:chargePointModel>MODEL-1000</ocpp:chargePointModel>
			<ocpp:chargePointSerialNumber>1337</ocpp:chargePointSerialNumber>
			<ocpp:chargeBoxSerialNumber>1337</ocpp:chargeBoxSerialNumber>
			<ocpp:firmwareVersion>1.7.0</ocpp:firmwareVersion>
		</ocpp:bootNotificationRequest>
	</S:Body>
</S:Envelope>
`

func TestSOAPUnmarshal(t *testing.T) {
	env, err := Unmarshal([]byte(bootNotification))
	if err != nil {
		t.Fatal(err)
	}
	_, ok := env.Body.Content.(*messages.BootNotificationRequest)
	if !ok {
		t.Fatal("Could not assert message type")
	}
}

const meterValues = `
<?xml version="1.0" encoding="UTF-8"?><S:Envelope xmlns:S="http://www.w3.org/2003/05/soap-envelope">
  <S:Header xmlns:wsa5="http://www.w3.org/2005/08/addressing">
    <wsa5:To>http://192.168.0.1:12811</wsa5:To>
    <wsa5:Action>/MeterValues</wsa5:Action>
    <wsa5:MessageID>uuid:77638163-b4f8-4cbe-af27-54362ea20bd6</wsa5:MessageID>
    <wsa5:From>
      <wsa5:Address>http://192.168.0.2:12811/Ocpp/ChargePointService</wsa5:Address>
    </wsa5:From>
    <chargeBoxIdentity xmlns="urn://Ocpp/Cs/2012/06/">BANANA</chargeBoxIdentity>
  </S:Header>
  <S:Body>
    <ocpp:meterValuesRequest xmlns:ocpp="urn://Ocpp/Cs/2012/06/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
      <ocpp:connectorId>0</ocpp:connectorId>
      <ocpp:values>
        <ocpp:timestamp>2019-04-05T21:45:00.000Z</ocpp:timestamp>
        <ocpp:value context="Sample.Clock" format="Raw" location="Outlet" measurand="Energy.Active.Import.Register" unit="Wh">0</ocpp:value>
      </ocpp:values>
    </ocpp:meterValuesRequest>
  </S:Body>
</S:Envelope>
`

func TestMeterValues(t *testing.T) {
	env, err := Unmarshal([]byte(meterValues))
	if err != nil {
		t.Fatal(err)
	}
	req, ok := env.Body.Content.(*messages.MeterValuesRequest)
	if !ok {
		t.Fatal("Could not assert message type")
	}
	if req.ConnectorId != 0 {
		t.Fatal("wrong connector")
	}
	value := req.Values[0].SampledValues[0]
	if value.Value != "0" || value.Location != "Outlet" ||
		value.Measurand != "Energy.Active.Import.Register" ||
		value.Unit != "Wh" || value.Format != "Raw" ||
		value.Context != "Sample.Clock" {
		t.Fail()
	}
}
