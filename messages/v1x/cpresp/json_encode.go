package cpresp

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
	// "IdTagInfo" field is required
	if m.IdTagInfo == nil {
		return nil, errors.New("idTagInfo is a required field")
	}
	// Marshal the "idTagInfo" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"idTagInfo\": ")
	if tmp, err := json.Marshal(m.IdTagInfo); err != nil {
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
	idTagInfoReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "idTagInfo":
			if err := json.Unmarshal([]byte(v), &m.IdTagInfo); err != nil {
				return err
			}
			idTagInfoReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if idTagInfo (a required property) was received
	if !idTagInfoReceived {
		return errors.New("\"idTagInfo\" is required but was not present")
	}
	return nil
}

func (m *BootNotification) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "CurrentTime" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "currentTime" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"currentTime\": ")
	if tmp, err := json.Marshal(m.CurrentTime.Format(time.RFC3339)); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Interval" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "interval" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"interval\": ")
	if tmp, err := json.Marshal(m.Interval); err != nil {
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

func (m *BootNotification) UnmarshalJSON(b []byte) error {
	currentTimeReceived := false
	intervalReceived := false
	statusReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "currentTime":
			if err := json.Unmarshal([]byte(v), &m.CurrentTime); err != nil {
				return err
			}
			currentTimeReceived = true
		case "interval":
			if err := json.Unmarshal([]byte(v), &m.Interval); err != nil {
				return err
			}
			intervalReceived = true
		case "status":
			if err := json.Unmarshal([]byte(v), &m.Status); err != nil {
				return err
			}
			statusReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if currentTime (a required property) was received
	if !currentTimeReceived {
		return errors.New("\"currentTime\" is required but was not present")
	}
	// check if interval (a required property) was received
	if !intervalReceived {
		return errors.New("\"interval\" is required but was not present")
	}
	// check if status (a required property) was received
	if !statusReceived {
		return errors.New("\"status\" is required but was not present")
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

func (m *DiagnosticsStatusNotification) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *DiagnosticsStatusNotification) UnmarshalJSON(b []byte) error {
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

func (m *FirmwareStatusNotification) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *FirmwareStatusNotification) UnmarshalJSON(b []byte) error {
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

func (m *Heartbeat) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "CurrentTime" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "currentTime" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"currentTime\": ")
	if tmp, err := json.Marshal(m.CurrentTime.Format(time.RFC3339)); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *Heartbeat) UnmarshalJSON(b []byte) error {
	currentTimeReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "currentTime":
			if err := json.Unmarshal([]byte(v), &m.CurrentTime); err != nil {
				return err
			}
			currentTimeReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if currentTime (a required property) was received
	if !currentTimeReceived {
		return errors.New("\"currentTime\" is required but was not present")
	}
	return nil
}

func (m *IdTagInfo) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	if m.ExpiryDate != nil {
		buf.WriteString("\"expiryDate\": ")
		if tmp, err := json.Marshal(m.ExpiryDate.Format(time.RFC3339)); err != nil {
			return nil, err
		} else {
			buf.Write(tmp)
		}
		buf.WriteString(",")
	}
	buf.WriteString("\"parentIdTag\": ")
	if tmp, err := json.Marshal(m.ParentIdTag); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
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

func (m *IdTagInfo) UnmarshalJSON(b []byte) error {
	statusReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "expiryDate":
			if err := json.Unmarshal([]byte(v), &m.ExpiryDate); err != nil {
				return err
			}
		case "parentIdTag":
			if err := json.Unmarshal([]byte(v), &m.ParentIdTag); err != nil {
				return err
			}
		case "status":
			if err := json.Unmarshal([]byte(v), &m.Status); err != nil {
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

func (m *MeterValues) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *MeterValues) UnmarshalJSON(b []byte) error {
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

func (m *StartTransaction) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "IdTagInfo" field is required
	if m.IdTagInfo == nil {
		return nil, errors.New("idTagInfo is a required field")
	}
	// Marshal the "idTagInfo" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"idTagInfo\": ")
	if tmp, err := json.Marshal(m.IdTagInfo); err != nil {
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

func (m *StartTransaction) UnmarshalJSON(b []byte) error {
	idTagInfoReceived := false
	transactionIdReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "idTagInfo":
			if err := json.Unmarshal([]byte(v), &m.IdTagInfo); err != nil {
				return err
			}
			idTagInfoReceived = true
		case "transactionId":
			if err := json.Unmarshal([]byte(v), &m.TransactionId); err != nil {
				return err
			}
			transactionIdReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if idTagInfo (a required property) was received
	if !idTagInfoReceived {
		return errors.New("\"idTagInfo\" is required but was not present")
	}
	// check if transactionId (a required property) was received
	if !transactionIdReceived {
		return errors.New("\"transactionId\" is required but was not present")
	}
	return nil
}

func (m *StatusNotification) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (m *StatusNotification) UnmarshalJSON(b []byte) error {
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

func (m *StopTransaction) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "idTagInfo" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"idTagInfo\": ")
	if tmp, err := json.Marshal(m.IdTagInfo); err != nil {
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
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "idTagInfo":
			if err := json.Unmarshal([]byte(v), &m.IdTagInfo); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}
