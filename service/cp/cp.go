package cp

import (
	"errors"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/cpres"
	"github.com/eduhenke/go-ocpp/service"
	"github.com/eduhenke/go-ocpp/wsconn"
)

type Service interface {
	Send(cpreq.ChargePointRequest) (cpres.ChargePointResponse, error)
}

type SoapService struct {
	*service.SoapService
}

func NewSoapService(csURL string) Service {
	return &SoapService{
		service.NewSoapService(csURL),
	}
}

func (service *SoapService) Send(req cpreq.ChargePointRequest) (cpres.ChargePointResponse, error) {
	rawResp, err := service.SoapService.Send(req)
	if err != nil {
		return nil, err
	}
	resp, ok := rawResp.(cpres.ChargePointResponse)
	if !ok {
		return nil, errors.New("response is not a cpresponse")
	}
	return resp, nil
}

type JsonService struct {
	*service.JsonService
}

func NewJsonService(conn *wsconn.Conn) Service {
	return &JsonService{service.NewJsonService(conn)}
}

func (service *JsonService) Send(req cpreq.ChargePointRequest) (cpres.ChargePointResponse, error) {
	rawResp, err := service.JsonService.Send(req)
	if err != nil {
		return nil, err
	}
	resp, ok := rawResp.(cpres.ChargePointResponse)
	if !ok {
		return nil, errors.New("response is not a cpresponse")
	}
	return resp, nil
}

