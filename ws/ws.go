package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func IsWebSocketUpgrade(r *http.Request) bool {
	return websocket.IsWebSocketUpgrade(r)
}

func IsNormalCloseError(err error) bool {
	return websocket.IsCloseError(err, websocket.CloseNormalClosure)
}
