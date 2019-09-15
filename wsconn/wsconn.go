package wsconn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/eduhenke/go-ocpp"
	ws "github.com/gorilla/websocket"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr, "xxx: ", log.Ldate|log.Ltime|log.Lshortfile)
}

type Conn struct {
	*ws.Conn
	sendMux sync.Mutex
}

var upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func New(w http.ResponseWriter, r *http.Request, supportedVersions []ocpp.Version) (*Conn, error) {
	upgraderHeader := http.Header{}
	for _, v := range supportedVersions {
		upgraderHeader.Add("Sec-WebSocket-Protocol", "ocpp"+string(v))
	}
	socket, err := upgrader.Upgrade(w, r, upgraderHeader)
	conn := &Conn{Conn: socket}
	return conn, err
}

func (c *Conn) WriteJSON(data interface{}) error {
	fmt.Println("LOCK...")
	c.sendMux.Lock()
	defer func() {
		c.sendMux.Unlock()
		fmt.Println("UNLOCK!")
	}()
	return c.Conn.WriteJSON(data)
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
	fmt.Println("FRAME:", frame, frame[0])
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
	}
	return nil, nil
}

func (c *Conn) ReadMessage() (Message, error) {
	_, messageBytes, err := c.Conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	messageBytes = bytes.TrimSpace(bytes.Replace(messageBytes, newline, space, -1))
	msg, err := UnmarshalMessage(messageBytes)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
