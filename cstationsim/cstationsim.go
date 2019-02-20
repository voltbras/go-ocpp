package cstationsim

import (
	"github.com/eduhenke/go-ocpp/services"
	soap "github.com/eduhenke/go-ocpp/soap"
)

type Station struct {
	client    *soap.Client
	CsService services.CentralSystemService
}

func NewStation(port, csURL string) *Station {
	client := soap.NewClient(csURL)
	return &Station{
		client:    client,
		CsService: services.NewCentralSystemService(client),
	}
}
