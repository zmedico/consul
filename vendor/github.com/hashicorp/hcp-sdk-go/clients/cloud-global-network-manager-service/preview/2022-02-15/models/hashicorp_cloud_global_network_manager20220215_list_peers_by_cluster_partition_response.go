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

// HashicorpCloudGlobalNetworkManager20220215ListPeersByClusterPartitionResponse hashicorp cloud global network manager 20220215 list peers by cluster partition response
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.ListPeersByClusterPartitionResponse
type HashicorpCloudGlobalNetworkManager20220215ListPeersByClusterPartitionResponse struct {

	// cluster partition peers
	ClusterPartitionPeers []*HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeers `json:"cluster_partition_peers"`
}

// Validate validates this hashicorp cloud global network manager 20220215 list peers by cluster partition response
func (m *HashicorpCloudGlobalNetworkManager20220215ListPeersByClusterPartitionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterPartitionPeers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ListPeersByClusterPartitionResponse) validateClusterPartitionPeers(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterPartitionPeers) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterPartitionPeers); i++ {
		if swag.IsZero(m.ClusterPartitionPeers[i]) { // not required
			continue
		}

		if m.ClusterPartitionPeers[i] != nil {
			if err := m.ClusterPartitionPeers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster_partition_peers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cluster_partition_peers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud global network manager 20220215 list peers by cluster partition response based on the context it is used
func (m *HashicorpCloudGlobalNetworkManager20220215ListPeersByClusterPartitionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterPartitionPeers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ListPeersByClusterPartitionResponse) contextValidateClusterPartitionPeers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClusterPartitionPeers); i++ {

		if m.ClusterPartitionPeers[i] != nil {
			if err := m.ClusterPartitionPeers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster_partition_peers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cluster_partition_peers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ListPeersByClusterPartitionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ListPeersByClusterPartitionResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudGlobalNetworkManager20220215ListPeersByClusterPartitionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
