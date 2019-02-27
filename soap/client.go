package soap

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	url string
}

// NewClient creates new SOAP client instance
func NewClient(url string) *Client {
	return &Client{
		url: url,
	}
}

// Call performs HTTP POST request
func (s *Client) Call(soapAction string, request, response interface{}) error {
	if len(s.url) == 0 {
		return errors.New("no URL to request")
	}
	envelope := toSendEnvelope{XMLNS: "http://www.w3.org/2003/05/soap-envelope"}
	envelope.Header = &toSendHeader{
		XMLNS:     "http://www.w3.org/2005/08/addressing",
		Action:    soapAction,
		To:        s.url,
		From:      toSendFrom{Address: "http://192.168.10.1:8080"},
		MessageID: "uuid:" + uuid.New().String(),
	}
	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}
	rawReq := buffer.Bytes()
	log.
		WithField("bytes", len(rawReq)).
		Trace("Sending request\n", string(rawReq))

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/soap+xml")
	req.Close = true

	tr := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, 30*time.Second)
		},
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, _ := ioutil.ReadAll(res.Body)

	if len(rawbody) == 0 && res.StatusCode != http.StatusOK {
		return errors.New("response returned with a non-OK status code: " + res.Status)
	}

	respEnvelope, err := Unmarshal(rawbody)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	respVal, contentVal := reflect.ValueOf(response), reflect.ValueOf(respEnvelope.Body.Content)
	if respVal.Type() != contentVal.Type() {
		panic("incompatible types:" + respVal.Type().String() + contentVal.Type().String())
	}

	respVal.Elem().Set(contentVal.Elem())

	return nil
}
