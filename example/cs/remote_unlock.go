package main

import (
	"errors"
	"time"

	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/cs"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpresp"
	"github.com/eduhenke/go-ocpp/messages/v1x/csreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/csresp"
)

func remote_unlock_main() {
	csys := cs.New()
	// this runs the central system on the given port
	// and handles each incoming ChargepointRequest
	go csys.Run(":12811", func(req cpreq.ChargePointRequest, cpID string) (cpresp.ChargePointResponse, error) {
		switch req.(type) {
		case *cpreq.BootNotification:
			return &cpresp.BootNotification{
				Status:      "Accepted",
				CurrentTime: time.Now(),
				Interval:    60,
			}, nil

		case *cpreq.Heartbeat:
			return &cpresp.Heartbeat{CurrentTime: time.Now()}, nil

		case *cpreq.StatusNotification:
			return &cpresp.StatusNotification{}, nil

		default:
			return nil, errors.New("Response not supported")
		}
	})

	cpID := "<CHARGEPOINT_ID>"
	// cpURL is only used if it's SOAP
	// if the communication is via WS
	// the cpURL will be ignored
	cpURL := "<CHARGEPOINT_URL>"
	cpVersion := ocpp.V16

	cpService, err := csys.GetServiceOf(cpID, cpVersion, cpURL)
	if err != nil {
		panic(err)
	}

	// this is the RFID tag to be
	// sent to the charger, if it's
	// only a purely remote transaction(i.e. no local action needed)
	// it can be anything(e.g. "VIRTUAL")
	tag := "VIRTUAL"
	cpResp, err := cpService.Send(&csreq.RemoteStartTransaction{
		IdTag:       tag,
		ConnectorId: 1,
	})
	if err != nil {
		// error on communicating
		panic(err)
	}
	resp := cpResp.(*csresp.RemoteStartTransaction)
	if resp.Status != "Accepted" {
		panic("not accepted")
	}

	select {}
}
