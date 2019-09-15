package cp

import (
	"github.com/eduhenke/go-ocpp/messages/v16/cpreq"
	"github.com/eduhenke/go-ocpp/messages/v16/cpres"
	"github.com/eduhenke/go-ocpp/soap"
)

type Service interface {
	Send(cpreq.ChargePointRequest) (cpres.ChargePointResponse, error)
}

type SoapService struct {
	client *soap.Client
}

func NewSoapService(csURL string) Service {
	client := soap.NewClient(csURL)
	return &SoapService{
		client: client,
	}
}

func (service *SoapService) Send(req cpreq.ChargePointRequest) (cpres.ChargePointResponse, error) {
	resp := req.GetResponseObject()
	err := service.client.Call(req.Action(), req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// type JSONService struct {
// }

// func NewJSONService() Service {
// 	return &JSONService{}
// }

// func (service *JSONService) Send(req cpreq.ChargePointRequest) (cpres.ChargePointResponse, error) {
// 	resp := req.GetResponseObject()
// 	err := service.client.Call(req.Action(), req, resp)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp, nil
// }
