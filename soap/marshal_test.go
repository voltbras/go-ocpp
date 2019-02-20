package soap

import (
	"testing"

	"github.com/eduhenke/go-ocpp/messages"
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
