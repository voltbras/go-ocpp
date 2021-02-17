package soap

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/voltbras/go-ocpp/internal/log"

	"github.com/voltbras/go-ocpp/messages"

	"github.com/voltbras/go-ocpp"
)

func Handle(w http.ResponseWriter, r *http.Request, handle ocpp.MessageHandler) error {
	defer r.Body.Close()

	rawReq, _ := ioutil.ReadAll(r.Body)
	log.Debug("Received request with %d bytes from: %s", len(rawReq), r.RemoteAddr)
	log.Debug("Received raw request: %s", string(rawReq))
	reqEnv, err := Unmarshal(rawReq)

	if err != nil {
		return fmt.Errorf("couldn't decode request %w", err)
	}

	req, ok := reqEnv.Body.Content.(messages.Request)
	if !ok {
		return errors.New("received message is not a request")
	}

	resp, err := handle(req, reqEnv.Header.ChargeBoxIdentity)
	if err != nil {
		log.Error("couldn't handle request: %w", err)
	}
	rawResp, err := Marshal(resp, err, "http://www.w3.org/2003/05/soap-envelope")
	if err != nil {
		return fmt.Errorf("couldn't encode response: %w", err)
	}

	log.Debug("Sending response with %d bytes", len(rawResp))
	log.Debug("Sending raw response: %s", string(rawResp))

	w.Header().Set("Content-Type", "application/soap+xml")
	w.Header().Set("Content-Encoding", "deflate")

	_, err = w.Write(rawResp)
	return err
}
