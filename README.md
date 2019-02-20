# go-ocpp
v1.5 OCPP implementation in Golang

# Usage
## Central System
Your `Handler` struct must implement the `HandleEnvelope` method:

```go
package handler

import (
	"errors"
	"time"

	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/csystem"
	"github.com/eduhenke/go-ocpp/messages"
)

type Handler struct{}

func (h *Handler) Run(port string) {
	csystem.ListenAndServe(port, h)
}

func (h *Handler) HandleEnvelope(env ocpp.Envelope) (interface{}, error) {
	// Return an error to the Station communicating to the Central System
	//
	// station, isAuthorized := getStation(env.Header.ChargeBoxIdentity)
	// if err != nil {
	// 	return nil, errors.New("charger not authorized to join network")
	// }

	switch req := env.Body.Content.(type) {
	case *messages.BootNotificationRequest:
		return messages.BootNotificationResponse{
            // accept chargepoint in the network
			Status:            messages.RegistrationStatusAccepted,
			CurrentTime:       time.Now(),
			HeartbeatInterval: 60,
		}, nil

	case *messages.HeartbeatRequest:
		return messages.HeartbeatResponse{CurrentTime: time.Now()}, nil

	case *messages.StatusNotificationRequest:
		if req.Status != messages.ChargePointStatusAvailable {
			// chargepoint is unavailable
		}
		return messages.StatusNotificationResponse{}, nil

	default:
		return nil, errors.New("Response not supported")
	}
}

```