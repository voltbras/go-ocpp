package csreq

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func (strct *CancelReservation) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "ReservationId" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "reservationId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"reservationId\": ")
	if tmp, err := json.Marshal(strct.ReservationId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *CancelReservation) UnmarshalJSON(b []byte) error {
	reservationIdReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "reservationId":
			if err := json.Unmarshal([]byte(v), &strct.ReservationId); err != nil {
				return err
			}
			reservationIdReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if reservationId (a required property) was received
	if !reservationIdReceived {
		return errors.New("\"reservationId\" is required but was not present")
	}
	return nil
}

func (strct *ChangeAvailability) MarshalJSON() ([]byte, error) {
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
	if tmp, err := json.Marshal(strct.ConnectorId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Type" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "type" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"type\": ")
	if tmp, err := json.Marshal(strct.Type); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ChangeAvailability) UnmarshalJSON(b []byte) error {
	connectorIdReceived := false
	typeReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "connectorId":
			if err := json.Unmarshal([]byte(v), &strct.ConnectorId); err != nil {
				return err
			}
			connectorIdReceived = true
		case "type":
			if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
				return err
			}
			typeReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if connectorId (a required property) was received
	if !connectorIdReceived {
		return errors.New("\"connectorId\" is required but was not present")
	}
	// check if type (a required property) was received
	if !typeReceived {
		return errors.New("\"type\" is required but was not present")
	}
	return nil
}

func (strct *ChangeConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Key" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "key" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"key\": ")
	if tmp, err := json.Marshal(strct.Key); err != nil {
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
	if tmp, err := json.Marshal(strct.Value); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ChangeConfiguration) UnmarshalJSON(b []byte) error {
	keyReceived := false
	valueReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "key":
			if err := json.Unmarshal([]byte(v), &strct.Key); err != nil {
				return err
			}
			keyReceived = true
		case "value":
			if err := json.Unmarshal([]byte(v), &strct.Value); err != nil {
				return err
			}
			valueReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if key (a required property) was received
	if !keyReceived {
		return errors.New("\"key\" is required but was not present")
	}
	// check if value (a required property) was received
	if !valueReceived {
		return errors.New("\"value\" is required but was not present")
	}
	return nil
}

func (strct *ChargingProfile) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "ChargingProfileId" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "chargingProfileId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingProfileId\": ")
	if tmp, err := json.Marshal(strct.ChargingProfileId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ChargingProfileKind" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "chargingProfileKind" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingProfileKind\": ")
	if tmp, err := json.Marshal(strct.ChargingProfileKind); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ChargingProfilePurpose" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "chargingProfilePurpose" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingProfilePurpose\": ")
	if tmp, err := json.Marshal(strct.ChargingProfilePurpose); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ChargingSchedule" field is required
	if strct.ChargingSchedule == nil {
		return nil, errors.New("chargingSchedule is a required field")
	}
	// Marshal the "chargingSchedule" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingSchedule\": ")
	if tmp, err := json.Marshal(strct.ChargingSchedule); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "recurrencyKind" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"recurrencyKind\": ")
	if tmp, err := json.Marshal(strct.RecurrencyKind); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "StackLevel" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "stackLevel" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"stackLevel\": ")
	if tmp, err := json.Marshal(strct.StackLevel); err != nil {
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
	if tmp, err := json.Marshal(strct.TransactionId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "validFrom" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"validFrom\": ")
	if tmp, err := json.Marshal(strct.ValidFrom.Format(time.RFC3339)); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "validTo" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"validTo\": ")
	if tmp, err := json.Marshal(strct.ValidTo.Format(time.RFC3339)); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ChargingProfile) UnmarshalJSON(b []byte) error {
	chargingProfileIdReceived := false
	chargingProfileKindReceived := false
	chargingProfilePurposeReceived := false
	chargingScheduleReceived := false
	stackLevelReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "chargingProfileId":
			if err := json.Unmarshal([]byte(v), &strct.ChargingProfileId); err != nil {
				return err
			}
			chargingProfileIdReceived = true
		case "chargingProfileKind":
			if err := json.Unmarshal([]byte(v), &strct.ChargingProfileKind); err != nil {
				return err
			}
			chargingProfileKindReceived = true
		case "chargingProfilePurpose":
			if err := json.Unmarshal([]byte(v), &strct.ChargingProfilePurpose); err != nil {
				return err
			}
			chargingProfilePurposeReceived = true
		case "chargingSchedule":
			if err := json.Unmarshal([]byte(v), &strct.ChargingSchedule); err != nil {
				return err
			}
			chargingScheduleReceived = true
		case "recurrencyKind":
			if err := json.Unmarshal([]byte(v), &strct.RecurrencyKind); err != nil {
				return err
			}
		case "stackLevel":
			if err := json.Unmarshal([]byte(v), &strct.StackLevel); err != nil {
				return err
			}
			stackLevelReceived = true
		case "transactionId":
			if err := json.Unmarshal([]byte(v), &strct.TransactionId); err != nil {
				return err
			}
		case "validFrom":
			if err := json.Unmarshal([]byte(v), &strct.ValidFrom); err != nil {
				return err
			}
		case "validTo":
			if err := json.Unmarshal([]byte(v), &strct.ValidTo); err != nil {
				return err
			}
		}
	}
	// check if chargingProfileId (a required property) was received
	if !chargingProfileIdReceived {
		return errors.New("\"chargingProfileId\" is required but was not present")
	}
	// check if chargingProfileKind (a required property) was received
	if !chargingProfileKindReceived {
		return errors.New("\"chargingProfileKind\" is required but was not present")
	}
	// check if chargingProfilePurpose (a required property) was received
	if !chargingProfilePurposeReceived {
		return errors.New("\"chargingProfilePurpose\" is required but was not present")
	}
	// check if chargingSchedule (a required property) was received
	if !chargingScheduleReceived {
		return errors.New("\"chargingSchedule\" is required but was not present")
	}
	// check if stackLevel (a required property) was received
	if !stackLevelReceived {
		return errors.New("\"stackLevel\" is required but was not present")
	}
	return nil
}

func (strct *ClearCache) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ClearCache) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k := range jsonMap {
		switch k {
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *ClearChargingProfile) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "chargingProfilePurpose" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingProfilePurpose\": ")
	if tmp, err := json.Marshal(strct.ChargingProfilePurpose); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "connectorId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"connectorId\": ")
	if tmp, err := json.Marshal(strct.ConnectorId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "id" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"id\": ")
	if tmp, err := json.Marshal(strct.Id); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "stackLevel" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"stackLevel\": ")
	if tmp, err := json.Marshal(strct.StackLevel); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ClearChargingProfile) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "chargingProfilePurpose":
			if err := json.Unmarshal([]byte(v), &strct.ChargingProfilePurpose); err != nil {
				return err
			}
		case "connectorId":
			if err := json.Unmarshal([]byte(v), &strct.ConnectorId); err != nil {
				return err
			}
		case "id":
			if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
				return err
			}
		case "stackLevel":
			if err := json.Unmarshal([]byte(v), &strct.StackLevel); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *DataTransfer) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "data" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"data\": ")
	if tmp, err := json.Marshal(strct.Data); err != nil {
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
	if tmp, err := json.Marshal(strct.MessageId); err != nil {
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
	if tmp, err := json.Marshal(strct.VendorId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *DataTransfer) UnmarshalJSON(b []byte) error {
	vendorIdReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "data":
			if err := json.Unmarshal([]byte(v), &strct.Data); err != nil {
				return err
			}
		case "messageId":
			if err := json.Unmarshal([]byte(v), &strct.MessageId); err != nil {
				return err
			}
		case "vendorId":
			if err := json.Unmarshal([]byte(v), &strct.VendorId); err != nil {
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

func (strct *GetCompositeSchedule) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "chargingRateUnit" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingRateUnit\": ")
	if tmp, err := json.Marshal(strct.ChargingRateUnit); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ConnectorId" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "connectorId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"connectorId\": ")
	if tmp, err := json.Marshal(strct.ConnectorId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Duration" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "duration" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"duration\": ")
	if tmp, err := json.Marshal(strct.Duration); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *GetCompositeSchedule) UnmarshalJSON(b []byte) error {
	connectorIdReceived := false
	durationReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "chargingRateUnit":
			if err := json.Unmarshal([]byte(v), &strct.ChargingRateUnit); err != nil {
				return err
			}
		case "connectorId":
			if err := json.Unmarshal([]byte(v), &strct.ConnectorId); err != nil {
				return err
			}
			connectorIdReceived = true
		case "duration":
			if err := json.Unmarshal([]byte(v), &strct.Duration); err != nil {
				return err
			}
			durationReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if connectorId (a required property) was received
	if !connectorIdReceived {
		return errors.New("\"connectorId\" is required but was not present")
	}
	// check if duration (a required property) was received
	if !durationReceived {
		return errors.New("\"duration\" is required but was not present")
	}
	return nil
}

func (strct *GetConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "key" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"key\": ")
	if tmp, err := json.Marshal(strct.Key); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *GetConfiguration) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "key":
			if err := json.Unmarshal([]byte(v), &strct.Key); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *GetDiagnostics) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Location" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "location" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"location\": ")
	if tmp, err := json.Marshal(strct.Location); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "retries" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"retries\": ")
	if tmp, err := json.Marshal(strct.Retries); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "retryInterval" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"retryInterval\": ")
	if tmp, err := json.Marshal(strct.RetryInterval); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "startTime" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"startTime\": ")
	if tmp, err := json.Marshal(strct.StartTime.Format(time.RFC3339)); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "stopTime" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"stopTime\": ")
	if tmp, err := json.Marshal(strct.StopTime.Format(time.RFC3339)); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *GetDiagnostics) UnmarshalJSON(b []byte) error {
	locationReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "location":
			if err := json.Unmarshal([]byte(v), &strct.Location); err != nil {
				return err
			}
			locationReceived = true
		case "retries":
			if err := json.Unmarshal([]byte(v), &strct.Retries); err != nil {
				return err
			}
		case "retryInterval":
			if err := json.Unmarshal([]byte(v), &strct.RetryInterval); err != nil {
				return err
			}
		case "startTime":
			if err := json.Unmarshal([]byte(v), &strct.StartTime); err != nil {
				return err
			}
		case "stopTime":
			if err := json.Unmarshal([]byte(v), &strct.StopTime); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if location (a required property) was received
	if !locationReceived {
		return errors.New("\"location\" is required but was not present")
	}
	return nil
}

func (strct *GetLocalListVersion) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *GetLocalListVersion) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k := range jsonMap {
		switch k {
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *IdTagInfo) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "expiryDate" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"expiryDate\": ")
	if tmp, err := json.Marshal(strct.ExpiryDate.Format(time.RFC3339)); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "parentIdTag" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"parentIdTag\": ")
	if tmp, err := json.Marshal(strct.ParentIdTag); err != nil {
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
	if tmp, err := json.Marshal(strct.Status); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *IdTagInfo) UnmarshalJSON(b []byte) error {
	statusReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "status":
			if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
				return err
			}
			statusReceived = true
		}
	}
	// check if status (a required property) was received
	if !statusReceived {
		return errors.New("\"status\" is required but was not present")
	}
	return nil
}

func (strct *LocalAuthorizationListItems) MarshalJSON() ([]byte, error) {
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
	if tmp, err := json.Marshal(strct.IdTag); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "idTagInfo" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"idTagInfo\": ")
	if tmp, err := json.Marshal(strct.IdTagInfo); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *LocalAuthorizationListItems) UnmarshalJSON(b []byte) error {
	idTagReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "idTag":
			if err := json.Unmarshal([]byte(v), &strct.IdTag); err != nil {
				return err
			}
			idTagReceived = true
		case "idTagInfo":
			if err := json.Unmarshal([]byte(v), &strct.IdTagInfo); err != nil {
				return err
			}
		}
	}
	// check if idTag (a required property) was received
	if !idTagReceived {
		return errors.New("\"idTag\" is required but was not present")
	}
	return nil
}

func (strct *RemoteStartTransaction) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "chargingProfile" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingProfile\": ")
	if tmp, err := json.Marshal(strct.ChargingProfile); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "connectorId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"connectorId\": ")
	if tmp, err := json.Marshal(strct.ConnectorId); err != nil {
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
	if tmp, err := json.Marshal(strct.IdTag); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *RemoteStartTransaction) UnmarshalJSON(b []byte) error {
	idTagReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "chargingProfile":
			if err := json.Unmarshal([]byte(v), &strct.ChargingProfile); err != nil {
				return err
			}
		case "connectorId":
			if err := json.Unmarshal([]byte(v), &strct.ConnectorId); err != nil {
				return err
			}
		case "idTag":
			if err := json.Unmarshal([]byte(v), &strct.IdTag); err != nil {
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

func (strct *RemoteStopTransaction) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "TransactionId" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "transactionId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"transactionId\": ")
	if tmp, err := json.Marshal(strct.TransactionId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *RemoteStopTransaction) UnmarshalJSON(b []byte) error {
	transactionIdReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "transactionId":
			if err := json.Unmarshal([]byte(v), &strct.TransactionId); err != nil {
				return err
			}
			transactionIdReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if transactionId (a required property) was received
	if !transactionIdReceived {
		return errors.New("\"transactionId\" is required but was not present")
	}
	return nil
}

func (strct *ReserveNow) MarshalJSON() ([]byte, error) {
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
	if tmp, err := json.Marshal(strct.ConnectorId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ExpiryDate" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "expiryDate" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"expiryDate\": ")
	if tmp, err := json.Marshal(strct.ExpiryDate.Format(time.RFC3339)); err != nil {
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
	if tmp, err := json.Marshal(strct.IdTag); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "parentIdTag" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"parentIdTag\": ")
	if tmp, err := json.Marshal(strct.ParentIdTag); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ReservationId" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "reservationId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"reservationId\": ")
	if tmp, err := json.Marshal(strct.ReservationId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ReserveNow) UnmarshalJSON(b []byte) error {
	connectorIdReceived := false
	expiryDateReceived := false
	idTagReceived := false
	reservationIdReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "connectorId":
			if err := json.Unmarshal([]byte(v), &strct.ConnectorId); err != nil {
				return err
			}
			connectorIdReceived = true
		case "expiryDate":
			if err := json.Unmarshal([]byte(v), &strct.ExpiryDate); err != nil {
				return err
			}
			expiryDateReceived = true
		case "idTag":
			if err := json.Unmarshal([]byte(v), &strct.IdTag); err != nil {
				return err
			}
			idTagReceived = true
		case "parentIdTag":
			if err := json.Unmarshal([]byte(v), &strct.ParentIdTag); err != nil {
				return err
			}
		case "reservationId":
			if err := json.Unmarshal([]byte(v), &strct.ReservationId); err != nil {
				return err
			}
			reservationIdReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if connectorId (a required property) was received
	if !connectorIdReceived {
		return errors.New("\"connectorId\" is required but was not present")
	}
	// check if expiryDate (a required property) was received
	if !expiryDateReceived {
		return errors.New("\"expiryDate\" is required but was not present")
	}
	// check if idTag (a required property) was received
	if !idTagReceived {
		return errors.New("\"idTag\" is required but was not present")
	}
	// check if reservationId (a required property) was received
	if !reservationIdReceived {
		return errors.New("\"reservationId\" is required but was not present")
	}
	return nil
}

func (strct *Reset) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Type" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "type" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"type\": ")
	if tmp, err := json.Marshal(strct.Type); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Reset) UnmarshalJSON(b []byte) error {
	typeReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "type":
			if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
				return err
			}
			typeReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if type (a required property) was received
	if !typeReceived {
		return errors.New("\"type\" is required but was not present")
	}
	return nil
}

func (strct *SendLocalList) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "ListVersion" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "listVersion" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"listVersion\": ")
	if tmp, err := json.Marshal(strct.ListVersion); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "localAuthorizationList" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"localAuthorizationList\": ")
	if tmp, err := json.Marshal(strct.LocalAuthorizationList); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "UpdateType" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "updateType" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"updateType\": ")
	if tmp, err := json.Marshal(strct.UpdateType); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *SendLocalList) UnmarshalJSON(b []byte) error {
	listVersionReceived := false
	updateTypeReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "listVersion":
			if err := json.Unmarshal([]byte(v), &strct.ListVersion); err != nil {
				return err
			}
			listVersionReceived = true
		case "localAuthorizationList":
			if err := json.Unmarshal([]byte(v), &strct.LocalAuthorizationList); err != nil {
				return err
			}
		case "updateType":
			if err := json.Unmarshal([]byte(v), &strct.UpdateType); err != nil {
				return err
			}
			updateTypeReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if listVersion (a required property) was received
	if !listVersionReceived {
		return errors.New("\"listVersion\" is required but was not present")
	}
	// check if updateType (a required property) was received
	if !updateTypeReceived {
		return errors.New("\"updateType\" is required but was not present")
	}
	return nil
}

func (strct *TriggerMessage) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "connectorId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"connectorId\": ")
	if tmp, err := json.Marshal(strct.ConnectorId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "RequestedMessage" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "requestedMessage" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"requestedMessage\": ")
	if tmp, err := json.Marshal(strct.RequestedMessage); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *TriggerMessage) UnmarshalJSON(b []byte) error {
	requestedMessageReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "connectorId":
			if err := json.Unmarshal([]byte(v), &strct.ConnectorId); err != nil {
				return err
			}
		case "requestedMessage":
			if err := json.Unmarshal([]byte(v), &strct.RequestedMessage); err != nil {
				return err
			}
			requestedMessageReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if requestedMessage (a required property) was received
	if !requestedMessageReceived {
		return errors.New("\"requestedMessage\" is required but was not present")
	}
	return nil
}

func (strct *ChargingSchedule) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "ChargingRateUnit" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "chargingRateUnit" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingRateUnit\": ")
	if tmp, err := json.Marshal(strct.ChargingRateUnit); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ChargingSchedulePeriod" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "chargingSchedulePeriod" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingSchedulePeriod\": ")
	if tmp, err := json.Marshal(strct.ChargingSchedulePeriod); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "duration" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"duration\": ")
	if tmp, err := json.Marshal(strct.Duration); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "minChargingRate" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"minChargingRate\": ")
	if tmp, err := json.Marshal(strct.MinChargingRate); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "startSchedule" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"startSchedule\": ")
	if tmp, err := json.Marshal(strct.StartSchedule.Format(time.RFC3339)); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ChargingSchedule) UnmarshalJSON(b []byte) error {
	chargingRateUnitReceived := false
	chargingSchedulePeriodReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "chargingRateUnit":
			if err := json.Unmarshal([]byte(v), &strct.ChargingRateUnit); err != nil {
				return err
			}
			chargingRateUnitReceived = true
		case "chargingSchedulePeriod":
			if err := json.Unmarshal([]byte(v), &strct.ChargingSchedulePeriod); err != nil {
				return err
			}
			chargingSchedulePeriodReceived = true
		case "duration":
			if err := json.Unmarshal([]byte(v), &strct.Duration); err != nil {
				return err
			}
		case "minChargingRate":
			if err := json.Unmarshal([]byte(v), &strct.MinChargingRate); err != nil {
				return err
			}
		case "startSchedule":
			if err := json.Unmarshal([]byte(v), &strct.StartSchedule); err != nil {
				return err
			}
		}
	}
	// check if chargingRateUnit (a required property) was received
	if !chargingRateUnitReceived {
		return errors.New("\"chargingRateUnit\" is required but was not present")
	}
	// check if chargingSchedulePeriod (a required property) was received
	if !chargingSchedulePeriodReceived {
		return errors.New("\"chargingSchedulePeriod\" is required but was not present")
	}
	return nil
}

func (strct *ChargingSchedulePeriodItems) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Limit" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "limit" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"limit\": ")
	if tmp, err := json.Marshal(strct.Limit); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "numberPhases" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"numberPhases\": ")
	if tmp, err := json.Marshal(strct.NumberPhases); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "StartPeriod" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "startPeriod" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"startPeriod\": ")
	if tmp, err := json.Marshal(strct.StartPeriod); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ChargingSchedulePeriodItems) UnmarshalJSON(b []byte) error {
	limitReceived := false
	startPeriodReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "limit":
			if err := json.Unmarshal([]byte(v), &strct.Limit); err != nil {
				return err
			}
			limitReceived = true
		case "numberPhases":
			if err := json.Unmarshal([]byte(v), &strct.NumberPhases); err != nil {
				return err
			}
		case "startPeriod":
			if err := json.Unmarshal([]byte(v), &strct.StartPeriod); err != nil {
				return err
			}
			startPeriodReceived = true
		}
	}
	// check if limit (a required property) was received
	if !limitReceived {
		return errors.New("\"limit\" is required but was not present")
	}
	// check if startPeriod (a required property) was received
	if !startPeriodReceived {
		return errors.New("\"startPeriod\" is required but was not present")
	}
	return nil
}

func (strct *CsChargingProfiles) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "ChargingProfileId" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "chargingProfileId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingProfileId\": ")
	if tmp, err := json.Marshal(strct.ChargingProfileId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ChargingProfileKind" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "chargingProfileKind" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingProfileKind\": ")
	if tmp, err := json.Marshal(strct.ChargingProfileKind); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ChargingProfilePurpose" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "chargingProfilePurpose" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingProfilePurpose\": ")
	if tmp, err := json.Marshal(strct.ChargingProfilePurpose); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ChargingSchedule" field is required
	if strct.ChargingSchedule == nil {
		return nil, errors.New("chargingSchedule is a required field")
	}
	// Marshal the "chargingSchedule" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingSchedule\": ")
	if tmp, err := json.Marshal(strct.ChargingSchedule); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "recurrencyKind" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"recurrencyKind\": ")
	if tmp, err := json.Marshal(strct.RecurrencyKind); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "StackLevel" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "stackLevel" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"stackLevel\": ")
	if tmp, err := json.Marshal(strct.StackLevel); err != nil {
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
	if tmp, err := json.Marshal(strct.TransactionId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "validFrom" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"validFrom\": ")
	if tmp, err := json.Marshal(strct.ValidFrom.Format(time.RFC3339)); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "validTo" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"validTo\": ")
	if tmp, err := json.Marshal(strct.ValidTo.Format(time.RFC3339)); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *CsChargingProfiles) UnmarshalJSON(b []byte) error {
	chargingProfileIdReceived := false
	chargingProfileKindReceived := false
	chargingProfilePurposeReceived := false
	chargingScheduleReceived := false
	stackLevelReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "chargingProfileId":
			if err := json.Unmarshal([]byte(v), &strct.ChargingProfileId); err != nil {
				return err
			}
			chargingProfileIdReceived = true
		case "chargingProfileKind":
			if err := json.Unmarshal([]byte(v), &strct.ChargingProfileKind); err != nil {
				return err
			}
			chargingProfileKindReceived = true
		case "chargingProfilePurpose":
			if err := json.Unmarshal([]byte(v), &strct.ChargingProfilePurpose); err != nil {
				return err
			}
			chargingProfilePurposeReceived = true
		case "chargingSchedule":
			if err := json.Unmarshal([]byte(v), &strct.ChargingSchedule); err != nil {
				return err
			}
			chargingScheduleReceived = true
		case "recurrencyKind":
			if err := json.Unmarshal([]byte(v), &strct.RecurrencyKind); err != nil {
				return err
			}
		case "stackLevel":
			if err := json.Unmarshal([]byte(v), &strct.StackLevel); err != nil {
				return err
			}
			stackLevelReceived = true
		case "transactionId":
			if err := json.Unmarshal([]byte(v), &strct.TransactionId); err != nil {
				return err
			}
		case "validFrom":
			if err := json.Unmarshal([]byte(v), &strct.ValidFrom); err != nil {
				return err
			}
		case "validTo":
			if err := json.Unmarshal([]byte(v), &strct.ValidTo); err != nil {
				return err
			}
		}
	}
	// check if chargingProfileId (a required property) was received
	if !chargingProfileIdReceived {
		return errors.New("\"chargingProfileId\" is required but was not present")
	}
	// check if chargingProfileKind (a required property) was received
	if !chargingProfileKindReceived {
		return errors.New("\"chargingProfileKind\" is required but was not present")
	}
	// check if chargingProfilePurpose (a required property) was received
	if !chargingProfilePurposeReceived {
		return errors.New("\"chargingProfilePurpose\" is required but was not present")
	}
	// check if chargingSchedule (a required property) was received
	if !chargingScheduleReceived {
		return errors.New("\"chargingSchedule\" is required but was not present")
	}
	// check if stackLevel (a required property) was received
	if !stackLevelReceived {
		return errors.New("\"stackLevel\" is required but was not present")
	}
	return nil
}

func (strct *SetChargingProfile) MarshalJSON() ([]byte, error) {
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
	if tmp, err := json.Marshal(strct.ConnectorId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "CsChargingProfiles" field is required
	if strct.CsChargingProfiles == nil {
		return nil, errors.New("csChargingProfiles is a required field")
	}
	// Marshal the "csChargingProfiles" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"csChargingProfiles\": ")
	if tmp, err := json.Marshal(strct.CsChargingProfiles); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *SetChargingProfile) UnmarshalJSON(b []byte) error {
	connectorIdReceived := false
	csChargingProfilesReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "connectorId":
			if err := json.Unmarshal([]byte(v), &strct.ConnectorId); err != nil {
				return err
			}
			connectorIdReceived = true
		case "csChargingProfiles":
			if err := json.Unmarshal([]byte(v), &strct.CsChargingProfiles); err != nil {
				return err
			}
			csChargingProfilesReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if connectorId (a required property) was received
	if !connectorIdReceived {
		return errors.New("\"connectorId\" is required but was not present")
	}
	// check if csChargingProfiles (a required property) was received
	if !csChargingProfilesReceived {
		return errors.New("\"csChargingProfiles\" is required but was not present")
	}
	return nil
}

func (strct *UnlockConnector) MarshalJSON() ([]byte, error) {
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
	if tmp, err := json.Marshal(strct.ConnectorId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *UnlockConnector) UnmarshalJSON(b []byte) error {
	connectorIdReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "connectorId":
			if err := json.Unmarshal([]byte(v), &strct.ConnectorId); err != nil {
				return err
			}
			connectorIdReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if connectorId (a required property) was received
	if !connectorIdReceived {
		return errors.New("\"connectorId\" is required but was not present")
	}
	return nil
}

func (strct *UpdateFirmware) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Location" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "location" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"location\": ")
	if tmp, err := json.Marshal(strct.Location); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "retries" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"retries\": ")
	if tmp, err := json.Marshal(strct.Retries); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "RetrieveDate" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "retrieveDate" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"retrieveDate\": ")
	if tmp, err := json.Marshal(strct.RetrieveDate.Format(time.RFC3339)); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "retryInterval" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"retryInterval\": ")
	if tmp, err := json.Marshal(strct.RetryInterval); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *UpdateFirmware) UnmarshalJSON(b []byte) error {
	locationReceived := false
	retrieveDateReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "location":
			if err := json.Unmarshal([]byte(v), &strct.Location); err != nil {
				return err
			}
			locationReceived = true
		case "retries":
			if err := json.Unmarshal([]byte(v), &strct.Retries); err != nil {
				return err
			}
		case "retrieveDate":
			if err := json.Unmarshal([]byte(v), &strct.RetrieveDate); err != nil {
				return err
			}
			retrieveDateReceived = true
		case "retryInterval":
			if err := json.Unmarshal([]byte(v), &strct.RetryInterval); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if location (a required property) was received
	if !locationReceived {
		return errors.New("\"location\" is required but was not present")
	}
	// check if retrieveDate (a required property) was received
	if !retrieveDateReceived {
		return errors.New("\"retrieveDate\" is required but was not present")
	}
	return nil
}
