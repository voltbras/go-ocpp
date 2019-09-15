package main

import (
	"errors"
	"time"

	"github.com/eduhenke/go-ocpp/csystem"
	"github.com/eduhenke/go-ocpp/messages/v16/cpreq"
	"github.com/eduhenke/go-ocpp/messages/v16/cpres"
)

func main() {

	csystem.ListenAndServe(":12811", func(req cpreq.ChargePointRequest) (cpres.ChargePointResponse, error) {
		switch req := req.(type) {
		case *cpreq.BootNotification:
			return &cpres.BootNotification{
				Status:      "Accepted",
				CurrentTime: time.Now(),
				Interval:    60,
			}, nil

		case *cpreq.Heartbeat:
			return &cpres.Heartbeat{CurrentTime: time.Now()}, nil

		case *cpreq.StatusNotification:
			if req.Status != "Available" {
				// chargepoint is unavailable
			}
			return &cpres.StatusNotification{}, nil

		default:
			return nil, errors.New("Response not supported")
		}

	})
	// sim := cstationsim.NewStation(":12812", "http://localhost:12811")
	// reply, err := sim.CsService.BootNotification(messages.BootNotificationRequest{
	// 	ChargePointVendor:       "Vendor",
	// 	ChargePointModel:        "Simulator-01",
	// 	ChargePointSerialNumber: "1337",
	// 	ChargeBoxSerialNumber:   "1337",
	// 	FirmwareVersion:         "1.7.0",
	// })
	// if err != nil {
	// 	fmt.Println("could't send boot notification:", err)
	// }
	// fmt.Println("got reply:", reply)
}
