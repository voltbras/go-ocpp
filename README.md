# go-ocpp

OCPP(1.5/1.6) implementation in Golang.

- v1.5, it's assumed it is SOAP
- v1.6, it's assumed it is JSON

## Usage

### Central System

Just pass a handler that takes in a `cpreq.ChargePointRequest` and returns a `(cpresp.ChargePointResponse, error)`.

In SOAP, error messages will be sent back via Fault, as specified in OCPP v1.5

In Websockets, error messages will be sent back as specified in OCPP-J v1.6

```go
csys := cs.New()
go csys.Run(":12811", func(req cpreq.ChargePointRequest, metadata cs.ChargePointRequestMetadata) (cpresp.ChargePointResponse, error) {
    // Return an error to the Station communicating to the Central System
    //
    // station, isAuthorized := getStation(metadata.ChargePointID)
    // if !isAuthorized {
    //   return nil, errors.New("charger not authorized to join network")
    // }
    // ---
    // Or check some specific header in the underlying HTTP request:
    // 
    // if shouldBlock(metadata.HTTPRequest) {
    //   return nil, errors.New("charger should send appropriate headers")
    // }

    switch req := req.(type) {
    case *cpreq.BootNotification:
        // accept chargepoint in the network
        return &cpresp.BootNotification{
            Status:      "Accepted",
            CurrentTime: time.Now(),
            Interval:    60,
        }, nil

    case *cpreq.Heartbeat:
        return &cpresp.Heartbeat{CurrentTime: time.Now()}, nil

    case *cpreq.StatusNotification:
        if req.Status != "Available" {
            // chargepoint is unavailable
        }
        return &cpresp.StatusNotification{}, nil

    default:
        return nil, errors.New("Response not supported")
    }
}
```

### Charge Point

Pass the required parameters to the constructor function, and then just send any request(`cpreq.*`).

TODO: assertion of the response type should be done inside the `.Send`?

```go
stationID := "id01"
centralSystemURL := "ws://localhost:12811"
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
```

### Logs

For more useful logging, do:

```go
ocpp.SetDebugLogger(log.New(os.Stdout, "DEBUG:", log.Ltime))
ocpp.SetErrorLogger(log.New(os.Stderr, "ERROR:", log.Ltime))
```
