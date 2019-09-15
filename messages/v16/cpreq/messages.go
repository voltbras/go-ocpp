package cpreq

import (
	"time"

	"github.com/eduhenke/go-ocpp/messages/v16/cpres"
)

// Authorize
type Authorize struct {
	chargepointRequest
	IdTag string `json:"idTag"`
}

// BootNotification
type BootNotification struct {
	chargepointRequest
	ChargeBoxSerialNumber   string `json:"chargeBoxSerialNumber,omitempty"`
	ChargePointModel        string `json:"chargePointModel"`
	ChargePointSerialNumber string `json:"chargePointSerialNumber,omitempty"`
	ChargePointVendor       string `json:"chargePointVendor"`
	FirmwareVersion         string `json:"firmwareVersion,omitempty"`
	Iccid                   string `json:"iccid,omitempty"`
	Imsi                    string `json:"imsi,omitempty"`
	MeterSerialNumber       string `json:"meterSerialNumber,omitempty"`
	MeterType               string `json:"meterType,omitempty"`
}

// DataTransfer
type DataTransfer struct {
	chargepointRequest
	Data      string `json:"data,omitempty"`
	MessageId string `json:"messageId,omitempty"`
	VendorId  string `json:"vendorId"`
}

// DiagnosticsStatusNotification
type DiagnosticsStatusNotification struct {
	chargepointRequest
	Status string `json:"status"`
}

// FirmwareStatusNotification
type FirmwareStatusNotification struct {
	chargepointRequest
	Status string `json:"status"`
}

// Heartbeat
type Heartbeat struct {
	chargepointRequest
}

// MeterValueItems
type MeterValueItems struct {
	SampledValue []*SampledValueItems `json:"sampledValue"`
	Timestamp    time.Time            `json:"timestamp"`
}

// MeterValues
type MeterValues struct {
	chargepointRequest
	ConnectorId   int                `json:"connectorId"`
	MeterValue    []*MeterValueItems `json:"meterValue"`
	TransactionId int                `json:"transactionId,omitempty"`
}

// SampledValueItems
type SampledValueItems struct {
	Context   string `json:"context,omitempty"`
	Format    string `json:"format,omitempty"`
	Location  string `json:"location,omitempty"`
	Measurand string `json:"measurand,omitempty"`
	Phase     string `json:"phase,omitempty"`
	Unit      string `json:"unit,omitempty"`
	Value     string `json:"value"`
}

// StartTransaction
type StartTransaction struct {
	chargepointRequest
	ConnectorId   int       `json:"connectorId"`
	IdTag         string    `json:"idTag"`
	MeterStart    int       `json:"meterStart"`
	ReservationId int       `json:"reservationId,omitempty"`
	Timestamp     time.Time `json:"timestamp"`
}

// StatusNotification
type StatusNotification struct {
	chargepointRequest
	ConnectorId     int        `json:"connectorId"`
	ErrorCode       string     `json:"errorCode"`
	Info            string     `json:"info,omitempty"`
	Status          string     `json:"status"`
	Timestamp       *time.Time `json:"timestamp,omitempty"`
	VendorErrorCode string     `json:"vendorErrorCode,omitempty"`
	VendorId        string     `json:"vendorId,omitempty"`
}

// StopTransaction
type StopTransaction struct {
	chargepointRequest
	IdTag           string                  `json:"idTag,omitempty"`
	MeterStop       int                     `json:"meterStop"`
	Reason          string                  `json:"reason,omitempty"`
	Timestamp       time.Time               `json:"timestamp"`
	TransactionData []*TransactionDataItems `json:"transactionData,omitempty"`
	TransactionId   int                     `json:"transactionId"`
}

// TransactionDataItems
type TransactionDataItems struct {
	SampledValue []*SampledValueItems `json:"sampledValue"`
	Timestamp    time.Time            `json:"timestamp"`
}

func (strct *Authorize) Action() string                     { return "Authorize" }
func (strct *BootNotification) Action() string              { return "BootNotification" }
func (strct *DataTransfer) Action() string                  { return "DataTransfer" }
func (strct *DiagnosticsStatusNotification) Action() string { return "DiagnosticsStatusNotification" }
func (strct *FirmwareStatusNotification) Action() string    { return "FirmwareStatusNotification" }
func (strct *Heartbeat) Action() string                     { return "Heartbeat" }
func (strct *MeterValues) Action() string                   { return "MeterValues" }
func (strct *StartTransaction) Action() string              { return "StartTransaction" }
func (strct *StatusNotification) Action() string            { return "StatusNotification" }
func (strct *StopTransaction) Action() string               { return "StopTransaction" }

func (strct *Authorize) GetResponseObject() cpres.ChargePointResponse { return &cpres.Authorize{} }
func (strct *BootNotification) GetResponseObject() cpres.ChargePointResponse {
	return &cpres.BootNotification{}
}
func (strct *DataTransfer) GetResponseObject() cpres.ChargePointResponse { return &cpres.DataTransfer{} }
func (strct *DiagnosticsStatusNotification) GetResponseObject() cpres.ChargePointResponse {
	return &cpres.DiagnosticsStatusNotification{}
}
func (strct *FirmwareStatusNotification) GetResponseObject() cpres.ChargePointResponse {
	return &cpres.FirmwareStatusNotification{}
}
func (strct *Heartbeat) GetResponseObject() cpres.ChargePointResponse   { return &cpres.Heartbeat{} }
func (strct *MeterValues) GetResponseObject() cpres.ChargePointResponse { return &cpres.MeterValues{} }
func (strct *StartTransaction) GetResponseObject() cpres.ChargePointResponse {
	return &cpres.StartTransaction{}
}
func (strct *StatusNotification) GetResponseObject() cpres.ChargePointResponse {
	return &cpres.StatusNotification{}
}
func (strct *StopTransaction) GetResponseObject() cpres.ChargePointResponse {
	return &cpres.StopTransaction{}
}
