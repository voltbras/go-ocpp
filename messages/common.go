package messages

import (
	"encoding/xml"
	"time"
)

type ChargeBoxIdentity string

// Type of string defining identification token, e.g. RFID or credit card number. To be treated as case insensitive.
type IdToken string

// Defines the authorization-status-value
type AuthorizationStatus string

const (
	AuthorizationStatusAccepted AuthorizationStatus = "Accepted"

	AuthorizationStatusBlocked AuthorizationStatus = "Blocked"

	AuthorizationStatusExpired AuthorizationStatus = "Expired"

	AuthorizationStatusInvalid AuthorizationStatus = "Invalid"

	AuthorizationStatusConcurrentTx AuthorizationStatus = "ConcurrentTx"
)

// Defines the status returned in DataTransfer.conf
type DataTransferStatus string

const (
	DataTransferStatusAccepted DataTransferStatus = "Accepted"

	DataTransferStatusRejected DataTransferStatus = "Rejected"

	DataTransferStatusUnknownMessageId DataTransferStatus = "UnknownMessageId"

	DataTransferStatusUnknownVendorId DataTransferStatus = "UnknownVendorId"
)

type IdTagInfo struct {
	Status AuthorizationStatus `xml:"status,omitempty"`

	ExpiryDate time.Time `xml:"expiryDate,omitempty"`

	ParentIdTag IdToken `xml:"parentIdTag,omitempty"`
}

type Location string

type DataTransferRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ dataTransferRequest"`

	VendorId string `xml:"vendorId,omitempty"`

	MessageId string `xml:"messageId,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DataTransferResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ dataTransferResponse"`

	Status *DataTransferStatus `xml:"status,omitempty"`

	Data string `xml:"data,omitempty"`
}
