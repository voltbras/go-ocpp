package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/voltbras/go-ocpp"
	"github.com/voltbras/go-ocpp/cs"
	"github.com/voltbras/go-ocpp/messages/v1x/cpreq"
	"github.com/voltbras/go-ocpp/messages/v1x/cpresp"
)

func main() {
	ocpp.SetDebugLogger(log.New(os.Stdout, "DEBUG:", log.Ltime))
	ocpp.SetErrorLogger(log.New(os.Stderr, "ERROR:", log.Ltime))
	csys := cs.New()
	go csys.Run(":12811", func(req cpreq.ChargePointRequest, metadata cs.ChargePointRequestMetadata) (cpresp.ChargePointResponse, error) {
		fmt.Printf("EXAMPLE(MAIN): Request from %s\n", metadata.ChargePointID)
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
