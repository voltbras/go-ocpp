package csres

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func (m *CancelReservation) MarshalJSON() ([]byte, error) {
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

func (m *CancelReservation) UnmarshalJSON(b []byte) error {
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

func (m *ChangeAvailability) MarshalJSON() ([]byte, error) {
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

func (m *ChangeAvailability) UnmarshalJSON(b []byte) error {
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

func (m *ChangeConfiguration) MarshalJSON() ([]byte, error) {
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

func (m *ChangeConfiguration) UnmarshalJSON(b []byte) error {
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

func (m *ChargingSchedule) MarshalJSON() ([]byte, error) {
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
	if tmp, err := json.Marshal(m.ChargingRateUnit); err != nil {
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
	if tmp, err := json.Marshal(m.ChargingSchedulePeriod); err != nil {
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
	if tmp, err := json.Marshal(m.Duration); err != nil {
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
	if tmp, err := json.Marshal(m.MinChargingRate); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}

	if m.StartSchedule != nil {
		buf.WriteString(",")
		buf.WriteString("\"startSchedule\": ")
		if tmp, err := json.Marshal(m.StartSchedule.Format(time.RFC3339)); err != nil {
			return nil, err
		} else {
			buf.Write(tmp)
		}
	}

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *ChargingSchedule) UnmarshalJSON(b []byte) error {
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
			if err := json.Unmarshal([]byte(v), &m.ChargingRateUnit); err != nil {
				return err
			}
			chargingRateUnitReceived = true
		case "chargingSchedulePeriod":
			if err := json.Unmarshal([]byte(v), &m.ChargingSchedulePeriod); err != nil {
				return err
			}
			chargingSchedulePeriodReceived = true
		case "duration":
			if err := json.Unmarshal([]byte(v), &m.Duration); err != nil {
				return err
			}
		case "minChargingRate":
			if err := json.Unmarshal([]byte(v), &m.MinChargingRate); err != nil {
				return err
			}
		case "startSchedule":
			if err := json.Unmarshal([]byte(v), &m.StartSchedule); err != nil {
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

func (m *ChargingSchedulePeriodItems) MarshalJSON() ([]byte, error) {
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
	if tmp, err := json.Marshal(m.Limit); err != nil {
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
	if tmp, err := json.Marshal(m.NumberPhases); err != nil {
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
	if tmp, err := json.Marshal(m.StartPeriod); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *ChargingSchedulePeriodItems) UnmarshalJSON(b []byte) error {
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
			if err := json.Unmarshal([]byte(v), &m.Limit); err != nil {
				return err
			}
			limitReceived = true
		case "numberPhases":
			if err := json.Unmarshal([]byte(v), &m.NumberPhases); err != nil {
				return err
			}
		case "startPeriod":
			if err := json.Unmarshal([]byte(v), &m.StartPeriod); err != nil {
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

func (m *ClearCache) MarshalJSON() ([]byte, error) {
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

func (m *ClearCache) UnmarshalJSON(b []byte) error {
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

func (m *ClearChargingProfile) MarshalJSON() ([]byte, error) {
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

func (m *ClearChargingProfile) UnmarshalJSON(b []byte) error {
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

func (m *ConfigurationKeyItems) MarshalJSON() ([]byte, error) {
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
	if tmp, err := json.Marshal(m.Key); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Readonly" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "readonly" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"readonly\": ")
	if tmp, err := json.Marshal(m.Readonly); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
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

func (m *ConfigurationKeyItems) UnmarshalJSON(b []byte) error {
	keyReceived := false
	readonlyReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "key":
			if err := json.Unmarshal([]byte(v), &m.Key); err != nil {
				return err
			}
			keyReceived = true
		case "readonly":
			if err := json.Unmarshal([]byte(v), &m.Readonly); err != nil {
				return err
			}
			readonlyReceived = true
		case "value":
			if err := json.Unmarshal([]byte(v), &m.Value); err != nil {
				return err
			}
		}
	}
	// check if key (a required property) was received
	if !keyReceived {
		return errors.New("\"key\" is required but was not present")
	}
	// check if readonly (a required property) was received
	if !readonlyReceived {
		return errors.New("\"readonly\" is required but was not present")
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

func (m *DataTransfer) UnmarshalJSON(b []byte) error {
	statusReceived := false
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

func (m *GetCompositeSchedule) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "chargingSchedule" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingSchedule\": ")
	if tmp, err := json.Marshal(m.ChargingSchedule); err != nil {
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
	if tmp, err := json.Marshal(m.ConnectorId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	// Marshal the "scheduleStart" field
	if m.ScheduleStart != nil {
		buf.WriteString(",")
		buf.WriteString("\"scheduleStart\": ")
		if tmp, err := json.Marshal(m.ScheduleStart.Format(time.RFC3339)); err != nil {
			return nil, err
		} else {
			buf.Write(tmp)
		}
	}
	// "Status" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "status" field
	buf.WriteString(",")
	buf.WriteString("\"status\": ")
	if tmp, err := json.Marshal(m.Status); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *GetCompositeSchedule) UnmarshalJSON(b []byte) error {
	statusReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "chargingSchedule":
			if err := json.Unmarshal([]byte(v), &m.ChargingSchedule); err != nil {
				return err
			}
		case "connectorId":
			if err := json.Unmarshal([]byte(v), &m.ConnectorId); err != nil {
				return err
			}
		case "scheduleStart":
			if err := json.Unmarshal([]byte(v), &m.ScheduleStart); err != nil {
				return err
			}
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

func (m *GetConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "configurationKey" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"configurationKey\": ")
	if tmp, err := json.Marshal(m.ConfigurationKey); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "unknownKey" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"unknownKey\": ")
	if tmp, err := json.Marshal(m.UnknownKey); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *GetConfiguration) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "configurationKey":
			if err := json.Unmarshal([]byte(v), &m.ConfigurationKey); err != nil {
				return err
			}
		case "unknownKey":
			if err := json.Unmarshal([]byte(v), &m.UnknownKey); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (m *GetDiagnostics) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "fileName" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"fileName\": ")
	if tmp, err := json.Marshal(m.FileName); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *GetDiagnostics) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "fileName":
			if err := json.Unmarshal([]byte(v), &m.FileName); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (m *GetLocalListVersion) MarshalJSON() ([]byte, error) {
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
	if tmp, err := json.Marshal(m.ListVersion); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *GetLocalListVersion) UnmarshalJSON(b []byte) error {
	listVersionReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "listVersion":
			if err := json.Unmarshal([]byte(v), &m.ListVersion); err != nil {
				return err
			}
			listVersionReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if listVersion (a required property) was received
	if !listVersionReceived {
		return errors.New("\"listVersion\" is required but was not present")
	}
	return nil
}

func (m *RemoteStartTransaction) MarshalJSON() ([]byte, error) {
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

func (m *RemoteStartTransaction) UnmarshalJSON(b []byte) error {
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

func (m *RemoteStopTransaction) MarshalJSON() ([]byte, error) {
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

func (m *RemoteStopTransaction) UnmarshalJSON(b []byte) error {
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

func (m *ReserveNow) MarshalJSON() ([]byte, error) {
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

func (m *ReserveNow) UnmarshalJSON(b []byte) error {
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

func (m *Reset) MarshalJSON() ([]byte, error) {
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

func (m *Reset) UnmarshalJSON(b []byte) error {
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

func (m *SendLocalList) MarshalJSON() ([]byte, error) {
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

func (m *SendLocalList) UnmarshalJSON(b []byte) error {
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

func (m *SetChargingProfile) MarshalJSON() ([]byte, error) {
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

func (m *SetChargingProfile) UnmarshalJSON(b []byte) error {
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

func (m *TriggerMessage) MarshalJSON() ([]byte, error) {
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

func (m *TriggerMessage) UnmarshalJSON(b []byte) error {
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

func (m *UnlockConnector) MarshalJSON() ([]byte, error) {
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

func (m *UnlockConnector) UnmarshalJSON(b []byte) error {
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

func (m *UpdateFirmware) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *UpdateFirmware) UnmarshalJSON(b []byte) error {
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
