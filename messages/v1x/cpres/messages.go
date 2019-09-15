package cpres

import (
	"encoding/xml"
	"time"
)

// Authorize
type Authorize struct {
	chargepointResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ authorizeResponse"`

	IdTagInfo *IdTagInfo `json:"idTagInfo" xml:"idTagInfo,omitempty"`
}

// BootNotification
type BootNotification struct {
	chargepointResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ bootNotificationResponse"`

	CurrentTime time.Time `json:"currentTime" xml:"currentTime,omitempty"`

	// Interval in 1.5 is heartbeatInterval
	Interval float64 `json:"interval" xml:"heartbeatInterval,omitempty"`
	Status   string  `json:"status" xml:"status,omitempty"`
}

// DataTransfer
type DataTransfer struct {
	chargepointResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ dataTransferResponse"`

	Data   string `json:"data, xml:"data,omitempty"omitempty"`
	Status string `json:"status" xml:"status,omitempty"`
}

// DiagnosticsStatusNotification
type DiagnosticsStatusNotification struct {
	chargepointResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ diagnosticsStatusNotificationResponse"`
}

// FirmwareStatusNotification
type FirmwareStatusNotification struct {
	chargepointResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ firmwareStatusNotificationResponse"`
}

// Heartbeat
type Heartbeat struct {
	chargepointResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ heartbeatResponse"`

	CurrentTime time.Time `json:"currentTime" xml:"currentTime,omitempty"`
}

// IdTagInfo
type IdTagInfo struct {
	ExpiryDate  *time.Time `json:"expiryDate,omitempty" xml:"expiryDate,omitempty"`
	ParentIdTag string     `json:"parentIdTag,omitempty" xml:"parentIdTag,omitempty"`
	Status      string     `json:"status" xml:"status,omitempty"`
}

// MeterValues
type MeterValues struct {
	chargepointResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ meterValuesResponse"`
}

// StartTransaction
type StartTransaction struct {
	chargepointResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ startTransactionResponse"`

	IdTagInfo     *IdTagInfo `json:"idTagInfo" xml:"idTagInfo,omitempty"`
	TransactionId int32      `json:"transactionId" xml:"transactionId,omitempty"`
}

// StatusNotification
type StatusNotification struct {
	chargepointResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ statusNotificationResponse"`
}

// StopTransaction
type StopTransaction struct {
	chargepointResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cs/2012/06/ stopTransactionResponse"`

	IdTagInfo *IdTagInfo `json:"idTagInfo, xml:"idTagInfo,omitempty"omitempty"`
}
