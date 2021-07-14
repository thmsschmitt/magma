// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkProbeDestinationDetails network probe destination details
// swagger:model network_probe_destination_details
type NetworkProbeDestinationDetails struct {

	// destination tls client certificate.
	// Required: true
	// Format: byte
	Certificate *strfmt.Base64 `json:"certificate"`

	// delivery address
	// Required: true
	DeliveryAddress string `json:"delivery_address"`

	// delivery type
	// Required: true
	// Enum: [all events_only]
	DeliveryType string `json:"delivery_type"`

	// destination tls client private key.
	// Required: true
	// Format: byte
	PrivateKey *strfmt.Base64 `json:"private_key"`

	// enables exporter to skip server certs verification.
	// Required: true
	SkipVerifyServer bool `json:"skip_verify_server"`
}

// Validate validates this network probe destination details
func (m *NetworkProbeDestinationDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeliveryAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeliveryType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkipVerifyServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkProbeDestinationDetails) validateCertificate(formats strfmt.Registry) error {

	if err := validate.Required("certificate", "body", m.Certificate); err != nil {
		return err
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *NetworkProbeDestinationDetails) validateDeliveryAddress(formats strfmt.Registry) error {

	if err := validate.RequiredString("delivery_address", "body", string(m.DeliveryAddress)); err != nil {
		return err
	}

	return nil
}

var networkProbeDestinationDetailsTypeDeliveryTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","events_only"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkProbeDestinationDetailsTypeDeliveryTypePropEnum = append(networkProbeDestinationDetailsTypeDeliveryTypePropEnum, v)
	}
}

const (

	// NetworkProbeDestinationDetailsDeliveryTypeAll captures enum value "all"
	NetworkProbeDestinationDetailsDeliveryTypeAll string = "all"

	// NetworkProbeDestinationDetailsDeliveryTypeEventsOnly captures enum value "events_only"
	NetworkProbeDestinationDetailsDeliveryTypeEventsOnly string = "events_only"
)

// prop value enum
func (m *NetworkProbeDestinationDetails) validateDeliveryTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, networkProbeDestinationDetailsTypeDeliveryTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NetworkProbeDestinationDetails) validateDeliveryType(formats strfmt.Registry) error {

	if err := validate.RequiredString("delivery_type", "body", string(m.DeliveryType)); err != nil {
		return err
	}

	// value enum
	if err := m.validateDeliveryTypeEnum("delivery_type", "body", m.DeliveryType); err != nil {
		return err
	}

	return nil
}

func (m *NetworkProbeDestinationDetails) validatePrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("private_key", "body", m.PrivateKey); err != nil {
		return err
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *NetworkProbeDestinationDetails) validateSkipVerifyServer(formats strfmt.Registry) error {

	if err := validate.Required("skip_verify_server", "body", bool(m.SkipVerifyServer)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkProbeDestinationDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkProbeDestinationDetails) UnmarshalBinary(b []byte) error {
	var res NetworkProbeDestinationDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
