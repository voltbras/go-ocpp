package cpreq

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func (m *Authorize) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "IdTag" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "idTag" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"idTag\": ")
	if tmp, err := json.Marshal(m.IdTag); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *Authorize) UnmarshalJSON(b []byte) error {
	idTagReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "idTag":
			if err := json.Unmarshal([]byte(v), &m.IdTag); err != nil {
				return err
			}
			idTagReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if idTag (a required property) was received
	if !idTagReceived {
		return errors.New("\"idTag\" is required but was not present")
	}
	return nil
}

func (m *BootNotification) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "chargeBoxSerialNumber" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargeBoxSerialNumber\": ")
	if tmp, err := json.Marshal(m.ChargeBoxSerialNumber); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ChargePointModel" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "chargePointModel" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargePointModel\": ")
	if tmp, err := json.Marshal(m.ChargePointModel); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "chargePointSerialNumber" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargePointSerialNumber\": ")
	if tmp, err := json.Marshal(m.ChargePointSerialNumber); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ChargePointVendor" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "chargePointVendor" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargePointVendor\": ")
	if tmp, err := json.Marshal(m.ChargePointVendor); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "firmwareVersion" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"firmwareVersion\": ")
	if tmp, err := json.Marshal(m.FirmwareVersion); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "iccid" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"iccid\": ")
	if tmp, err := json.Marshal(m.Iccid); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "imsi" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"imsi\": ")
	if tmp, err := json.Marshal(m.Imsi); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "meterSerialNumber" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"meterSerialNumber\": ")
	if tmp, err := json.Marshal(m.MeterSerialNumber); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "meterType" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"meterType\": ")
	if tmp, err := json.Marshal(m.MeterType); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *BootNotification) UnmarshalJSON(b []byte) error {
	chargePointModelReceived := false
	chargePointVendorReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "chargeBoxSerialNumber":
			if err := json.Unmarshal([]byte(v), &m.ChargeBoxSerialNumber); err != nil {
				return err
			}
		case "chargePointModel":
			if err := json.Unmarshal([]byte(v), &m.ChargePointModel); err != nil {
				return err
			}
			chargePointModelReceived = true
		case "chargePointSerialNumber":
			if err := json.Unmarshal([]byte(v), &m.ChargePointSerialNumber); err != nil {
				return err
			}
		case "chargePointVendor":
			if err := json.Unmarshal([]byte(v), &m.ChargePointVendor); err != nil {
				return err
			}
			chargePointVendorReceived = true
		case "firmwareVersion":
			if err := json.Unmarshal([]byte(v), &m.FirmwareVersion); err != nil {
				return err
			}
		case "iccid":
			if err := json.Unmarshal([]byte(v), &m.Iccid); err != nil {
				return err
			}
		case "imsi":
			if err := json.Unmarshal([]byte(v), &m.Imsi); err != nil {
				return err
			}
		case "meterSerialNumber":
			if err := json.Unmarshal([]byte(v), &m.MeterSerialNumber); err != nil {
				return err
			}
		case "meterType":
			if err := json.Unmarshal([]byte(v), &m.MeterType); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if chargePointModel (a required property) was received
	if !chargePointModelReceived {
		return errors.New("\"chargePointModel\" is required but was not present")
	}
	// check if chargePointVendor (a required property) was received
	if !chargePointVendorReceived {
		return errors.New("\"chargePointVendor\" is required but was not present")
	}
	return nil
}

func (m *DataTransfer) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "data" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"data\": ")
	if tmp, err := json.Marshal(m.Data); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "messageId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"messageId\": ")
	if tmp, err := json.Marshal(m.MessageId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "VendorId" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "vendorId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"vendorId\": ")
	if tmp, err := json.Marshal(m.VendorId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *DataTransfer) UnmarshalJSON(b []byte) error {
	vendorIdReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "data":
			if err := json.Unmarshal([]byte(v), &m.Data); err != nil {
				return err
			}
		case "messageId":
			if err := json.Unmarshal([]byte(v), &m.MessageId); err != nil {
				return err
			}
		case "vendorId":
			if err := json.Unmarshal([]byte(v), &m.VendorId); err != nil {
				return err
			}
			vendorIdReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if vendorId (a required property) was received
	if !vendorIdReceived {
		return errors.New("\"vendorId\" is required but was not present")
	}
	return nil
}

func (m *DiagnosticsStatusNotification) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Status" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "status" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"status\": ")
	if tmp, err := json.Marshal(m.Status); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *DiagnosticsStatusNotification) UnmarshalJSON(b []byte) error {
	statusReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "status":
			if err := json.Unmarshal([]byte(v), &m.Status); err != nil {
				return err
			}
			statusReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if status (a required property) was received
	if !statusReceived {
		return errors.New("\"status\" is required but was not present")
	}
	return nil
}

func (m *FirmwareStatusNotification) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Status" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "status" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"status\": ")
	if tmp, err := json.Marshal(m.Status); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *FirmwareStatusNotification) UnmarshalJSON(b []byte) error {
	statusReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "status":
			if err := json.Unmarshal([]byte(v), &m.Status); err != nil {
				return err
			}
			statusReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if status (a required property) was received
	if !statusReceived {
		return errors.New("\"status\" is required but was not present")
	}
	return nil
}

func (m *Heartbeat) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *Heartbeat) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, _ := range jsonMap {
		switch k {
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (m *MeterValueItems) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "SampledValue" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "sampledValue" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"sampledValue\": ")
	if tmp, err := json.Marshal(m.SampledValues); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Timestamp" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "timestamp" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"timestamp\": ")
	if tmp, err := json.Marshal(m.Timestamp.Format(time.RFC3339)); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *MeterValueItems) UnmarshalJSON(b []byte) error {
	sampledValueReceived := false
	timestampReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "sampledValue":
			if err := json.Unmarshal([]byte(v), &m.SampledValues); err != nil {
				return err
			}
			sampledValueReceived = true
		case "timestamp":
			if err := json.Unmarshal([]byte(v), &m.Timestamp); err != nil {
				return err
			}
			timestampReceived = true
		}
	}
	// check if sampledValue (a required property) was received
	if !sampledValueReceived {
		return errors.New("\"sampledValue\" is required but was not present")
	}
	// check if timestamp (a required property) was received
	if !timestampReceived {
		return errors.New("\"timestamp\" is required but was not present")
	}
	return nil
}

func (m *MeterValues) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "ConnectorId" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "connectorId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"connectorId\": ")
	if tmp, err := json.Marshal(m.ConnectorId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "MeterValue" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "meterValue" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"meterValue\": ")
	if tmp, err := json.Marshal(m.MeterValue); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "transactionId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"transactionId\": ")
	if tmp, err := json.Marshal(m.TransactionId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *MeterValues) UnmarshalJSON(b []byte) error {
	connectorIdReceived := false
	meterValueReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "connectorId":
			if err := json.Unmarshal([]byte(v), &m.ConnectorId); err != nil {
				return err
			}
			connectorIdReceived = true
		case "meterValue":
			if err := json.Unmarshal([]byte(v), &m.MeterValue); err != nil {
				return err
			}
			meterValueReceived = true
		case "transactionId":
			if err := json.Unmarshal([]byte(v), &m.TransactionId); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if connectorId (a required property) was received
	if !connectorIdReceived {
		return errors.New("\"connectorId\" is required but was not present")
	}
	// check if meterValue (a required property) was received
	if !meterValueReceived {
		return errors.New("\"meterValue\" is required but was not present")
	}
	return nil
}

func (m *SampledValue) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "context" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"context\": ")
	if tmp, err := json.Marshal(m.Context); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "format" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"format\": ")
	if tmp, err := json.Marshal(m.Format); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "location" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"location\": ")
	if tmp, err := json.Marshal(m.Location); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "measurand" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"measurand\": ")
	if tmp, err := json.Marshal(m.Measurand); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "phase" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"phase\": ")
	if tmp, err := json.Marshal(m.Phase); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "unit" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"unit\": ")
	if tmp, err := json.Marshal(m.Unit); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Value" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "value" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"value\": ")
	if tmp, err := json.Marshal(m.Value); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *SampledValue) UnmarshalJSON(b []byte) error {
	valueReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "context":
			if err := json.Unmarshal([]byte(v), &m.Context); err != nil {
				return err
			}
		case "format":
			if err := json.Unmarshal([]byte(v), &m.Format); err != nil {
				return err
			}
		case "location":
			if err := json.Unmarshal([]byte(v), &m.Location); err != nil {
				return err
			}
		case "measurand":
			if err := json.Unmarshal([]byte(v), &m.Measurand); err != nil {
				return err
			}
		case "phase":
			if err := json.Unmarshal([]byte(v), &m.Phase); err != nil {
				return err
			}
		case "unit":
			if err := json.Unmarshal([]byte(v), &m.Unit); err != nil {
				return err
			}
		case "value":
			if err := json.Unmarshal([]byte(v), &m.Value); err != nil {
				return err
			}
			valueReceived = true
		}
	}
	// check if value (a required property) was received
	if !valueReceived {
		return errors.New("\"value\" is required but was not present")
	}
	return nil
}

func (m *StartTransaction) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "ConnectorId" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "connectorId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"connectorId\": ")
	if tmp, err := json.Marshal(m.ConnectorId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "IdTag" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "idTag" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"idTag\": ")
	if tmp, err := json.Marshal(m.IdTag); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "MeterStart" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "meterStart" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"meterStart\": ")
	if tmp, err := json.Marshal(m.MeterStart); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "reservationId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"reservationId\": ")
	if tmp, err := json.Marshal(m.ReservationId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Timestamp" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "timestamp" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"timestamp\": ")
	if tmp, err := json.Marshal(m.Timestamp.Format(time.RFC3339)); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *StartTransaction) UnmarshalJSON(b []byte) error {
	connectorIdReceived := false
	idTagReceived := false
	meterStartReceived := false
	timestampReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "connectorId":
			if err := json.Unmarshal([]byte(v), &m.ConnectorId); err != nil {
				return err
			}
			connectorIdReceived = true
		case "idTag":
			if err := json.Unmarshal([]byte(v), &m.IdTag); err != nil {
				return err
			}
			idTagReceived = true
		case "meterStart":
			if err := json.Unmarshal([]byte(v), &m.MeterStart); err != nil {
				return err
			}
			meterStartReceived = true
		case "reservationId":
			if err := json.Unmarshal([]byte(v), &m.ReservationId); err != nil {
				return err
			}
		case "timestamp":
			if err := json.Unmarshal([]byte(v), &m.Timestamp); err != nil {
				return err
			}
			timestampReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if connectorId (a required property) was received
	if !connectorIdReceived {
		return errors.New("\"connectorId\" is required but was not present")
	}
	// check if idTag (a required property) was received
	if !idTagReceived {
		return errors.New("\"idTag\" is required but was not present")
	}
	// check if meterStart (a required property) was received
	if !meterStartReceived {
		return errors.New("\"meterStart\" is required but was not present")
	}
	// check if timestamp (a required property) was received
	if !timestampReceived {
		return errors.New("\"timestamp\" is required but was not present")
	}
	return nil
}

func (m *StatusNotification) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "ConnectorId" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "connectorId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"connectorId\": ")
	if tmp, err := json.Marshal(m.ConnectorId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ErrorCode" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "errorCode" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"errorCode\": ")
	if tmp, err := json.Marshal(m.ErrorCode); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "info" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"info\": ")
	if tmp, err := json.Marshal(m.Info); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Status" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "status" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"status\": ")
	if tmp, err := json.Marshal(m.Status); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	// Marshal the "timestamp" field
	if m.Timestamp != nil {
		buf.WriteString(",")
		buf.WriteString("\"timestamp\": ")
		if tmp, err := json.Marshal(m.Timestamp.Format(time.RFC3339)); err != nil {
			return nil, err
		} else {
			buf.Write(tmp)
		}
	}
	comma = true
	// Marshal the "vendorErrorCode" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"vendorErrorCode\": ")
	if tmp, err := json.Marshal(m.VendorErrorCode); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "vendorId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"vendorId\": ")
	if tmp, err := json.Marshal(m.VendorId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *StatusNotification) UnmarshalJSON(b []byte) error {
	connectorIdReceived := false
	errorCodeReceived := false
	statusReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "connectorId":
			if err := json.Unmarshal([]byte(v), &m.ConnectorId); err != nil {
				return err
			}
			connectorIdReceived = true
		case "errorCode":
			if err := json.Unmarshal([]byte(v), &m.ErrorCode); err != nil {
				return err
			}
			errorCodeReceived = true
		case "info":
			if err := json.Unmarshal([]byte(v), &m.Info); err != nil {
				return err
			}
		case "status":
			if err := json.Unmarshal([]byte(v), &m.Status); err != nil {
				return err
			}
			statusReceived = true
		case "timestamp":
			if err := json.Unmarshal([]byte(v), &m.Timestamp); err != nil {
				return err
			}
		case "vendorErrorCode":
			if err := json.Unmarshal([]byte(v), &m.VendorErrorCode); err != nil {
				return err
			}
		case "vendorId":
			if err := json.Unmarshal([]byte(v), &m.VendorId); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if connectorId (a required property) was received
	if !connectorIdReceived {
		return errors.New("\"connectorId\" is required but was not present")
	}
	// check if errorCode (a required property) was received
	if !errorCodeReceived {
		return errors.New("\"errorCode\" is required but was not present")
	}
	// check if status (a required property) was received
	if !statusReceived {
		return errors.New("\"status\" is required but was not present")
	}
	return nil
}

func (m *StopTransaction) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "idTag" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"idTag\": ")
	if tmp, err := json.Marshal(m.IdTag); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "MeterStop" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "meterStop" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"meterStop\": ")
	if tmp, err := json.Marshal(m.MeterStop); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "reason" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"reason\": ")
	if tmp, err := json.Marshal(m.Reason); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Timestamp" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "timestamp" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"timestamp\": ")
	if tmp, err := json.Marshal(m.Timestamp.Format(time.RFC3339)); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "transactionData" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"transactionData\": ")
	if tmp, err := json.Marshal(m.TransactionData); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "TransactionId" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "transactionId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"transactionId\": ")
	if tmp, err := json.Marshal(m.TransactionId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *StopTransaction) UnmarshalJSON(b []byte) error {
	meterStopReceived := false
	timestampReceived := false
	transactionIdReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "idTag":
			if err := json.Unmarshal([]byte(v), &m.IdTag); err != nil {
				return err
			}
		case "meterStop":
			if err := json.Unmarshal([]byte(v), &m.MeterStop); err != nil {
				return err
			}
			meterStopReceived = true
		case "reason":
			if err := json.Unmarshal([]byte(v), &m.Reason); err != nil {
				return err
			}
		case "timestamp":
			if err := json.Unmarshal([]byte(v), &m.Timestamp); err != nil {
				return err
			}
			timestampReceived = true
		case "transactionData":
			if err := json.Unmarshal([]byte(v), &m.TransactionData); err != nil {
				return err
			}
		case "transactionId":
			if err := json.Unmarshal([]byte(v), &m.TransactionId); err != nil {
				return err
			}
			transactionIdReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if meterStop (a required property) was received
	if !meterStopReceived {
		return errors.New("\"meterStop\" is required but was not present")
	}
	// check if timestamp (a required property) was received
	if !timestampReceived {
		return errors.New("\"timestamp\" is required but was not present")
	}
	// check if transactionId (a required property) was received
	if !transactionIdReceived {
		return errors.New("\"transactionId\" is required but was not present")
	}
	return nil
}

func (m *TransactionDataItems) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "SampledValue" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "sampledValue" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"sampledValue\": ")
	if tmp, err := json.Marshal(m.SampledValues); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Timestamp" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "timestamp" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"timestamp\": ")
	if tmp, err := json.Marshal(m.Timestamp.Format(time.RFC3339)); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *TransactionDataItems) UnmarshalJSON(b []byte) error {
	sampledValueReceived := false
	timestampReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "sampledValue":
			if err := json.Unmarshal([]byte(v), &m.SampledValues); err != nil {
				return err
			}
			sampledValueReceived = true
		case "timestamp":
			if err := json.Unmarshal([]byte(v), &m.Timestamp); err != nil {
				return err
			}
			timestampReceived = true
		}
	}
	// check if sampledValue (a required property) was received
	if !sampledValueReceived {
		return errors.New("\"sampledValue\" is required but was not present")
	}
	// check if timestamp (a required property) was received
	if !timestampReceived {
		return errors.New("\"timestamp\" is required but was not present")
	}
	return nil
}
