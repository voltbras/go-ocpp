package csres

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

func (strct *CancelReservation) UnmarshalJSON(b []byte) error {
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

func (strct *ChangeAvailability) MarshalJSON() ([]byte, error) {
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

func (strct *ChangeAvailability) UnmarshalJSON(b []byte) error {
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

func (strct *ChangeConfiguration) MarshalJSON() ([]byte, error) {
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

func (strct *ChangeConfiguration) UnmarshalJSON(b []byte) error {
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

func (strct *ClearCache) MarshalJSON() ([]byte, error) {
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

func (strct *ClearCache) UnmarshalJSON(b []byte) error {
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

func (strct *ClearChargingProfile) MarshalJSON() ([]byte, error) {
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

func (strct *ClearChargingProfile) UnmarshalJSON(b []byte) error {
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

func (strct *ConfigurationKeyItems) MarshalJSON() ([]byte, error) {
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
	// "Readonly" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "readonly" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"readonly\": ")
	if tmp, err := json.Marshal(strct.Readonly); err != nil {
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

func (strct *ConfigurationKeyItems) UnmarshalJSON(b []byte) error {
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
			if err := json.Unmarshal([]byte(v), &strct.Key); err != nil {
				return err
			}
			keyReceived = true
		case "readonly":
			if err := json.Unmarshal([]byte(v), &strct.Readonly); err != nil {
				return err
			}
			readonlyReceived = true
		case "value":
			if err := json.Unmarshal([]byte(v), &strct.Value); err != nil {
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

func (strct *DataTransfer) UnmarshalJSON(b []byte) error {
	statusReceived := false
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
		case "status":
			if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
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

func (strct *GetCompositeSchedule) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
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
	// Marshal the "scheduleStart" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"scheduleStart\": ")
	if tmp, err := json.Marshal(strct.ScheduleStart.Format(time.RFC3339)); err != nil {
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

func (strct *GetCompositeSchedule) UnmarshalJSON(b []byte) error {
	statusReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "chargingSchedule":
			if err := json.Unmarshal([]byte(v), &strct.ChargingSchedule); err != nil {
				return err
			}
		case "connectorId":
			if err := json.Unmarshal([]byte(v), &strct.ConnectorId); err != nil {
				return err
			}
		case "scheduleStart":
			if err := json.Unmarshal([]byte(v), &strct.ScheduleStart); err != nil {
				return err
			}
		case "status":
			if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
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

func (strct *GetConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "configurationKey" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"configurationKey\": ")
	if tmp, err := json.Marshal(strct.ConfigurationKey); err != nil {
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
	if tmp, err := json.Marshal(strct.UnknownKey); err != nil {
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
		case "configurationKey":
			if err := json.Unmarshal([]byte(v), &strct.ConfigurationKey); err != nil {
				return err
			}
		case "unknownKey":
			if err := json.Unmarshal([]byte(v), &strct.UnknownKey); err != nil {
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
	// Marshal the "fileName" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"fileName\": ")
	if tmp, err := json.Marshal(strct.FileName); err != nil {
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
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "fileName":
			if err := json.Unmarshal([]byte(v), &strct.FileName); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *GetLocalListVersion) MarshalJSON() ([]byte, error) {
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *GetLocalListVersion) UnmarshalJSON(b []byte) error {
	listVersionReceived := false
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

func (strct *RemoteStartTransaction) MarshalJSON() ([]byte, error) {
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

func (strct *RemoteStartTransaction) UnmarshalJSON(b []byte) error {
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

func (strct *RemoteStopTransaction) MarshalJSON() ([]byte, error) {
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

func (strct *RemoteStopTransaction) UnmarshalJSON(b []byte) error {
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

func (strct *ReserveNow) MarshalJSON() ([]byte, error) {
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

func (strct *ReserveNow) UnmarshalJSON(b []byte) error {
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

func (strct *Reset) MarshalJSON() ([]byte, error) {
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

func (strct *Reset) UnmarshalJSON(b []byte) error {
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

func (strct *SendLocalList) MarshalJSON() ([]byte, error) {
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

func (strct *SendLocalList) UnmarshalJSON(b []byte) error {
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

func (strct *SetChargingProfile) MarshalJSON() ([]byte, error) {
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

func (strct *SetChargingProfile) UnmarshalJSON(b []byte) error {
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

func (strct *TriggerMessage) MarshalJSON() ([]byte, error) {
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

func (strct *TriggerMessage) UnmarshalJSON(b []byte) error {
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

func (strct *UnlockConnector) MarshalJSON() ([]byte, error) {
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

func (strct *UnlockConnector) UnmarshalJSON(b []byte) error {
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

func (strct *UpdateFirmware) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *UpdateFirmware) UnmarshalJSON(b []byte) error {
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
