// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudLocationRegion Region identifies a Cloud data-plane region.
//
// swagger:model hashicorp.cloud.location.Region
type HashicorpCloudLocationRegion struct {

	// provider is the named cloud provider ("aws", "gcp", "azure")
	Provider string `json:"provider,omitempty"`

	// region is the cloud region ("us-west1", "us-east1")
	Region string `json:"region,omitempty"`
}

// Validate validates this hashicorp cloud location region
func (m *HashicorpCloudLocationRegion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud location region based on context it is used
func (m *HashicorpCloudLocationRegion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudLocationRegion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudLocationRegion) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudLocationRegion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
