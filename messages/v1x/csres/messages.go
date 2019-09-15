package csres

import (
	"time"
)

// CancelReservation
type CancelReservation struct {
	centralSystemResponse
	Status string `json:"status"`
}

// ChangeAvailability
type ChangeAvailability struct {
	centralSystemResponse
	Status string `json:"status"`
}

// ChangeConfiguration
type ChangeConfiguration struct {
	centralSystemResponse
	Status string `json:"status"`
}

// ChargingSchedule
type ChargingSchedule struct {
	ChargingRateUnit       string                         `json:"chargingRateUnit"`
	ChargingSchedulePeriod []*ChargingSchedulePeriodItems `json:"chargingSchedulePeriod"`
	Duration               int                            `json:"duration,omitempty"`
	MinChargingRate        float64                        `json:"minChargingRate,omitempty"`
	StartSchedule          *time.Time                     `json:"startSchedule,omitempty"`
}

// ChargingSchedulePeriodItems
type ChargingSchedulePeriodItems struct {
	Limit        float64 `json:"limit"`
	NumberPhases int     `json:"numberPhases,omitempty"`
	StartPeriod  int     `json:"startPeriod"`
}

// ClearCache
type ClearCache struct {
	centralSystemResponse
	Status string `json:"status"`
}

// ClearChargingProfile
type ClearChargingProfile struct {
	centralSystemResponse
	Status string `json:"status"`
}

// ConfigurationKeyItems
type ConfigurationKeyItems struct {
	Key      string `json:"key"`
	Readonly bool   `json:"readonly"`
	Value    string `json:"value,omitempty"`
}

// DataTransfer
type DataTransfer struct {
	centralSystemResponse
	Data   string `json:"data,omitempty"`
	Status string `json:"status"`
}

// GetCompositeSchedule
type GetCompositeSchedule struct {
	centralSystemResponse
	ChargingSchedule *ChargingSchedule `json:"chargingSchedule,omitempty"`
	ConnectorId      int               `json:"connectorId,omitempty"`
	ScheduleStart    *time.Time        `json:"scheduleStart,omitempty"`
	Status           string            `json:"status"`
}

// GetConfiguration
type GetConfiguration struct {
	centralSystemResponse
	ConfigurationKey []*ConfigurationKeyItems `json:"configurationKey,omitempty"`
	UnknownKey       []string                 `json:"unknownKey,omitempty"`
}

// GetDiagnostics
type GetDiagnostics struct {
	centralSystemResponse
	FileName string `json:"fileName,omitempty"`
}

// GetLocalListVersion
type GetLocalListVersion struct {
	centralSystemResponse
	ListVersion int `json:"listVersion"`
}

// RemoteStartTransaction
type RemoteStartTransaction struct {
	centralSystemResponse
	Status string `json:"status"`
}

// RemoteStopTransaction
type RemoteStopTransaction struct {
	centralSystemResponse
	Status string `json:"status"`
}

// ReserveNow
type ReserveNow struct {
	centralSystemResponse
	Status string `json:"status"`
}

// Reset
type Reset struct {
	centralSystemResponse
	Status string `json:"status"`
}

// SendLocalList
type SendLocalList struct {
	centralSystemResponse
	Status string `json:"status"`
}

// TriggerMessage
type TriggerMessage struct {
	centralSystemResponse
	Status string `json:"status"`
}

// UnlockConnector
type UnlockConnector struct {
	centralSystemResponse
	Status string `json:"status"`
}

// UpdateFirmware
type UpdateFirmware struct {
	centralSystemResponse
}

// SetChargingProfile
type SetChargingProfile struct {
	centralSystemResponse
	Status string `json:"status"`
}
