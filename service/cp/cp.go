package cp

import (
	"errors"
	"github.com/eduhenke/go-ocpp/messages/v1x/csreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/csres"
	"github.com/eduhenke/go-ocpp/service"
	"github.com/eduhenke/go-ocpp/soap"
	"github.com/eduhenke/go-ocpp/wsconn"
)

type Service interface {
	Send(csreq.CentralSystemRequest) (csres.CentralSystemResponse, error)
}

type SoapService struct {
	*service.SoapService
}

func NewSoapService(csURL string, options *soap.CallOptions) Service {
	return &SoapService{
		service.NewSoapService(csURL, options),
	}
}

func (service *SoapService) Send(req csreq.CentralSystemRequest) (csres.CentralSystemResponse, error) {
	rawResp, err := service.SoapService.Send(req)
	if err != nil {
		return nil, err
	}
	resp, ok := rawResp.(csres.CentralSystemResponse)
	if !ok {
		return nil, errors.New("response is not a csresponse")
	}
	return resp, nil
}

type JsonService struct {
	*service.JsonService
}

func NewJsonService(conn *wsconn.Conn) Service {
	return &JsonService{service.NewJsonService(conn)}
}

func (service *JsonService) Send(req csreq.CentralSystemRequest) (csres.CentralSystemResponse, error) {
	rawResp, err := service.JsonService.Send(req)
	if err != nil {
		return nil, err
	}
	resp, ok := rawResp.(csres.CentralSystemResponse)
	if !ok {
		return nil, errors.New("response is not a csresponse")
	}
	return resp, nil
}
