package csreq

import (
	"time"

	"github.com/eduhenke/go-ocpp/messages"
	"github.com/eduhenke/go-ocpp/messages/v1x/csres"
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
	ConnectorId     int32            `json:"connectorId,omitempty"`
	IdTag           string           `json:"idTag"`
}

// RemoteStopTransaction
type RemoteStopTransaction struct {
	centralSystemRequest
	TransactionId int32 `json:"transactionId"`
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

func (m *CancelReservation) Action() string      { return "CancelReservation" }
func (m *ChangeAvailability) Action() string     { return "ChangeAvailability" }
func (m *ChangeConfiguration) Action() string    { return "ChangeConfiguration" }
func (m *ClearCache) Action() string             { return "ClearCache" }
func (m *ClearChargingProfile) Action() string   { return "ClearChargingProfile" }
func (m *DataTransfer) Action() string           { return "DataTransfer" }
func (m *GetCompositeSchedule) Action() string   { return "GetCompositeSchedule" }
func (m *GetConfiguration) Action() string       { return "GetConfiguration" }
func (m *GetDiagnostics) Action() string         { return "GetDiagnostics" }
func (m *GetLocalListVersion) Action() string    { return "GetLocalListVersion" }
func (m *RemoteStartTransaction) Action() string { return "RemoteStartTransaction" }
func (m *RemoteStopTransaction) Action() string  { return "RemoteStopTransaction" }
func (m *ReserveNow) Action() string             { return "ReserveNow" }
func (m *Reset) Action() string                  { return "Reset" }
func (m *SendLocalList) Action() string          { return "SendLocalList" }
func (m *TriggerMessage) Action() string         { return "TriggerMessage" }
func (m *UnlockConnector) Action() string        { return "UnlockConnector" }
func (m *UpdateFirmware) Action() string         { return "UpdateFirmware" }
func (m *SetChargingProfile) Action() string     { return "SetChargingProfile" }

func (m *CancelReservation) GetCentralSystemResponse() csres.CentralSystemResponse {
	return &csres.CancelReservation{}
}
func (m *ChangeAvailability) GetCentralSystemResponse() csres.CentralSystemResponse {
	return &csres.ChangeAvailability{}
}
func (m *ChangeConfiguration) GetCentralSystemResponse() csres.CentralSystemResponse {
	return &csres.ChangeConfiguration{}
}
func (m *ClearCache) GetCentralSystemResponse() csres.CentralSystemResponse {
	return &csres.ClearCache{}
}
func (m *ClearChargingProfile) GetCentralSystemResponse() csres.CentralSystemResponse {
	return &csres.ClearChargingProfile{}
}
func (m *DataTransfer) GetCentralSystemResponse() csres.CentralSystemResponse {
	return &csres.DataTransfer{}
}
func (m *GetCompositeSchedule) GetCentralSystemResponse() csres.CentralSystemResponse {
	return &csres.GetCompositeSchedule{}
}
func (m *GetConfiguration) GetCentralSystemResponse() csres.CentralSystemResponse {
	return &csres.GetConfiguration{}
}
func (m *GetDiagnostics) GetCentralSystemResponse() csres.CentralSystemResponse {
	return &csres.GetDiagnostics{}
}
func (m *GetLocalListVersion) GetCentralSystemResponse() csres.CentralSystemResponse {
	return &csres.GetLocalListVersion{}
}
func (m *RemoteStartTransaction) GetCentralSystemResponse() csres.CentralSystemResponse {
	return &csres.RemoteStartTransaction{}
}
func (m *RemoteStopTransaction) GetCentralSystemResponse() csres.CentralSystemResponse {
	return &csres.RemoteStopTransaction{}
}
func (m *ReserveNow) GetCentralSystemResponse() csres.CentralSystemResponse {
	return &csres.ReserveNow{}
}
func (m *Reset) GetCentralSystemResponse() csres.CentralSystemResponse { return &csres.Reset{} }
func (m *SendLocalList) GetCentralSystemResponse() csres.CentralSystemResponse {
	return &csres.SendLocalList{}
}
func (m *TriggerMessage) GetCentralSystemResponse() csres.CentralSystemResponse {
	return &csres.TriggerMessage{}
}
func (m *UnlockConnector) GetCentralSystemResponse() csres.CentralSystemResponse {
	return &csres.UnlockConnector{}
}
func (m *UpdateFirmware) GetCentralSystemResponse() csres.CentralSystemResponse {
	return &csres.UpdateFirmware{}
}
func (m *SetChargingProfile) GetCentralSystemResponse() csres.CentralSystemResponse {
	return &csres.SetChargingProfile{}
}

func (m *CancelReservation) GetResponse() messages.Response {

	return m.GetCentralSystemResponse()
}
func (m *ChangeAvailability) GetResponse() messages.Response {
	return m.GetCentralSystemResponse()
}
func (m *ChangeConfiguration) GetResponse() messages.Response {
	return m.GetCentralSystemResponse()
}
func (m *ClearCache) GetResponse() messages.Response {
	return m.GetCentralSystemResponse()
}
func (m *ClearChargingProfile) GetResponse() messages.Response {
	return m.GetCentralSystemResponse()
}
func (m *DataTransfer) GetResponse() messages.Response {
	return m.GetCentralSystemResponse()
}
func (m *GetCompositeSchedule) GetResponse() messages.Response {
	return m.GetCentralSystemResponse()
}
func (m *GetConfiguration) GetResponse() messages.Response {
	return m.GetCentralSystemResponse()
}
func (m *GetDiagnostics) GetResponse() messages.Response {
	return m.GetCentralSystemResponse()
}
func (m *GetLocalListVersion) GetResponse() messages.Response {
	return m.GetCentralSystemResponse()
}
func (m *RemoteStartTransaction) GetResponse() messages.Response {
	return m.GetCentralSystemResponse()
}
func (m *RemoteStopTransaction) GetResponse() messages.Response {
	return m.GetCentralSystemResponse()
}
func (m *ReserveNow) GetResponse() messages.Response {
	return m.GetCentralSystemResponse()
}
func (m *Reset) GetResponse() messages.Response {
	return m.GetCentralSystemResponse()
}
func (m *SendLocalList) GetResponse() messages.Response {
	return m.GetCentralSystemResponse()
}
func (m *TriggerMessage) GetResponse() messages.Response {
	return m.GetCentralSystemResponse()
}
func (m *UnlockConnector) GetResponse() messages.Response {
	return m.GetCentralSystemResponse()
}
func (m *UpdateFirmware) GetResponse() messages.Response {
	return m.GetCentralSystemResponse()
}
func (m *SetChargingProfile) GetResponse() messages.Response {
	return m.GetCentralSystemResponse()
}
