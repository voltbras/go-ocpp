package cpres

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// Authorize
type Authorize struct {
	chargepointResponse
	IdTagInfo *IdTagInfo `json:"idTagInfo"`
}

// BootNotification
type BootNotification struct {
	chargepointResponse
	CurrentTime time.Time `json:"currentTime"`
	Interval    float64   `json:"interval"`
	Status      string    `json:"status"`
}

// DataTransfer
type DataTransfer struct {
	chargepointResponse
	Data   string `json:"data,omitempty"`
	Status string `json:"status"`
}

// DiagnosticsStatusNotification
type DiagnosticsStatusNotification struct {
	chargepointResponse
}

// FirmwareStatusNotification
type FirmwareStatusNotification struct {
	chargepointResponse
}

// Heartbeat
type Heartbeat struct {
	chargepointResponse
	CurrentTime time.Time `json:"currentTime"`
}

// IdTagInfo
type IdTagInfo struct {
	ExpiryDate  *time.Time `json:"expiryDate,omitempty"`
	ParentIdTag string     `json:"parentIdTag,omitempty"`
	Status      string     `json:"status"`
}

// MeterValues
type MeterValues struct {
	chargepointResponse
}

// StartTransaction
type StartTransaction struct {
	chargepointResponse
	IdTagInfo     *IdTagInfo `json:"idTagInfo"`
	TransactionId int        `json:"transactionId"`
}

// StatusNotification
type StatusNotification struct {
	chargepointResponse
}

// StopTransaction
type StopTransaction struct {
	chargepointResponse
	IdTagInfo *IdTagInfo `json:"idTagInfo,omitempty"`
}
