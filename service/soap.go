package service

import (
	"github.com/eduhenke/go-ocpp/messages"
	"github.com/eduhenke/go-ocpp/soap"
)

type SoapService struct {
	client *soap.Client
}

func NewSoapService(URL string) *SoapService {
	client := soap.NewClient(URL)
	return &SoapService{
		client: client,
	}
}

func (service *SoapService) Send(req messages.Request) (messages.Response, error) {
	resp := req.GetResponse()
	err := service.client.Call(req.Action(), req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}