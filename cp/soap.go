package cp

import (
	"context"
	"errors"
	"net/http"

	"github.com/eduhenke/go-ocpp/internal/log"
	"github.com/eduhenke/go-ocpp/messages"
	"github.com/eduhenke/go-ocpp/messages/v1x/csreq"
	"github.com/eduhenke/go-ocpp/soap"
)

func handleSoap(ctx context.Context, port string, cshandler CentralSystemMessageHandler) error {
	handler := soapHandler{cshandler}
	server := &http.Server{Addr: port, Handler: handler}

	serverErrorCh := make(chan error)
	go func() {
		serverErrorCh <- server.ListenAndServe()
	}()
	select {
	case <-ctx.Done():
		return nil
	case err := <-serverErrorCh:
		return err
	}
}

type soapHandler struct{ cshandler CentralSystemMessageHandler }

func (s soapHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("New SOAP request")

	err := soap.Handle(w, r, func(request messages.Request, cpID string) (messages.Response, error) {
		req, ok := request.(csreq.CentralSystemRequest)
		if !ok {
			return nil, errors.New("request is not a cprequest")
		}
		return s.cshandler(req)
	})
	if err != nil {
		log.Error("Couldn't handle SOAP request: %w", err)
	}
}
