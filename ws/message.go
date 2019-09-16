package ws

import (
	"encoding/json"
	"github.com/eduhenke/go-ocpp/messages"
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
	Type() MessageType
	ID() MessageID
}

type CallMessage struct {
	id MessageID
	Action
	Payload map[string]interface{}
}

func NewCallMessage(id MessageID, action Action, payload map[string]interface{}) *CallMessage {
	return &CallMessage{
		id:               id,
		Action: action,
		Payload: payload,
	}
}

func (call *CallMessage) Type() MessageType {
	return Call
}

func (call *CallMessage) ID() MessageID {
	return call.id
}

func (call *CallMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{call.Type(), call.id, call.Action, call.Payload})
}

type CallResultMessage struct {
	id      MessageID
	Payload interface{} //map[string]interface{}
}

func NewCallResult(id MessageID, payload interface{}) *CallResultMessage {
	return &CallResultMessage{
		id:      id,
		Payload: payload,
	}
}

func (result *CallResultMessage) Type() MessageType {
	return CallResult
}

func (result *CallResultMessage) ID() MessageID {
	return result.id
}
func (result *CallResultMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{result.Type(), result.id, result.Payload})
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

func (err *CallErrorMessage) Type() MessageType {
	return CallError
}

func (err *CallErrorMessage) ID() MessageID {
	return err.id
}
func (err *CallErrorMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{err.Type(), err.id, err.errorCode, err.errorDescription, err.errorDetails})
}

func unmarshalResponse(id MessageID, resp messages.Response, err error) Message {
	if err != nil {
		return NewCallErrorMessage(id, InternalError, err.Error())
	}
	return NewCallResult(id, resp)
}

func UnmarshalRequest(id MessageID, req messages.Request) (*CallMessage, error) {
	var inInterface map[string]interface{}
	inrec, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &inInterface)
	if err != nil {
		return nil, err
	}
	return NewCallMessage(id, Action(req.Action()), inInterface), nil
}