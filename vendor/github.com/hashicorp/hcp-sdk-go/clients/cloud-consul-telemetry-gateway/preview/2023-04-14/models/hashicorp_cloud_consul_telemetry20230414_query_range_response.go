// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulTelemetry20230414QueryRangeResponse QueryRangeResponse
//
// swagger:model hashicorp.cloud.consul_telemetry_20230414.QueryRangeResponse
type HashicorpCloudConsulTelemetry20230414QueryRangeResponse struct {

	// result
	Result []*HashicorpCloudConsulTelemetry20230414MetricMatrix `json:"result"`
}

// Validate validates this hashicorp cloud consul telemetry 20230414 query range response
func (m *HashicorpCloudConsulTelemetry20230414QueryRangeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulTelemetry20230414QueryRangeResponse) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(m.Result) { // not required
		return nil
	}

	for i := 0; i < len(m.Result); i++ {
		if swag.IsZero(m.Result[i]) { // not required
			continue
		}

		if m.Result[i] != nil {
			if err := m.Result[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud consul telemetry 20230414 query range response based on the context it is used
func (m *HashicorpCloudConsulTelemetry20230414QueryRangeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulTelemetry20230414QueryRangeResponse) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Result); i++ {

		if m.Result[i] != nil {
			if err := m.Result[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulTelemetry20230414QueryRangeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulTelemetry20230414QueryRangeResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulTelemetry20230414QueryRangeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
