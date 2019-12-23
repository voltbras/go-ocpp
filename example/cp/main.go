package main

import (
	"fmt"

	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/cp"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpresp"
)

func main() {
	stationID := "19400010"
	centralSystemURL := "ws://ocpp.voltbras.com.br:80"
	st, err := cp.NewChargePoint(stationID, centralSystemURL, ocpp.V16, ocpp.JSON) // or ocpp.SOAP
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
