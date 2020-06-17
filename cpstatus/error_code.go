package cpstatus

type ErrorCode string

const (
	// Failure to lock or unlock connector.
	ConnectorLockFailure ErrorCode = "ConnectorLockFailure"
	// Communication failure with the vehicle, might be Mode 3 or other communication protocol problem. This is not a real error in the sense that the Charge Point doesn’t need to go to the faulted state. Instead, it should go to the SuspendedEVSE state.
	EVCommunicationError ErrorCode = "EVCommunicationError"
	// Ground fault circuit interrupter has been activated.
	GroundFailure ErrorCode = "GroundFailure"
	// Temperature inside Charge Point is too high.
	HighTemperature ErrorCode = "HighTemperature"
	// Error in internal hard- or software component.
	InternalError ErrorCode = "InternalError"
	// The authorization information received from the Central System is in conflict with the LocalAuthorizationList.
	LocalListConflict ErrorCode = "LocalListConflict"
	// No error to report.
	NoError ErrorCode = "NoError"
	// Other type of error. More information in vendorErrorCode.
	OtherError ErrorCode = "OtherError"
	// Over current protection device has tripped.
	OverCurrentFailure ErrorCode = "OverCurrentFailure"
	// Voltage has risen above an acceptable level.
	OverVoltage ErrorCode = "OverVoltage"
	// Failure to read electrical/energy/power meter.
	PowerMeterFailure ErrorCode = "PowerMeterFailure"
	// Failure to control power switch.
	PowerSwitchFailure ErrorCode = "PowerSwitchFailure"
	// Failure with idTag reader.
	ReaderFailure ErrorCode = "ReaderFailure"
	// Unable to perform a reset.
	ResetFailure ErrorCode = "ResetFailure"
	// Voltage has dropped below an acceptable level.
	UnderVoltage ErrorCode = "UnderVoltage"
	// Wireless communication device reports a weak signal.
	WeakSignal ErrorCode = "WeakSignal"
)
