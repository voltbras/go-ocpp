package cp

import (
	"github.com/eduhenke/go-ocpp/messages/v16/csreq"
	"github.com/eduhenke/go-ocpp/messages/v16/csres"
	"github.com/eduhenke/go-ocpp/soap"
)

type Service interface {
	Send(csreq.CentralSystemRequest) (csres.CentralSystemResponse, error)
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

func (service *SoapService) Send(req csreq.CentralSystemRequest) (csres.CentralSystemResponse, error) {
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

// func (service *JSONService) Send(req csreq.CentralSystemRequest) (csres.CentralSystemResponse, error) {
// 	resp := req.GetResponseObject()
// 	err := service.client.Call(req.Action(), req, resp)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp, nil
// }
