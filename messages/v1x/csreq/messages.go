package csreq

import (
	"encoding/xml"
	"time"

	"github.com/eduhenke/go-ocpp/messages"
	"github.com/eduhenke/go-ocpp/messages/v1x/csresp"
)

// CancelReservation
type CancelReservation struct {
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ cancelReservationRequest"`

	ReservationId int `json:"reservationId" xml:"reservationId"`
}

// ChangeAvailability
type ChangeAvailability struct {
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ changeAvailabilityRequest"`

	ConnectorId int    `json:"connectorId" xml:"connectorId"`
	Type        string `json:"type" xml:"type"`
}

// ChangeConfiguration
type ChangeConfiguration struct {
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ changeConfigurationRequest"`

	Key   string `json:"key" xml:"key"`
	Value string `json:"value" xml:"value"`
}

// ChargingProfile
type ChargingProfile struct {
	ChargingProfileId      int               `json:"chargingProfileId" xml:"chargingProfileId"`
	ChargingProfileKind    string            `json:"chargingProfileKind" xml:"chargingProfileKind"`
	ChargingProfilePurpose string            `json:"chargingProfilePurpose" xml:"chargingProfilePurpose"`
	ChargingSchedule       *ChargingSchedule `json:"chargingSchedule" xml:"chargingSchedule"`
	RecurrencyKind         string            `json:"recurrencyKind,omitempty" xml:"recurrencyKind,omitempty"`
	StackLevel             int               `json:"stackLevel" xml:"stackLevel"`
	TransactionId          int               `json:"transactionId,omitempty" xml:"transactionId,omitempty"`
	ValidFrom              *time.Time        `json:"validFrom,omitempty" xml:"validFrom,omitempty"`
	ValidTo                *time.Time        `json:"validTo,omitempty" xml:"validTo,omitempty"`
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
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ clearCacheRequest"`
}

// ClearChargingProfile
type ClearChargingProfile struct {
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ clearChargingProfileRequest"`

	ChargingProfilePurpose string `json:"chargingProfilePurpose,omitempty" xml:"chargingProfilePurpose,omitempty"`
	ConnectorId            int    `json:"connectorId,omitempty" xml:"connectorId,omitempty"`
	Id                     int    `json:"id,omitempty" xml:"id,omitempty"`
	StackLevel             int    `json:"stackLevel,omitempty" xml:"stackLevel,omitempty"`
}

// DataTransfer
type DataTransfer struct {
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ dataTransferRequest"`

	Data      string `json:"data,omitempty" xml:"data,omitempty"`
	MessageId string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	VendorId  string `json:"vendorId" xml:"vendorId"`
}

// GetCompositeSchedule
type GetCompositeSchedule struct {
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ getCompositeScheduleRequest"`

	ChargingRateUnit string `json:"chargingRateUnit,omitempty" xml:"chargingRateUnit,omitempty"`
	ConnectorId      int    `json:"connectorId" xml:"connectorId"`
	Duration         int    `json:"duration" xml:"duration"`
}

// GetConfiguration
type GetConfiguration struct {
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ getConfigurationRequest"`

	Key []string `json:"key,omitempty" xml:"key,omitempty"`
}

// GetDiagnostics
type GetDiagnostics struct {
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ getDiagnosticsRequest"`

	Location      string     `json:"location" xml:"location"`
	Retries       int        `json:"retries,omitempty" xml:"retries,omitempty"`
	RetryInterval int        `json:"retryInterval,omitempty" xml:"retryInterval,omitempty"`
	StartTime     *time.Time `json:"startTime,omitempty" xml:"startTime,omitempty"`
	StopTime      *time.Time `json:"stopTime,omitempty" xml:"stopTime,omitempty"`
}

// GetLocalListVersion
type GetLocalListVersion struct {
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ getLocalListVersionRequest"`
}

// IdTagInfo
type IdTagInfo struct {
	ExpiryDate  *time.Time `json:"expiryDate,omitempty" xml:"expiryDate,omitempty"`
	ParentIdTag string     `json:"parentIdTag,omitempty" xml:"parentIdTag,omitempty"`
	Status      string     `json:"status" xml:"status"`
}

// LocalAuthorizationListItems
type LocalAuthorizationListItems struct {
	IdTag     string     `json:"idTag" xml:"idTag"`
	IdTagInfo *IdTagInfo `json:"idTagInfo,omitempty" xml:"idTagInfo,omitempty"`
}

// RemoteStartTransaction
type RemoteStartTransaction struct {
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ remoteStartTransactionRequest"`

	IdTag           string           `json:"idTag" xml:"idTag"`
	ChargingProfile *ChargingProfile `json:"chargingProfile,omitempty" xml:"chargingProfile,omitempty"`
	ConnectorId     int32            `json:"connectorId,omitempty" xml:"connectorId,omitempty"`
}

// RemoteStopTransaction
type RemoteStopTransaction struct {
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ remoteStopTransactionRequest"`

	TransactionId int32 `json:"transactionId" xml:"transactionId"`
}

// ReserveNow
type ReserveNow struct {
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ reserveNowRequest"`

	ConnectorId   int       `json:"connectorId" xml:"connectorId"`
	ExpiryDate    time.Time `json:"expiryDate" xml:"expiryDate"`
	IdTag         string    `json:"idTag" xml:"idTag"`
	ParentIdTag   string    `json:"parentIdTag,omitempty" xml:"parentIdTag,omitempty"`
	ReservationId int       `json:"reservationId" xml:"reservationId"`
}

// Reset
type Reset struct {
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ resetRequest"`

	Type string `json:"type" xml:"type"`
}

// SendLocalList
type SendLocalList struct {
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ sendLocalListRequest"`

	ListVersion            int                            `json:"listVersion" xml:"listVersion"`
	LocalAuthorizationList []*LocalAuthorizationListItems `json:"localAuthorizationList,omitempty" xml:"localAuthorizationList,omitempty"`
	UpdateType             string                         `json:"updateType" xml:"updateType"`
}

// TriggerMessage
type TriggerMessage struct {
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ triggerMessageRequest"`

	ConnectorId      int    `json:"connectorId,omitempty" xml:"connectorId,omitempty"`
	RequestedMessage string `json:"requestedMessage" xml:"requestedMessage"`
}

// UnlockConnector
type UnlockConnector struct {
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ unlockConnectorRequest"`

	ConnectorId int `json:"connectorId" xml:"connectorId"`
}

// UpdateFirmware
type UpdateFirmware struct {
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ updateFirmwareRequest"`

	Location      string    `json:"location" xml:"location"`
	Retries       float64   `json:"retries,omitempty" xml:"retries,omitempty"`
	RetrieveDate  time.Time `json:"retrieveDate" xml:"retrieveDate"`
	RetryInterval float64   `json:"retryInterval,omitempty" xml:"retryInterval,omitempty"`
}

// CsChargingProfiles
type CsChargingProfiles struct {
	ChargingProfileId      int               `json:"chargingProfileId" xml:"chargingProfileId"`
	ChargingProfileKind    string            `json:"chargingProfileKind" xml:"chargingProfileKind"`
	ChargingProfilePurpose string            `json:"chargingProfilePurpose" xml:"chargingProfilePurpose"`
	ChargingSchedule       *ChargingSchedule `json:"chargingSchedule" xml:"chargingSchedule"`
	RecurrencyKind         string            `json:"recurrencyKind,omitempty" xml:"recurrencyKind,omitempty"`
	StackLevel             int               `json:"stackLevel" xml:"stackLevel"`
	TransactionId          int               `json:"transactionId,omitempty" xml:"transactionId,omitempty"`
	ValidFrom              time.Time         `json:"validFrom,omitempty" xml:"validFrom,omitempty"`
	ValidTo                time.Time         `json:"validTo,omitempty" xml:"validTo,omitempty"`
}

// SetChargingProfile
type SetChargingProfile struct {
	centralSystemRequest
	XMLName xml.Name `json:"-" xml:"urn://Ocpp/Cp/2012/06/ setChargingProfileRequest"`

	ConnectorId        int                 `json:"connectorId" xml:"connectorId"`
	CsChargingProfiles *CsChargingProfiles `json:"csChargingProfiles" xml:"csChargingProfiles"`
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

func (m *CancelReservation) GetResponse() messages.Response {
	return &csresp.CancelReservation{}
}
func (m *ChangeAvailability) GetResponse() messages.Response {
	return &csresp.ChangeAvailability{}
}
func (m *ChangeConfiguration) GetResponse() messages.Response {
	return &csresp.ChangeConfiguration{}
}
func (m *ClearCache) GetResponse() messages.Response {
	return &csresp.ClearCache{}
}
func (m *ClearChargingProfile) GetResponse() messages.Response {
	return &csresp.ClearChargingProfile{}
}
func (m *DataTransfer) GetResponse() messages.Response {
	return &csresp.DataTransfer{}
}
func (m *GetCompositeSchedule) GetResponse() messages.Response {
	return &csresp.GetCompositeSchedule{}
}
func (m *GetConfiguration) GetResponse() messages.Response {
	return &csresp.GetConfiguration{}
}
func (m *GetDiagnostics) GetResponse() messages.Response {
	return &csresp.GetDiagnostics{}
}
func (m *GetLocalListVersion) GetResponse() messages.Response {
	return &csresp.GetLocalListVersion{}
}
func (m *RemoteStartTransaction) GetResponse() messages.Response {
	return &csresp.RemoteStartTransaction{}
}
func (m *RemoteStopTransaction) GetResponse() messages.Response {
	return &csresp.RemoteStopTransaction{}
}
func (m *ReserveNow) GetResponse() messages.Response {
	return &csresp.ReserveNow{}
}
func (m *Reset) GetResponse() messages.Response {
	return &csresp.Reset{}
}
func (m *SendLocalList) GetResponse() messages.Response {
	return &csresp.SendLocalList{}
}
func (m *TriggerMessage) GetResponse() messages.Response {
	return &csresp.TriggerMessage{}
}
func (m *UnlockConnector) GetResponse() messages.Response {
	return &csresp.UnlockConnector{}
}
func (m *UpdateFirmware) GetResponse() messages.Response {
	return &csresp.UpdateFirmware{}
}
func (m *SetChargingProfile) GetResponse() messages.Response {
	return &csresp.SetChargingProfile{}
}
