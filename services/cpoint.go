package services

import (
	msg "github.com/eduhenke/go-ocpp/messages"
	"github.com/eduhenke/go-ocpp/soap"
)

type ChargePointService interface {
	UnlockConnector(request msg.UnlockConnectorRequest) (*msg.UnlockConnectorResponse, error)

	Reset(request msg.ResetRequest) (*msg.ResetResponse, error)

	ChangeAvailability(request msg.ChangeAvailabilityRequest) (*msg.ChangeAvailabilityResponse, error)

	GetDiagnostics(request msg.GetDiagnosticsRequest) (*msg.GetDiagnosticsResponse, error)

	ClearCache(request msg.ClearCacheRequest) (*msg.ClearCacheResponse, error)

	UpdateFirmware(request msg.UpdateFirmwareRequest) (*msg.UpdateFirmwareResponse, error)

	ChangeConfiguration(request msg.ChangeConfigurationRequest) (*msg.ChangeConfigurationResponse, error)

	RemoteStartTransaction(request msg.RemoteStartTransactionRequest) (*msg.RemoteStartTransactionResponse, error)

	RemoteStopTransaction(request msg.RemoteStopTransactionRequest) (*msg.RemoteStopTransactionResponse, error)

	CancelReservation(request msg.CancelReservationRequest) (*msg.CancelReservationResponse, error)

	DataTransfer(request msg.DataTransferRequest) (*msg.DataTransferResponse, error)

	GetConfiguration(request msg.GetConfigurationRequest) (*msg.GetConfigurationResponse, error)

	GetLocalListVersion(request msg.GetLocalListVersionRequest) (*msg.GetLocalListVersionResponse, error)

	ReserveNow(request msg.ReserveNowRequest) (*msg.ReserveNowResponse, error)

	SendLocalList(request msg.SendLocalListRequest) (*msg.SendLocalListResponse, error)
}

type chargePointService struct {
	client *soap.Client
}

func NewChargePointService(client *soap.Client) ChargePointService {
	return &chargePointService{
		client: client,
	}
}

func (service *chargePointService) UnlockConnector(request msg.UnlockConnectorRequest) (*msg.UnlockConnectorResponse, error) {
	response := new(msg.UnlockConnectorResponse)
	err := service.client.Call("/UnlockConnector", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) Reset(request msg.ResetRequest) (*msg.ResetResponse, error) {
	response := new(msg.ResetResponse)
	err := service.client.Call("/Reset", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) ChangeAvailability(request msg.ChangeAvailabilityRequest) (*msg.ChangeAvailabilityResponse, error) {
	response := new(msg.ChangeAvailabilityResponse)
	err := service.client.Call("/ChangeAvailability", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) GetDiagnostics(request msg.GetDiagnosticsRequest) (*msg.GetDiagnosticsResponse, error) {
	response := new(msg.GetDiagnosticsResponse)
	err := service.client.Call("/GetDiagnostics", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) ClearCache(request msg.ClearCacheRequest) (*msg.ClearCacheResponse, error) {
	response := new(msg.ClearCacheResponse)
	err := service.client.Call("/ClearCache", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) UpdateFirmware(request msg.UpdateFirmwareRequest) (*msg.UpdateFirmwareResponse, error) {
	response := new(msg.UpdateFirmwareResponse)
	err := service.client.Call("/UpdateFirmware", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) ChangeConfiguration(request msg.ChangeConfigurationRequest) (*msg.ChangeConfigurationResponse, error) {
	response := new(msg.ChangeConfigurationResponse)
	err := service.client.Call("/ChangeConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) RemoteStartTransaction(request msg.RemoteStartTransactionRequest) (*msg.RemoteStartTransactionResponse, error) {
	response := new(msg.RemoteStartTransactionResponse)
	err := service.client.Call("/RemoteStartTransaction", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) RemoteStopTransaction(request msg.RemoteStopTransactionRequest) (*msg.RemoteStopTransactionResponse, error) {
	response := new(msg.RemoteStopTransactionResponse)
	err := service.client.Call("/RemoteStopTransaction", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) CancelReservation(request msg.CancelReservationRequest) (*msg.CancelReservationResponse, error) {
	response := new(msg.CancelReservationResponse)
	err := service.client.Call("/CancelReservation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) DataTransfer(request msg.DataTransferRequest) (*msg.DataTransferResponse, error) {
	response := new(msg.DataTransferResponse)
	err := service.client.Call("/DataTransfer", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) GetConfiguration(request msg.GetConfigurationRequest) (*msg.GetConfigurationResponse, error) {
	response := new(msg.GetConfigurationResponse)
	err := service.client.Call("/GetConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) GetLocalListVersion(request msg.GetLocalListVersionRequest) (*msg.GetLocalListVersionResponse, error) {
	response := new(msg.GetLocalListVersionResponse)
	err := service.client.Call("/GetLocalListVersion", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) ReserveNow(request msg.ReserveNowRequest) (*msg.ReserveNowResponse, error) {
	response := new(msg.ReserveNowResponse)
	err := service.client.Call("/ReserveNow", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) SendLocalList(request msg.SendLocalListRequest) (*msg.SendLocalListResponse, error) {
	response := new(msg.SendLocalListResponse)
	err := service.client.Call("/SendLocalList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
