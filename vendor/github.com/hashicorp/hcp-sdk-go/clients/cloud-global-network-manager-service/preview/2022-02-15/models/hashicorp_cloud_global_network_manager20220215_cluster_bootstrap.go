// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudGlobalNetworkManager20220215ClusterBootstrap hashicorp cloud global network manager 20220215 cluster bootstrap
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.ClusterBootstrap
type HashicorpCloudGlobalNetworkManager20220215ClusterBootstrap struct {

	// bootstrap_expect is the number of consul servers that is expected to join a quorum
	BootstrapExpect int32 `json:"bootstrap_expect,omitempty"`

	// consul_config is the embedded consul configuration.
	ConsulConfig string `json:"consul_config,omitempty"`

	// gossip_key is the consul gossip key for bootstrapping the configuration. Should we move this into a Secrets message?
	GossipKey string `json:"gossip_key,omitempty"`

	// id is the name of the cluster
	ID string `json:"id,omitempty"`

	// server_tls is the public key, private key and certificate_authorities for bootstrapping consul RPC TLS.
	ServerTLS *HashicorpCloudGlobalNetworkManager20220215ServerTLS `json:"server_tls,omitempty"`
}

// Validate validates this hashicorp cloud global network manager 20220215 cluster bootstrap
func (m *HashicorpCloudGlobalNetworkManager20220215ClusterBootstrap) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServerTLS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterBootstrap) validateServerTLS(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerTLS) { // not required
		return nil
	}

	if m.ServerTLS != nil {
		if err := m.ServerTLS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server_tls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("server_tls")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud global network manager 20220215 cluster bootstrap based on the context it is used
func (m *HashicorpCloudGlobalNetworkManager20220215ClusterBootstrap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServerTLS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterBootstrap) contextValidateServerTLS(ctx context.Context, formats strfmt.Registry) error {

	if m.ServerTLS != nil {
		if err := m.ServerTLS.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server_tls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("server_tls")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ClusterBootstrap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ClusterBootstrap) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudGlobalNetworkManager20220215ClusterBootstrap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
