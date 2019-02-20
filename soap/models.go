package soap

import (
	"encoding/xml"

	"github.com/eduhenke/go-ocpp/messages"
)

type Envelope receivedEnvelope
type Header receivedHeader
type Body receivedBody

type receivedEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  *receivedHeader
	Body    receivedBody
}

type receivedFrom struct {
	XMLName xml.Name `xml:"From"`
	Address string   `xml:"Address"`
}

type receivedHeader struct {
	XMLName xml.Name `xml:"Header"`

	To                string `xml:"To"`
	Action            string `xml:"Action"`
	MessageID         string `xml:"MessageID"`
	From              receivedFrom
	ChargeBoxIdentity string `xml:"urn://Ocpp/Cs/2012/06/ chargeBoxIdentity"`
}

type receivedBody struct {
	XMLName xml.Name `xml:"Body"`

	Content interface{}
	Fault   *receivedFault
}

type receivedFault struct {
	XMLName xml.Name `xml:"Fault,omitempty"`

	Reason struct {
		XMLName xml.Name `xml:"Reason,omitempty"`
		Text    string   `xml:"Text,omitempty"`
	}
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

func (f *receivedFault) Error() string {
	return f.Reason.Text
}

type toSendEnvelope struct {
	XMLName xml.Name `xml:"S:Envelope"`
	XMLNS   string   `xml:"xmlns:S,attr"`

	Header *toSendHeader
	Body   toSendBody
}

type toSendFrom struct {
	XMLName xml.Name `xml:"wsa5:From"`
	Address string   `xml:"wsa5:Address"`
}

type toSendHeader struct {
	XMLName xml.Name `xml:"S:Header"`
	XMLNS   string   `xml:"xmlns:wsa5,attr"`

	To                string `xml:"wsa5:To"`
	Action            string `xml:"wsa5:Action"`
	MessageID         string `xml:"wsa5:MessageID"`
	From              toSendFrom
	ChargeBoxIdentity messages.ChargeBoxIdentity `xml:"urn://Ocpp/Cs/2012/06/ chargeBoxIdentity"`
}

type toSendBody struct {
	XMLName xml.Name `xml:"S:Body"`

	Content interface{}
	Fault   *toSendFault
}

type toSendFault struct {
	XMLName xml.Name `xml:"S:Fault,omitempty"`

	Code   string `xml:"Code,omitempty"`
	Reason string `xml:"reason,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

func (f *toSendFault) Error() string {
	return f.Reason
}
