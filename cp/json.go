package cp

import (
	"time"

	"github.com/eduhenke/go-ocpp/internal/log"
	"github.com/eduhenke/go-ocpp/internal/service"
	"github.com/eduhenke/go-ocpp/messages/v1x/csreq"
	"github.com/eduhenke/go-ocpp/ws"
)

const (
	websocketConnectionRetryInterval = 5 * time.Second
)

func (cp *chargePoint) Connection() *ws.Conn {
	return cp.conn
}

// tries to reach CS, if succeeded handle
func (cp *chargePoint) getNewWebsocketConnection() error {
	conn, err := ws.Dial(cp.identity, cp.centralSystemURL, cp.version)
	if err != nil {
		return err
	}
	cp.conn = conn
	cp.CentralSystem = service.NewCentralSystemJSON(cp.conn)
	// closing the channel will make the reads non blocking
	close(cp.connectedChan)
	return nil
}

func (cp *chargePoint) handleWebsocketConnection(cshandler CentralSystemMessageHandler) {
	for {
		select {
		case <-cp.ctx.Done():
			cp.conn.Close()
			return
		case <-cp.conn.WaitClose():
			log.Debug("Closed connection of Central System")
			cp.connectedChan = make(chan struct{})
			// try to connect until it is established
			for {
				err := cp.getNewWebsocketConnection()
				if err != nil {
					log.Error("On restarting connection with Central System: %w", err)
					<-time.After(websocketConnectionRetryInterval)
				} else {
					break
				}
			}
		case <-cp.conn.ReadMessageAsync():
			continue
		case req := <-cp.conn.Requests():
			cprequest, ok := req.Request.(csreq.CentralSystemRequest)
			if !ok {
				log.Error(csreq.ErrorNotCentralSystemRequest.Error())
			}
			cpresponse, err := cshandler(cprequest)
			err = cp.conn.SendResponse(req.MessageID, cpresponse, err)
			if err != nil {
				log.Error(err.Error())
			}
		}
	}
}

func (cp *chargePoint) WaitConnect() <-chan struct{} {
	return cp.connectedChan
}

func (cp *chargePoint) WaitDisconnect() <-chan struct{} {
	return cp.conn.WaitClose()
}
