package soap

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type HandlerCallback (func(Envelope) (interface{}, error))

func Handler(w http.ResponseWriter, r *http.Request, f HandlerCallback) {
	defer r.Body.Close()

	req, err := logAndUnmarshal(r)

	if err != nil {
		log.WithError(err).Error("couldn't decode request")
		return
	}

	resp, err := f(Envelope(req))
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
