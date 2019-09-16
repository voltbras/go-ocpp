package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/csystem"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpresp"
)

func main() {
	ocpp.SetDebugLogger(log.New(os.Stdout, "DEBUG:", log.Ltime))
	ocpp.SetErrorLogger(log.New(os.Stderr, "ERROR:", log.Ltime))
	csys := csystem.New()
	go csys.Run(":12811", func(req cpreq.ChargePointRequest, cpID string) (cpresp.ChargePointResponse, error) {
		fmt.Printf("EXAMPLE(MAIN): Request from %s\n", cpID)
		switch req := req.(type) {
		case *cpreq.BootNotification:
			return &cpresp.BootNotification{
				Status:      "Accepted",
				CurrentTime: time.Now(),
				Interval:    60,
			}, nil

		case *cpreq.Heartbeat:
			fmt.Printf("EXAMPLE(MAIN): Heartbeat\n")
			return &cpresp.Heartbeat{CurrentTime: time.Now()}, nil

		case *cpreq.StatusNotification:
			if req.Status != "Available" {
				// chargepoint is unavailable
			}
			return &cpresp.StatusNotification{}, nil

		default:
			fmt.Printf("EXAMPLE(MAIN): action not supported: %s\n", req.Action())
			return nil, errors.New("Response not supported")
		}

	})
	select {}
}
