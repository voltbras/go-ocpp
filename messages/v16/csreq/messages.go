package csreq

import (
	"time"

	"github.com/eduhenke/go-ocpp/messages/v16/csres"
)

// CancelReservation
type CancelReservation struct {
	centralSystemRequest
	ReservationId int `json:"reservationId"`
}

// ChangeAvailability
type ChangeAvailability struct {
	centralSystemRequest
	ConnectorId int    `json:"connectorId"`
	Type        string `json:"type"`
}

// ChangeConfiguration
type ChangeConfiguration struct {
	centralSystemRequest
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ChargingProfile
type ChargingProfile struct {
	ChargingProfileId      int               `json:"chargingProfileId"`
	ChargingProfileKind    string            `json:"chargingProfileKind"`
	ChargingProfilePurpose string            `json:"chargingProfilePurpose"`
	ChargingSchedule       *ChargingSchedule `json:"chargingSchedule"`
	RecurrencyKind         string            `json:"recurrencyKind,omitempty"`
	StackLevel             int               `json:"stackLevel"`
	TransactionId          int               `json:"transactionId,omitempty"`
	ValidFrom              *time.Time        `json:"validFrom,omitempty"`
	ValidTo                *time.Time        `json:"validTo,omitempty"`
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
	centralSystemRequest
}

// ClearChargingProfile
type ClearChargingProfile struct {
	centralSystemRequest
	ChargingProfilePurpose string `json:"chargingProfilePurpose,omitempty"`
	ConnectorId            int    `json:"connectorId,omitempty"`
	Id                     int    `json:"id,omitempty"`
	StackLevel             int    `json:"stackLevel,omitempty"`
}

// DataTransfer
type DataTransfer struct {
	centralSystemRequest
	Data      string `json:"data,omitempty"`
	MessageId string `json:"messageId,omitempty"`
	VendorId  string `json:"vendorId"`
}

// GetCompositeSchedule
type GetCompositeSchedule struct {
	centralSystemRequest
	ChargingRateUnit string `json:"chargingRateUnit,omitempty"`
	ConnectorId      int    `json:"connectorId"`
	Duration         int    `json:"duration"`
}

// GetConfiguration
type GetConfiguration struct {
	centralSystemRequest
	Key []string `json:"key,omitempty"`
}

// GetDiagnostics
type GetDiagnostics struct {
	centralSystemRequest
	Location      string     `json:"location"`
	Retries       int        `json:"retries,omitempty"`
	RetryInterval int        `json:"retryInterval,omitempty"`
	StartTime     *time.Time `json:"startTime,omitempty"`
	StopTime      *time.Time `json:"stopTime,omitempty"`
}

// GetLocalListVersion
type GetLocalListVersion struct {
	centralSystemRequest
}

// IdTagInfo
type IdTagInfo struct {
	ExpiryDate  *time.Time `json:"expiryDate,omitempty"`
	ParentIdTag string     `json:"parentIdTag,omitempty"`
	Status      string     `json:"status"`
}

// LocalAuthorizationListItems
type LocalAuthorizationListItems struct {
	IdTag     string     `json:"idTag"`
	IdTagInfo *IdTagInfo `json:"idTagInfo,omitempty"`
}

// RemoteStartTransaction
type RemoteStartTransaction struct {
	centralSystemRequest
	ChargingProfile *ChargingProfile `json:"chargingProfile,omitempty"`
	ConnectorId     int              `json:"connectorId,omitempty"`
	IdTag           string           `json:"idTag"`
}

// RemoteStopTransaction
type RemoteStopTransaction struct {
	centralSystemRequest
	TransactionId int `json:"transactionId"`
}

// ReserveNow
type ReserveNow struct {
	centralSystemRequest
	ConnectorId   int       `json:"connectorId"`
	ExpiryDate    time.Time `json:"expiryDate"`
	IdTag         string    `json:"idTag"`
	ParentIdTag   string    `json:"parentIdTag,omitempty"`
	ReservationId int       `json:"reservationId"`
}

// Reset
type Reset struct {
	centralSystemRequest
	Type string `json:"type"`
}

// SendLocalList
type SendLocalList struct {
	centralSystemRequest
	ListVersion            int                            `json:"listVersion"`
	LocalAuthorizationList []*LocalAuthorizationListItems `json:"localAuthorizationList,omitempty"`
	UpdateType             string                         `json:"updateType"`
}

// TriggerMessage
type TriggerMessage struct {
	centralSystemRequest
	ConnectorId      int    `json:"connectorId,omitempty"`
	RequestedMessage string `json:"requestedMessage"`
}

// UnlockConnector
type UnlockConnector struct {
	centralSystemRequest
	ConnectorId int `json:"connectorId"`
}

// UpdateFirmware
type UpdateFirmware struct {
	centralSystemRequest
	Location      string    `json:"location"`
	Retries       float64   `json:"retries,omitempty"`
	RetrieveDate  time.Time `json:"retrieveDate"`
	RetryInterval float64   `json:"retryInterval,omitempty"`
}

// CsChargingProfiles
type CsChargingProfiles struct {
	ChargingProfileId      int               `json:"chargingProfileId"`
	ChargingProfileKind    string            `json:"chargingProfileKind"`
	ChargingProfilePurpose string            `json:"chargingProfilePurpose"`
	ChargingSchedule       *ChargingSchedule `json:"chargingSchedule"`
	RecurrencyKind         string            `json:"recurrencyKind,omitempty"`
	StackLevel             int               `json:"stackLevel"`
	TransactionId          int               `json:"transactionId,omitempty"`
	ValidFrom              time.Time         `json:"validFrom,omitempty"`
	ValidTo                time.Time         `json:"validTo,omitempty"`
}

// SetChargingProfile
type SetChargingProfile struct {
	centralSystemRequest
	ConnectorId        int                 `json:"connectorId"`
	CsChargingProfiles *CsChargingProfiles `json:"csChargingProfiles"`
}

func (strct *CancelReservation) Action() string      { return "CancelReservation" }
func (strct *ChangeAvailability) Action() string     { return "ChangeAvailability" }
func (strct *ChangeConfiguration) Action() string    { return "ChangeConfiguration" }
func (strct *ClearCache) Action() string             { return "ClearCache" }
func (strct *ClearChargingProfile) Action() string   { return "ClearChargingProfile" }
func (strct *DataTransfer) Action() string           { return "DataTransfer" }
func (strct *GetCompositeSchedule) Action() string   { return "GetCompositeSchedule" }
func (strct *GetConfiguration) Action() string       { return "GetConfiguration" }
func (strct *GetDiagnostics) Action() string         { return "GetDiagnostics" }
func (strct *GetLocalListVersion) Action() string    { return "GetLocalListVersion" }
func (strct *RemoteStartTransaction) Action() string { return "RemoteStartTransaction" }
func (strct *RemoteStopTransaction) Action() string  { return "RemoteStopTransaction" }
func (strct *ReserveNow) Action() string             { return "ReserveNow" }
func (strct *Reset) Action() string                  { return "Reset" }
func (strct *SendLocalList) Action() string          { return "SendLocalList" }
func (strct *TriggerMessage) Action() string         { return "TriggerMessage" }
func (strct *UnlockConnector) Action() string        { return "UnlockConnector" }
func (strct *UpdateFirmware) Action() string         { return "UpdateFirmware" }
func (strct *SetChargingProfile) Action() string     { return "SetChargingProfile" }

func (strct *CancelReservation) GetResponseObject() csres.CentralSystemResponse {
	return &csres.CancelReservation{}
}
func (strct *ChangeAvailability) GetResponseObject() csres.CentralSystemResponse {
	return &csres.ChangeAvailability{}
}
func (strct *ChangeConfiguration) GetResponseObject() csres.CentralSystemResponse {
	return &csres.ChangeConfiguration{}
}
func (strct *ClearCache) GetResponseObject() csres.CentralSystemResponse { return &csres.ClearCache{} }
func (strct *ClearChargingProfile) GetResponseObject() csres.CentralSystemResponse {
	return &csres.ClearChargingProfile{}
}
func (strct *DataTransfer) GetResponseObject() csres.CentralSystemResponse {
	return &csres.DataTransfer{}
}
func (strct *GetCompositeSchedule) GetResponseObject() csres.CentralSystemResponse {
	return &csres.GetCompositeSchedule{}
}
func (strct *GetConfiguration) GetResponseObject() csres.CentralSystemResponse {
	return &csres.GetConfiguration{}
}
func (strct *GetDiagnostics) GetResponseObject() csres.CentralSystemResponse {
	return &csres.GetDiagnostics{}
}
func (strct *GetLocalListVersion) GetResponseObject() csres.CentralSystemResponse {
	return &csres.GetLocalListVersion{}
}
func (strct *RemoteStartTransaction) GetResponseObject() csres.CentralSystemResponse {
	return &csres.RemoteStartTransaction{}
}
func (strct *RemoteStopTransaction) GetResponseObject() csres.CentralSystemResponse {
	return &csres.RemoteStopTransaction{}
}
func (strct *ReserveNow) GetResponseObject() csres.CentralSystemResponse { return &csres.ReserveNow{} }
func (strct *Reset) GetResponseObject() csres.CentralSystemResponse      { return &csres.Reset{} }
func (strct *SendLocalList) GetResponseObject() csres.CentralSystemResponse {
	return &csres.SendLocalList{}
}
func (strct *TriggerMessage) GetResponseObject() csres.CentralSystemResponse {
	return &csres.TriggerMessage{}
}
func (strct *UnlockConnector) GetResponseObject() csres.CentralSystemResponse {
	return &csres.UnlockConnector{}
}
func (strct *UpdateFirmware) GetResponseObject() csres.CentralSystemResponse {
	return &csres.UpdateFirmware{}
}
func (strct *SetChargingProfile) GetResponseObject() csres.CentralSystemResponse {
	return &csres.SetChargingProfile{}
}
