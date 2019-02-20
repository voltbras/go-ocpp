package cstationsim

import (
	log "github.com/sirupsen/logrus"
	msg "github.com/eduhenke/go-ocpp/messages"
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

func (st *Station) Boot() {
	reply, err := st.CsService.BootNotification(msg.BootNotificationRequest{
		ChargePointVendor:       "Vendor",
		ChargePointModel:        "Simulator-01",
		ChargePointSerialNumber: "1337",
		ChargeBoxSerialNumber:   "1337",
		FirmwareVersion:         "1.7.0",
	})
	if err != nil {
		log.WithError(err).Error("could't send boot notification")
		return
	}
	log.Info("got reply:", reply)
}
