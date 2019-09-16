package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
)

func IsWebSocketUpgrade(r *http.Request) bool {
	return websocket.IsWebSocketUpgrade(r)
}