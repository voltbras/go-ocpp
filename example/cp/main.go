package main

import (
	"errors"
	"fmt"

	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/cp"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpresp"
	"github.com/eduhenke/go-ocpp/messages/v1x/csreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/csresp"
)

func main() {
	stationID := "5"
	centralSystemURL := "ws://localhost:12811"
	st, err := cp.New(stationID, centralSystemURL, ocpp.V16, ocpp.JSON, func(cprequest csreq.CentralSystemRequest) (csresp.CentralSystemResponse, error) {
		return nil, errors.New("not supported")
	}) // or ocpp.SOAP
	if err != nil {
		fmt.Println("could not create charge point:", err)
		return
	}
	rawResp, err := st.Send(&cpreq.Heartbeat{})
	if err != nil {
		fmt.Println("could't send heartbeat:", err)
		return
	}
	resp, ok := rawResp.(*cpresp.Heartbeat)
	if !ok {
		fmt.Println("response is not a heartbeat response")
		return
	}
	fmt.Println("got reply:", resp)
}
