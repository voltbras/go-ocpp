package messages

import (
	"encoding/xml"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

// String type of max 25 chars that is to be treated as case insensitive.
type ChargeBoxSerialNumber string

// String type of max 20 chars that is to be treated as case insensitive.
type ChargePointModel string

// String type of max 25 chars that is to be treated as case insensitive.
type ChargePointSerialNumber string

// String type of max 20 chars that is to be treated as case insensitive.
type ChargePointVendor string

// String type of max 50 chars that is to be treated as case insensitive.
type FirmwareVersion string

// String type of max 20 chars that is to be treated as case insensitive.
type IccidString string

// String type of max 20 chars that is to be treated as case insensitive.
type ImsiString string

// String type of max 25 chars that is to be treated as case insensitive.
type MeterSerialNumber string

// String type of max 25 chars that is to be treated as case insensitive.
type MeterType string

type ReadingContext string

const (
	ReadingContextInterruptionBegin ReadingContext = "Interruption.Begin"

	ReadingContextInterruptionEnd ReadingContext = "Interruption.End"

	ReadingContextSampleClock ReadingContext = "Sample.Clock"

	ReadingContextSamplePeriodic ReadingContext = "Sample.Periodic"

	ReadingContextTransactionBegin ReadingContext = "Transaction.Begin"

	ReadingContextTransactionEnd ReadingContext = "Transaction.End"
)

type Measurand string

const (
	MeasurandEnergyActiveExportRegister Measurand = "Energy.Active.Export.Register"

	MeasurandEnergyActiveImportRegister Measurand = "Energy.Active.Import.Register"

	MeasurandEnergyReactiveExportRegister Measurand = "Energy.Reactive.Export.Register"

	MeasurandEnergyReactiveImportRegister Measurand = "Energy.Reactive.Import.Register"

	MeasurandEnergyActiveExportInterval Measurand = "Energy.Active.Export.Interval"

	MeasurandEnergyActiveImportInterval Measurand = "Energy.Active.Import.Interval"

	MeasurandEnergyReactiveExportInterval Measurand = "Energy.Reactive.Export.Interval"

	MeasurandEnergyReactiveImportInterval Measurand = "Energy.Reactive.Import.Interval"

	MeasurandPowerActiveExport Measurand = "Power.Active.Export"

	MeasurandPowerActiveImport Measurand = "Power.Active.Import"

	MeasurandPowerReactiveExport Measurand = "Power.Reactive.Export"

	MeasurandPowerReactiveImport Measurand = "Power.Reactive.Import"

	MeasurandCurrentExport Measurand = "Current.Export"

	MeasurandCurrentImport Measurand = "Current.Import"

	MeasurandVoltage Measurand = "Voltage"

	MeasurandTemperature Measurand = "Temperature"
)

type ValueFormat string

const (
	ValueFormatRaw ValueFormat = "Raw"

	ValueFormatSignedData ValueFormat = "SignedData"
)

type UnitOfMeasure string

const (
	UnitOfMeasureWh UnitOfMeasure = "Wh"

	UnitOfMeasureKWh UnitOfMeasure = "kWh"

	UnitOfMeasureVarh UnitOfMeasure = "varh"

	UnitOfMeasureKvarh UnitOfMeasure = "kvarh"

	UnitOfMeasureW UnitOfMeasure = "W"

	UnitOfMeasureKW UnitOfMeasure = "kW"

	UnitOfMeasureVar_ UnitOfMeasure = "var"

	UnitOfMeasureKvar UnitOfMeasure = "kvar"

	UnitOfMeasureAmp UnitOfMeasure = "Amp"

	UnitOfMeasureVolt UnitOfMeasure = "Volt"

	UnitOfMeasureCelsius UnitOfMeasure = "Celsius"
)

const (
	LocationInlet Location = "Inlet"

	LocationOutlet Location = "Outlet"

	LocationBody Location = "Body"
)

// Defines the registration-status-value
type RegistrationStatus string

const (
	RegistrationStatusAccepted RegistrationStatus = "Accepted"

	RegistrationStatusRejected RegistrationStatus = "Rejected"
)

// Defines the charge-point-error-value
type ChargePointErrorCode string

const (
	ChargePointErrorCodeConnectorLockFailure ChargePointErrorCode = "ConnectorLockFailure"

	ChargePointErrorCodeHighTemperature ChargePointErrorCode = "HighTemperature"

	ChargePointErrorCodeMode3Error ChargePointErrorCode = "Mode3Error"

	ChargePointErrorCodeNoError ChargePointErrorCode = "NoError"

	ChargePointErrorCodePowerMeterFailure ChargePointErrorCode = "PowerMeterFailure"

	ChargePointErrorCodePowerSwitchFailure ChargePointErrorCode = "PowerSwitchFailure"

	ChargePointErrorCodeReaderFailure ChargePointErrorCode = "ReaderFailure"

	ChargePointErrorCodeResetFailure ChargePointErrorCode = "ResetFailure"

	ChargePointErrorCodeGroundFailure ChargePointErrorCode = "GroundFailure"

	ChargePointErrorCodeOverCurrentFailure ChargePointErrorCode = "OverCurrentFailure"

	ChargePointErrorCodeUnderVoltage ChargePointErrorCode = "UnderVoltage"

	ChargePointErrorCodeWeakSignal ChargePointErrorCode = "WeakSignal"

	ChargePointErrorCodeOtherError ChargePointErrorCode = "OtherError"
)

// Defines the charge-point-status-value
type ChargePointStatus string

const (
	ChargePointStatusAvailable ChargePointStatus = "Available"

	ChargePointStatusOccupied ChargePointStatus = "Occupied"

	ChargePointStatusFaulted ChargePointStatus = "Faulted"

	ChargePointStatusUnavailable ChargePointStatus = "Unavailable"

	ChargePointStatusReserved ChargePointStatus = "Reserved"
)

// Defines the firmware-status-value
type FirmwareStatus string

const (
	FirmwareStatusDownloaded FirmwareStatus = "Downloaded"

	FirmwareStatusDownloadFailed FirmwareStatus = "DownloadFailed"

	FirmwareStatusInstallationFailed FirmwareStatus = "InstallationFailed"

	FirmwareStatusInstalled FirmwareStatus = "Installed"
)

// Defines the diagnostics-status-value
type DiagnosticsStatus string

const (
	DiagnosticsStatusUploaded DiagnosticsStatus = "Uploaded"

	DiagnosticsStatusUploadFailed DiagnosticsStatus = "UploadFailed"
)

type AuthorizeRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cs/2012/06/ authorizeRequest"`

	IdTag IdToken `xml:"idTag,omitempty"`
}

type AuthorizeResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cs/2012/06/ authorizeResponse"`

	IdTagInfo IdTagInfo `xml:"idTagInfo,omitempty"`
}

type StartTransactionRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cs/2012/06/ startTransactionRequest"`

	ConnectorId int32 `xml:"connectorId,omitempty"`

	IdTag IdToken `xml:"idTag,omitempty"`

	Timestamp time.Time `xml:"timestamp,omitempty"`

	MeterStart int32 `xml:"meterStart,omitempty"`

	ReservationId int32 `xml:"reservationId,omitempty"`
}

type StartTransactionResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cs/2012/06/ startTransactionResponse"`

	TransactionId int32 `xml:"transactionId,omitempty"`

	IdTagInfo IdTagInfo `xml:"idTagInfo,omitempty"`
}

type TransactionData struct {
	Values []MeterValue `xml:"values,omitempty"`
}

type StopTransactionRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cs/2012/06/ stopTransactionRequest"`

	TransactionId int32 `xml:"transactionId,omitempty"`

	IdTag IdToken `xml:"idTag,omitempty"`

	Timestamp time.Time `xml:"timestamp,omitempty"`

	MeterStop int32 `xml:"meterStop,omitempty"`

	TransactionData []TransactionData `xml:"transactionData,omitempty"`
}

type StopTransactionResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cs/2012/06/ stopTransactionResponse"`

	IdTagInfo *IdTagInfo `xml:"idTagInfo,omitempty"`
}

type HeartbeatRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cs/2012/06/ heartbeatRequest"`
}

type HeartbeatResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cs/2012/06/ heartbeatResponse"`

	CurrentTime time.Time `xml:"currentTime,omitempty"`
}
type SampledValue struct {
	Value string `xml:",chardata"`

	Context ReadingContext `xml:"context,attr,omitempty"`

	Format ValueFormat `xml:"format,attr,omitempty"`

	Measurand Measurand `xml:"measurand,attr,omitempty"`

	Location Location `xml:"location,attr,omitempty"`

	Unit UnitOfMeasure `xml:"unit,attr,omitempty"`
}

type MeterValue struct {
	Timestamp time.Time `xml:"timestamp,omitempty"`

	SampledValues []SampledValue`xml:"value,omitempty" bson:"sampledValues,omitempty" json:"sampledValues,omitempty"`
}

type MeterValuesRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cs/2012/06/ meterValuesRequest"`

	ConnectorId int32 `xml:"connectorId,omitempty"`

	TransactionId int32 `xml:"transactionId,omitempty"`

	Values []MeterValue `xml:"values,omitempty"`
}

type MeterValuesResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cs/2012/06/ meterValuesResponse"`
}

type BootNotificationRequest struct {
	XMLName xml.Name `xml:"bootNotificationRequest"`

	ChargePointVendor ChargePointVendor `xml:"chargePointVendor,omitempty"`

	ChargePointModel ChargePointModel `xml:"chargePointModel,omitempty"`

	ChargePointSerialNumber ChargePointSerialNumber `xml:"chargePointSerialNumber,omitempty"`

	ChargeBoxSerialNumber ChargeBoxSerialNumber `xml:"chargeBoxSerialNumber,omitempty"`

	FirmwareVersion FirmwareVersion `xml:"firmwareVersion,omitempty"`

	Iccid IccidString `xml:"iccid,omitempty"`

	Imsi ImsiString `xml:"imsi,omitempty"`

	MeterType MeterType `xml:"meterType,omitempty"`

	MeterSerialNumber MeterSerialNumber `xml:"meterSerialNumber,omitempty"`
}

type BootNotificationResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cs/2012/06/ bootNotificationResponse"`

	Status RegistrationStatus `xml:"status,omitempty"`

	CurrentTime time.Time `xml:"currentTime,omitempty"`

	HeartbeatInterval int32 `xml:"heartbeatInterval,omitempty"`
}

type StatusNotificationRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cs/2012/06/ statusNotificationRequest"`

	ConnectorId int32 `xml:"connectorId,omitempty"`

	Status ChargePointStatus `xml:"status,omitempty"`

	ErrorCode ChargePointErrorCode `xml:"errorCode,omitempty"`

	Info string `xml:"info,omitempty"`

	Timestamp time.Time `xml:"timestamp,omitempty"`

	VendorId string `xml:"vendorId,omitempty"`

	VendorErrorCode string `xml:"vendorErrorCode,omitempty"`
}

type StatusNotificationResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cs/2012/06/ statusNotificationResponse"`
}

type FirmwareStatusNotificationRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cs/2012/06/ firmwareStatusNotificationRequest"`

	Status FirmwareStatus `xml:"status,omitempty"`
}

type FirmwareStatusNotificationResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cs/2012/06/ firmwareStatusNotificationResponse"`
}

type DiagnosticsStatusNotificationRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cs/2012/06/ diagnosticsStatusNotificationRequest"`

	Status DiagnosticsStatus `xml:"status,omitempty"`
}

type DiagnosticsStatusNotificationResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cs/2012/06/ diagnosticsStatusNotificationResponse"`
}
