package cpreq

import (
	"encoding/xml"
	"time"

	"github.com/eduhenke/go-ocpp/messages"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpres"
)

// Authorize
type Authorize struct {
	chargepointRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ authorizeRequest"`

	IdTag string `json:"idTag" xml:"idTag,omitempty"`
}

// BootNotification
type BootNotification struct {
	chargepointRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ bootNotificationRequest"`

	ChargeBoxSerialNumber   string `json:"chargeBoxSerialNumber,omitempty" xml:"chargeBoxSerialNumber,omitempty"`
	ChargePointModel        string `json:"chargePointModel" xml:"chargePointModel,omitempty"`
	ChargePointSerialNumber string `json:"chargePointSerialNumber,omitempty" xml:"chargePointSerialNumber,omitempty"`
	ChargePointVendor       string `json:"chargePointVendor" xml:"chargePointVendor,omitempty"`
	FirmwareVersion         string `json:"firmwareVersion,omitempty" xml:"firmwareVersion,omitempty"`
	Iccid                   string `json:"iccid,omitempty" xml:"iccid,omitempty"`
	Imsi                    string `json:"imsi,omitempty" xml:"imsi,omitempty"`
	MeterSerialNumber       string `json:"meterSerialNumber,omitempty" xml:"meterSerialNumber,omitempty"`
	MeterType               string `json:"meterType,omitempty" xml:"meterType,omitempty"`
}

// DataTransfer
type DataTransfer struct {
	chargepointRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ dataTransferRequest"`

	Data      string `json:"data,omitempty" xml:"data,omitempty"`
	MessageId string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	VendorId  string `json:"vendorId" xml:"vendorId,omitempty"`
}

// DiagnosticsStatusNotification
type DiagnosticsStatusNotification struct {
	chargepointRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ diagnosticsStatusNotificationRequest"`

	Status string `json:"status" xml:"status,omitempty"`
}

// FirmwareStatusNotification
type FirmwareStatusNotification struct {
	chargepointRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ firmwareStatusNotificationRequest"`

	Status string `json:"status" xml:"status,omitempty"`
}

// Heartbeat
type Heartbeat struct {
	chargepointRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ heartbeatRequest"`
}

// MeterValueItems
type MeterValueItems struct {
	// SampledValues in OCPP v1.5 is value
	SampledValues []*SampledValue `json:"sampledValue" xml:"value,omitempty"`
	Timestamp     time.Time       `json:"timestamp" xml:"timestamp,omitempty"`
}

// MeterValues
type MeterValues struct {
	chargepointRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ meterValuesRequest"`

	ConnectorId int `json:"connectorId" xml:"connectorId,omitempty"`
	// MeterValue in OCPP v1.5 is values
	MeterValue    []*MeterValueItems `json:"meterValue" xml:"values,omitempty"`
	TransactionId int32              `json:"transactionId,omitempty" xml:"transactionId,omitempty"`
}

// SampledValue
type SampledValue struct {
	Context   string `json:"context,omitempty" xml:"context,attr,omitempty"`
	Format    string `json:"format,omitempty" xml:"format,attr,omitempty"`
	Location  string `json:"location,omitempty" xml:"location,attr,omitempty"`
	Measurand string `json:"measurand,omitempty" xml:"measurand,attr,omitempty"`
	// Phase not present in OCPP v1.5
	Phase string `json:"phase,omitempty" xml:"-"`
	Unit  string `json:"unit,omitempty" xml:"unit,attr,omitempty"`

	Value string `json:"value" xml:",chardata"`
}

// StartTransaction
type StartTransaction struct {
	chargepointRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ startTransactionRequest"`

	ConnectorId   int       `json:"connectorId" xml:"connectorId,omitempty"`
	IdTag         string    `json:"idTag" xml:"idTag,omitempty"`
	MeterStart    int       `json:"meterStart" xml:"meterStart,omitempty"`
	ReservationId int       `json:"reservationId,omitempty" xml:"reservationId,omitempty"`
	Timestamp     time.Time `json:"timestamp" xml:"timestamp,omitempty"`
}

// StatusNotification
type StatusNotification struct {
	chargepointRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ statusNotificationRequest"`

	ConnectorId     int        `json:"connectorId" xml:"connectorId,omitempty"`
	ErrorCode       string     `json:"errorCode" xml:"errorCode,omitempty"`
	Info            string     `json:"info,omitempty" xml:"info,omitempty"`
	Status          string     `json:"status" xml:"status,omitempty"`
	Timestamp       *time.Time `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	VendorErrorCode string     `json:"vendorErrorCode,omitempty" xml:"vendorErrorCode,omitempty"`
	VendorId        string     `json:"vendorId,omitempty" xml:"vendorId,omitempty"`
}

// StopTransaction
type StopTransaction struct {
	chargepointRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ stopTransactionRequest"`

	IdTag           string                  `json:"idTag,omitempty" xml:"idTag,omitempty"`
	MeterStop       int                     `json:"meterStop" xml:"meterStop,omitempty"`
	Reason          string                  `json:"reason,omitempty" xml:"reason,omitempty"`
	Timestamp       time.Time               `json:"timestamp" xml:"timestamp,omitempty"`
	TransactionData []*TransactionDataItems `json:"transactionData,omitempty" xml:"transactionData,omitempty"`
	TransactionId   int                     `json:"transactionId" xml:"transactionId,omitempty"`
}

// TransactionDataItems
type TransactionDataItems struct {
	SampledValues []*SampledValue `json:"sampledValue" xml:"sampledValue,omitempty"`
	Timestamp     time.Time       `json:"timestamp" xml:"timestamp,omitempty"`
}

func (m *Authorize) Action() string                     { return "Authorize" }
func (m *BootNotification) Action() string              { return "BootNotification" }
func (m *DataTransfer) Action() string                  { return "DataTransfer" }
func (m *DiagnosticsStatusNotification) Action() string { return "DiagnosticsStatusNotification" }
func (m *FirmwareStatusNotification) Action() string    { return "FirmwareStatusNotification" }
func (m *Heartbeat) Action() string                     { return "Heartbeat" }
func (m *MeterValues) Action() string                   { return "MeterValues" }
func (m *StartTransaction) Action() string              { return "StartTransaction" }
func (m *StatusNotification) Action() string            { return "StatusNotification" }
func (m *StopTransaction) Action() string               { return "StopTransaction" }

func (m *Authorize) GetResponse() messages.Response        { return &cpres.Authorize{} }
func (m *BootNotification) GetResponse() messages.Response { return &cpres.BootNotification{} }
func (m *DataTransfer) GetResponse() messages.Response     { return &cpres.DataTransfer{} }
func (m *DiagnosticsStatusNotification) GetResponse() messages.Response {
	return &cpres.DiagnosticsStatusNotification{}
}
func (m *FirmwareStatusNotification) GetResponse() messages.Response {
	return &cpres.FirmwareStatusNotification{}
}
func (m *Heartbeat) GetResponse() messages.Response          { return &cpres.Heartbeat{} }
func (m *MeterValues) GetResponse() messages.Response        { return &cpres.MeterValues{} }
func (m *StartTransaction) GetResponse() messages.Response   { return &cpres.StartTransaction{} }
func (m *StatusNotification) GetResponse() messages.Response { return &cpres.StatusNotification{} }
func (m *StopTransaction) GetResponse() messages.Response    { return &cpres.StopTransaction{} }
