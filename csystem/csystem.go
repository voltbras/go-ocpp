package csystem

import (
	"log"
	"net/http"

	"github.com/eduhenke/go-ocpp"
	"github.com/eduhenke/go-ocpp/soap"
)

// EnvelopeHandler handles the OCPP messages coming from the charger
type EnvelopeHandler interface {
	HandleEnvelope(env ocpp.Envelope) (interface{}, error)
}

func ListenAndServe(port string, handler EnvelopeHandler) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: check whether JSON/SOAP and their versions
		soap.Handler(w, r, func(env soap.Envelope) (interface{}, error) {
			return handler.HandleEnvelope((ocpp.Envelope)(env))
		})
	})
	log.Fatal(http.ListenAndServe(port, nil))
}
