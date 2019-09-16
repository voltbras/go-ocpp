package main

import (
	"errors"
	"fmt"
	"github.com/eduhenke/go-ocpp"
	"log"
	"os"
	"time"

	"github.com/eduhenke/go-ocpp/csystem"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpres"
)

func main() {
	ocpp.SetDebugLogger(log.New(os.Stdout, "DEBUG:", log.Ltime))
	ocpp.SetErrorLogger(log.New(os.Stderr, "ERROR:", log.Ltime))
	csys := csystem.New()
	go csys.Run(":12811", func(req cpreq.ChargePointRequest, cpID string) (cpres.ChargePointResponse, error) {
		fmt.Printf("EXAMPLE(MAIN): Request from %s\n", cpID)
		switch req := req.(type) {
		case *cpreq.BootNotification:
			return &cpres.BootNotification{
				Status:      "Accepted",
				CurrentTime: time.Now(),
				Interval:    60,
			}, nil

		case *cpreq.Heartbeat:
			fmt.Printf("EXAMPLE(MAIN): Heartbeat\n")
			return &cpres.Heartbeat{CurrentTime: time.Now()}, nil

		case *cpreq.StatusNotification:
			if req.Status != "Available" {
				// chargepoint is unavailable
			}
			return &cpres.StatusNotification{}, nil

		default:
			fmt.Printf("EXAMPLE(MAIN): action not supported: %s\n", req.Action())
			return nil, errors.New("Response not supported")
		}

	})
	select {}
}
