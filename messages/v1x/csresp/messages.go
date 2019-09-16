package csresp

import (
	"encoding/xml"
	"time"
)

// CancelReservation
type CancelReservation struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ cancelReservationResponse"`

	Status string `json:"status" xml:"status"`
}

// ChangeAvailability
type ChangeAvailability struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ changeAvailabilityResponse"`

	Status string `json:"status" xml:"status"`
}

// ChangeConfiguration
type ChangeConfiguration struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ changeConfigurationResponse"`

	Status string `json:"status" xml:"status"`
}

// ChargingSchedule
type ChargingSchedule struct {
	ChargingRateUnit       string                         `json:"chargingRateUnit" xml:"chargingRateUnit"`
	ChargingSchedulePeriod []*ChargingSchedulePeriodItems `json:"chargingSchedulePeriod" xml:"chargingSchedulePeriod"`
	Duration               int                            `json:"duration,omitempty" xml:"duration,omitempty"`
	MinChargingRate        float64                        `json:"minChargingRate,omitempty" xml:"minChargingRate,omitempty"`
	StartSchedule          *time.Time                     `json:"startSchedule,omitempty" xml:"startSchedule,omitempty"`
}

// ChargingSchedulePeriodItems
type ChargingSchedulePeriodItems struct {
	Limit        float64 `json:"limit" xml:"limit"`
	NumberPhases int     `json:"numberPhases,omitempty" xml:"numberPhases,omitempty"`
	StartPeriod  int     `json:"startPeriod" xml:"startPeriod"`
}

// ClearCache
type ClearCache struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ clearCacheResponse"`

	Status string `json:"status" xml:"status"`
}

// ClearChargingProfile
type ClearChargingProfile struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ clearChargingProfileResponse"`

	Status string `json:"status" xml:"status"`
}

// ConfigurationKeyItems
type ConfigurationKeyItems struct {
	Key      string `json:"key" xml:"key"`
	Readonly bool   `json:"readonly" xml:"readonly"`
	Value    string `json:"value,omitempty" xml:"value,omitempty"`
}

// DataTransfer
type DataTransfer struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ dataTransferResponse"`

	Data   string `json:"data,omitempty" xml:"data,omitempty"`
	Status string `json:"status" xml:"status"`
}

// GetCompositeSchedule
type GetCompositeSchedule struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ getCompositeScheduleResponse"`

	ChargingSchedule *ChargingSchedule `json:"chargingSchedule,omitempty" xml:"chargingSchedule,omitempty"`
	ConnectorId      int               `json:"connectorId,omitempty" xml:"connectorId,omitempty"`
	ScheduleStart    *time.Time        `json:"scheduleStart,omitempty" xml:"scheduleStart,omitempty"`
	Status           string            `json:"status" xml:"status"`
}

// GetConfiguration
type GetConfiguration struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ getConfigurationResponse"`

	ConfigurationKey []*ConfigurationKeyItems `json:"configurationKey,omitempty" xml:"configurationKey,omitempty"`
	UnknownKey       []string                 `json:"unknownKey,omitempty" xml:"unknownKey,omitempty"`
}

// GetDiagnostics
type GetDiagnostics struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ getDiagnosticsResponse"`

	FileName string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

// GetLocalListVersion
type GetLocalListVersion struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ getLocalListVersionResponse"`

	ListVersion int `json:"listVersion" xml:"listVersion"`
}

// RemoteStartTransaction
type RemoteStartTransaction struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ remoteStartTransactionResponse"`

	Status string `json:"status" xml:"status"`
}

// RemoteStopTransaction
type RemoteStopTransaction struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ remoteStopTransactionResponse"`

	Status string `json:"status" xml:"status"`
}

// ReserveNow
type ReserveNow struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ reserveNowResponse"`

	Status string `json:"status" xml:"status"`
}

// Reset
type Reset struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ resetResponse"`

	Status string `json:"status" xml:"status"`
}

// SendLocalList
type SendLocalList struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ sendLocalListResponse"`

	Status string `json:"status" xml:"status"`
}

// TriggerMessage
type TriggerMessage struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ triggerMessageResponse"`

	Status string `json:"status" xml:"status"`
}

// UnlockConnector
type UnlockConnector struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ unlockConnectorResponse"`

	Status string `json:"status" xml:"status"`
}

// UpdateFirmware
type UpdateFirmware struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ updateFirmwareResponse"`
}

// SetChargingProfile
type SetChargingProfile struct {
	centralSystemResponse
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ setChargingProfileResponse"`

	Status string `json:"status" xml:"status"`
}
