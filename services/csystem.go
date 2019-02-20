package services

import (
	msg "github.com/eduhenke/go-ocpp/messages"
	"github.com/eduhenke/go-ocpp/soap"
)

type CentralSystemService interface {
	Authorize(request msg.AuthorizeRequest) (*msg.AuthorizeResponse, error)

	StartTransaction(request msg.StartTransactionRequest) (*msg.StartTransactionResponse, error)

	StopTransaction(request msg.StopTransactionRequest) (*msg.StopTransactionResponse, error)

	Heartbeat(request msg.HeartbeatRequest) (*msg.HeartbeatResponse, error)

	MeterValues(request msg.MeterValuesRequest) (*msg.MeterValuesResponse, error)

	BootNotification(request msg.BootNotificationRequest) (*msg.BootNotificationResponse, error)

	StatusNotification(request msg.StatusNotificationRequest) (*msg.StatusNotificationResponse, error)

	FirmwareStatusNotification(request msg.FirmwareStatusNotificationRequest) (*msg.FirmwareStatusNotificationResponse, error)

	DiagnosticsStatusNotification(request msg.DiagnosticsStatusNotificationRequest) (*msg.DiagnosticsStatusNotificationResponse, error)

	DataTransfer(request msg.DataTransferRequest) (*msg.DataTransferResponse, error)
}

type centralSystemService struct {
	client *soap.Client
}

func NewCentralSystemService(client *soap.Client) CentralSystemService {
	return &centralSystemService{
		client: client,
	}
}

func (service centralSystemService) Authorize(request msg.AuthorizeRequest) (*msg.AuthorizeResponse, error) {
	response := new(msg.AuthorizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service centralSystemService) StartTransaction(request msg.StartTransactionRequest) (*msg.StartTransactionResponse, error) {
	response := new(msg.StartTransactionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service centralSystemService) StopTransaction(request msg.StopTransactionRequest) (*msg.StopTransactionResponse, error) {
	response := new(msg.StopTransactionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service centralSystemService) Heartbeat(request msg.HeartbeatRequest) (*msg.HeartbeatResponse, error) {
	response := new(msg.HeartbeatResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service centralSystemService) MeterValues(request msg.MeterValuesRequest) (*msg.MeterValuesResponse, error) {
	response := new(msg.MeterValuesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service centralSystemService) BootNotification(request msg.BootNotificationRequest) (*msg.BootNotificationResponse, error) {
	response := new(msg.BootNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service centralSystemService) StatusNotification(request msg.StatusNotificationRequest) (*msg.StatusNotificationResponse, error) {
	response := new(msg.StatusNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service centralSystemService) FirmwareStatusNotification(request msg.FirmwareStatusNotificationRequest) (*msg.FirmwareStatusNotificationResponse, error) {
	response := new(msg.FirmwareStatusNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service centralSystemService) DiagnosticsStatusNotification(request msg.DiagnosticsStatusNotificationRequest) (*msg.DiagnosticsStatusNotificationResponse, error) {
	response := new(msg.DiagnosticsStatusNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service centralSystemService) DataTransfer(request msg.DataTransferRequest) (*msg.DataTransferResponse, error) {
	response := new(msg.DataTransferResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
