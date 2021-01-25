package ws

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/eduhenke/go-ocpp/internal"
	"github.com/eduhenke/go-ocpp/internal/log"
	"github.com/eduhenke/go-ocpp/messages"
	"github.com/eduhenke/go-ocpp/messages/req"
	"github.com/google/uuid"

	"github.com/eduhenke/go-ocpp"
	"github.com/gorilla/websocket"
)

type CallResponse struct {
	response messages.Response
	err      error
}

type Conn struct {
	ctx       context.Context
	cancelCtx context.CancelFunc
	*websocket.Conn
	sendMux      sync.Mutex
	sentMessages map[MessageID]*CallMessage
	requests     chan struct {
		messages.Request
		MessageID
	}
	responsesOf map[MessageID]chan CallResponse
}

func newConn(socket *websocket.Conn) *Conn {
	ctx, cancel := context.WithCancel(context.Background())
	return &Conn{
		Conn:         socket,
		sentMessages: make(map[MessageID]*CallMessage, 0),
		requests: make(chan struct {
			messages.Request
			MessageID
		}, 0),
		ctx:         ctx,
		cancelCtx:   cancel,
		responsesOf: make(map[MessageID]chan CallResponse),
	}
}

func Dial(identity, csURL string, version ocpp.Version) (*Conn, error) {
	dialer := websocket.Dialer{
		Subprotocols: []string{ocppVersionToProtocol(version)},
	}
	socket, _, err := dialer.Dial(csURL+"/"+identity, http.Header{})
	if err != nil {
		return nil, err
	}
	return newConn(socket), err
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ocppVersionToProtocol(version ocpp.Version) string {
	switch version {
	case ocpp.V15:
		return "ocpp1.5"
	case ocpp.V16:
		return "ocpp1.6"
	}
	return ""
}

func Handshake(w http.ResponseWriter, r *http.Request, supportedVersions []ocpp.Version) (*Conn, error) {
	upgraderHeader := http.Header{}
	for _, v := range supportedVersions {
		upgraderHeader.Add("Sec-WebSocket-Protocol", ocppVersionToProtocol(v))
	}
	socket, err := upgrader.Upgrade(w, r, upgraderHeader)
	conn := newConn(socket)
	return conn, err
}

func (c *Conn) WriteJSON(data interface{}) error {
	c.sendMux.Lock()
	defer c.sendMux.Unlock()
	return c.Conn.WriteJSON(data)
}

func (c *Conn) Close() error {
	err := c.Conn.Close()
	c.cancelCtx()
	return err
}

func (c *Conn) WaitClose() <-chan struct{} {
	return c.ctx.Done()
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func UnmarshalMessage(msg []byte) (Message, error) {
	var frame []interface{}
	err := json.Unmarshal(msg, &frame)
	if err != nil {
		return nil, fmt.Errorf("on unmarshalling websocket message: %w", err)
	}
	msgType, ok := frame[0].(float64)
	if !ok {
		return nil, fmt.Errorf("first field is not a message type: %w", err)
	}
	idStr, ok := frame[1].(string)
	if !ok {
		return nil, fmt.Errorf("second field is not a message ID: %w", err)
	}
	id := MessageID(idStr)
	switch MessageType(msgType) {
	case Call:
		actionStr, ok := frame[2].(string)
		if !ok {
			return nil, fmt.Errorf("third field is not action: %w", err)
		}
		action := Action(actionStr)
		payload, ok := frame[3].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("fourth field is not payload: %w", err)
		}
		return &CallMessage{id, action, payload}, nil
	case CallResult:
		payload, ok := frame[2].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("third field is not payload: %w", err)
		}
		return &CallResultMessage{id, payload}, nil
	case CallError:
		codeStr, ok := frame[2].(string)
		if !ok {
			return nil, fmt.Errorf("third field is not error code: %w", err)
		}
		description, ok := frame[3].(string)
		if !ok {
			return nil, fmt.Errorf("fourth field is not error description: %w", err)
		}
		details, ok := frame[4].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("fifth field is not error details: %w", err)
		}
		return &CallErrorMessage{id, ErrorCode(codeStr), description, details}, nil
	}
	return nil, nil
}

func (c *Conn) ReadMessageAsync() <-chan error {
	readMessageResultChannel := make(chan error)
	go func() {
		readMessageResultChannel <- c.ReadMessage()
	}()
	return readMessageResultChannel
}

func (c *Conn) ReadMessage() error {
	_, messageBytes, err := c.Conn.ReadMessage()
	if err != nil {
		c.Close()
		return err
	}
	messageBytes = bytes.TrimSpace(bytes.Replace(messageBytes, newline, space, -1))
	log.Debug("Received a message, raw: %v", string(messageBytes))
	msg, err := UnmarshalMessage(messageBytes)
	if err != nil {
		return err
	}

	log.Debug("Received a message, parsed: %v", msg)

	if msg.Type() == CallResult || msg.Type() == CallError {
		_, ok := c.sentMessages[msg.ID()]
		if !ok {
			return errors.New("received call error/result without sending any call message")
		}
	}

	var wserr ErrorCode
	switch m := msg.(type) {
	case *CallMessage:
		var req messages.Request
		req, wserr = c.callToRequest(m)
		if wserr != Nil {
			msg := NewCallErrorMessage(msg.ID(), wserr, "on handling message")
			return c.sendMessage(msg)
		}
		c.requests <- struct {
			messages.Request
			MessageID
		}{req, msg.ID()}
	case *CallResultMessage:
		var resp messages.Response
		resp, wserr = c.callResultToResponse(m)
		c.responsesOf[m.ID()] <- CallResponse{
			response: resp,
			err:      wserr,
		}
	case *CallErrorMessage:
		c.responsesOf[m.ID()] <- CallResponse{
			response: nil,
			err:      m,
		}
	}
	return nil
}

func (c *Conn) callToRequest(call *CallMessage) (messages.Request, ErrorCode) {
	req := req.FromActionName(string(call.Action))
	if req == nil {
		return nil, NotSupported
	}
	originalPayload, err := json.Marshal(call.Payload)
	if err != nil {
		return nil, GenericError
	}
	err = json.Unmarshal(originalPayload, req)
	if err != nil {
		return nil, FormationViolation
	}
	return req, Nil
}

func (c *Conn) callResultToResponse(result *CallResultMessage) (messages.Response, ErrorCode) {
	id := result.ID()
	call, ok := c.sentMessages[id]
	if !ok {
		return nil, GenericError
	}
	if call == nil {
		return nil, NotSupported
	}
	resp := req.FromActionName(string(call.Action)).GetResponse()
	originalPayload, err := json.Marshal(result.Payload)
	if err != nil {
		return nil, GenericError
	}
	err = json.Unmarshal(originalPayload, resp)
	if err != nil {
		return nil, FormationViolation
	}
	return resp, Nil
}
func (c *Conn) SendResponse(id MessageID, response messages.Response, err error) error {
	return c.sendMessage(unmarshalResponse(id, response, err))
}
func (c *Conn) sendMessage(msg Message) error {
	c.sendMux.Lock()
	defer c.sendMux.Unlock()

	bts, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("on marshalling message: %w", err)
	}
	log.Debug("Sending message [raw]: %v", string(bts))
	err = c.Conn.WriteMessage(websocket.TextMessage, bts)
	if err != nil {
		return fmt.Errorf("on sending message: %w", err)
	}
	log.Debug("Sent message!")
	return nil
}

func (c *Conn) Requests() <-chan struct {
	messages.Request
	MessageID
} {
	return c.requests
}

func (c *Conn) SendRequest(request messages.Request) (messages.Response, error) {
	id := MessageID(uuid.New().String())
	msg, err := UnmarshalRequest(id, request)
	if err != nil {
		return nil, err
	}
	c.sentMessages[id] = msg
	c.responsesOf[id] = make(chan CallResponse)
	err = c.sendMessage(msg)
	if err != nil {
		return nil, err
	}
	select {
	case callResponse := <-c.responsesOf[id]:
		delete(c.sentMessages, id)
		delete(c.responsesOf, id)
		if callResponse.err != Nil {
			return nil, callResponse.err
		}
		return callResponse.response, nil
	case <-time.After(internal.DefaultRequestTimeout):
		return nil, fmt.Errorf("request timeout exceeded")
	}
}
