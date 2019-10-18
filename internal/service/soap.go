package service

import (
	"github.com/eduhenke/go-ocpp/messages"
	"github.com/eduhenke/go-ocpp/soap"
)

type SOAP struct {
	client  *soap.Client
	options *soap.CallOptions
}

func NewSOAP(URL string, options *soap.CallOptions) *SOAP {
	client := soap.NewClient(URL)
	return &SOAP{
		client:  client,
		options: options,
	}
}

func (service *SOAP) Send(req messages.Request) (messages.Response, error) {
	resp := req.GetResponse()
	err := service.client.Call(req.Action(), req, resp, service.options)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
