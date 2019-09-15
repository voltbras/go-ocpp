package service

import (
	"github.com/eduhenke/go-ocpp/messages"
	"github.com/eduhenke/go-ocpp/soap"
)

type SoapService struct {
	client *soap.Client
	options *soap.CallOptions
}

func NewSoapService(URL string, options *soap.CallOptions) *SoapService {
	client := soap.NewClient(URL)
	return &SoapService{
		client: client,
		options: options,
	}
}

func (service *SoapService) Send(req messages.Request) (messages.Response, error) {
	resp := req.GetResponse()
	err := service.client.Call(req.Action(), req, resp, service.options)
	if err != nil {
		return nil, err
	}
	return resp, nil
}