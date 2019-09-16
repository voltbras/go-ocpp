# go-ocpp

OCPP(1.5/1.6) implementation in Golang.

- v1.5, it's assumed it is SOAP
- v1.6, it's assumed it is JSON

## Usage

### Central System

Your `Handler` struct must implement the `HandleEnvelope` method:

```go
package main

import (
    "errors"
    "time"

    "github.com/eduhenke/go-ocpp"
    "github.com/eduhenke/go-ocpp/csystem"
    "github.com/eduhenke/go-ocpp/messages/v1x/cpreq"
    "github.com/eduhenke/go-ocpp/messages/v1x/cpres"
)

func main() {
    csys := csystem.New()
    go csys.Run(":12811", func(req cpreq.ChargePointRequest, cpID string) (cpres.ChargePointResponse, error) {
        // Return an error to the Station communicating to the Central System
        //
        // station, isAuthorized := getStation(cpID)
        // if err != nil {
        //   return nil, errors.New("charger not authorized to join network")
        // }

        switch req := req.(type) {
        case *cpreq.BootNotification:
            // accept chargepoint in the network
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
}
```

### ChargePoint Simulator

```go
package main

import (
    "fmt"

    "github.com/eduhenke/go-ocpp"
    "github.com/eduhenke/go-ocpp/cstationsim"
    "github.com/eduhenke/go-ocpp/messages/v1x/cpreq"
    "github.com/eduhenke/go-ocpp/messages/v1x/cpres"
)

func main() {
    stationId := "id01"
    centralSystemURL := "ws://localhost:12811"
    st := cstationsim.NewStation(stationId, centralSystemURL, ocpp.JSON) // or ocpp.SOAP
    rawResp, err := st.CsService.Send(&cpreq.Heartbeat{})
    if err != nil {
        fmt.Println("could't send heartbeat:", err)
    }
    resp, ok := rawResp.(*cpres.Heartbeat)
    if !ok {
        fmt.Println("response is not a heartbeat response")
    }
    fmt.Println("got reply:", resp)
}
```

### Logs

For more useful logging, do:

```go
ocpp.SetDebugLogger(log.New(os.Stdout, "DEBUG:", log.Ltime))
ocpp.SetErrorLogger(log.New(os.Stderr, "ERROR:", log.Ltime))
```
