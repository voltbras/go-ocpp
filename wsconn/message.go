package wsconn

import (
	"encoding/json"
)

type Action string
type MessageID string
type MessageType float64

const (
	Call       MessageType = 2
	CallResult MessageType = 3
	CallError  MessageType = 4
)

type Message interface {
	MessageType() MessageType
	MessageID() MessageID
}

type CallMessage struct {
	id      MessageID
	Action
	Payload map[string]interface{}
}

func (msg *CallMessage) MessageType() MessageType {
	return Call
}

func (msg *CallMessage) MessageID() MessageID {
	return msg.id
}

func (msg *CallMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{msg.MessageType(), msg.id, msg.Action, msg.Payload})
}


type CallResultMessage struct {
	id      MessageID
	payload interface{} //map[string]interface{}
}

func NewCallResult(id MessageID, payload interface{}) *CallResultMessage {
	return &CallResultMessage{
		id:      id,
		payload: payload,
	}
}

func (msg *CallResultMessage) MessageType() MessageType {
	return CallResult
}

func (msg *CallResultMessage) MessageID() MessageID {
	return msg.id
}
func (msg *CallResultMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{msg.MessageType(), msg.id, msg.payload})
}


type ErrorCode string

const (
	// NotSupported when Requested Action is recognized but not supported by the receiver
	NotSupported ErrorCode = "NotSupported"
	// InternalError when An internal error occurred and the receiver was not able to process the requested Action successfully
	InternalError ErrorCode = "InternalError"
	// ProtocolError when Payload for Action is incomplete
	ProtocolError ErrorCode = "ProtocolError"
	// SecurityError when During the processing of Action a security issue occurred preventing receiver from completing the Action successfully
	SecurityError ErrorCode = "SecurityError"
	// FormationViolation when Payload for Action is syntactically incorrect or not conform the PDU structure for Action
	FormationViolation ErrorCode = "FormationViolation"
	// PropertyConstraintViolation when Payload is syntactically correct but at least one field contains an invalid value
	PropertyConstraintViolation ErrorCode = "PropertyConstraintViolation"
	// OccurenceConstraintViolation when Payload for Action is syntactically correct but at least one of the fields violates occurence constraints
	OccurenceConstraintViolation ErrorCode = "OccurenceConstraintViolation"
	// TypeConstraintViolation when Payload for Action is syntactically correct but at least one of the fields violates data type constraints (e.g. “somestring”: 12)
	TypeConstraintViolation ErrorCode = "TypeConstraintViolation"
	// GenericError when Any other error not covered by the previous ones
	GenericError ErrorCode = "GenericError"
	// Nil no error
	Nil ErrorCode = ""
)

type CallErrorMessage struct {
	id               MessageID
	errorCode        ErrorCode
	errorDescription string
	errorDetails     map[string]interface{}
}

func NewCallErrorMessage(id MessageID, errorCode ErrorCode, errorDescription string) *CallErrorMessage {
	return &CallErrorMessage{
		id:               id,
		errorCode:        errorCode,
		errorDescription: errorDescription,
		errorDetails:     nil,
	}
}

func (msg *CallErrorMessage) MessageType() MessageType {
	return CallError
}

func (msg *CallErrorMessage) MessageID() MessageID {
	return msg.id
}
func (msg *CallErrorMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{msg.MessageType(), msg.id, msg.errorCode, msg.errorDescription, msg.errorDetails})
}
