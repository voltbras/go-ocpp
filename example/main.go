package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/cstationsim"
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

func main() {
	csys := Handler{}
	go csys.Run(":12811")

	sim := cstationsim.NewStation(":12812", "http://localhost:12811")
	reply, err := sim.CsService.BootNotification(messages.BootNotificationRequest{
		ChargePointVendor:       "Vendor",
		ChargePointModel:        "Simulator-01",
		ChargePointSerialNumber: "1337",
		ChargeBoxSerialNumber:   "1337",
		FirmwareVersion:         "1.7.0",
	})
	if err != nil {
		fmt.Println("could't send boot notification:", err)
	}
	fmt.Println("got reply:", reply)
}
