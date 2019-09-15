package soap

import (
	"net/http"

	"github.com/eduhenke/go-ocpp/messages"

	"github.com/eduhenke/go-ocpp"

	log "github.com/sirupsen/logrus"
)

func Handler(w http.ResponseWriter, r *http.Request, handle ocpp.MessageHandler) {
	defer r.Body.Close()

	reqEnv, err := logAndUnmarshal(r)

	if err != nil {
		log.WithError(err).Error("couldn't decode request")
		return
	}

	req, ok := reqEnv.Body.Content.(messages.Request)
	if !ok {
		panic("received message is not a request")
	}

	resp, err := handle(req, reqEnv.Header.ChargeBoxIdentity)
	if err != nil {
		log.WithError(err).Error("couldn't handle request")
	}
	rawResp, err := Marshal(resp, err, "http://www.w3.org/2003/05/soap-envelope")
	if err != nil {
		log.WithError(err).Error("couldn't encode request")
	}

	log.
		WithField("bytes", len(rawResp)).
		Trace("Sending response\n", string(rawResp))

	if err != nil {
		log.WithError(err).Error("couldn't encode response")
		return
	}
	w.Header().Set("Content-Type", "application/soap+xml")
	w.Header().Set("Content-Encoding", "deflate")
	w.Write(rawResp)
}
