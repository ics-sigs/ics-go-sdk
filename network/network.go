package network

import (
	"context"
	"fmt"
	"github.com/ics-sigs/ics-go-sdk/client/methods"
	"github.com/ics-sigs/ics-go-sdk/client/types"
)

func (n *NetworkService) GetNetworkByName(ctx context.Context, name string) (*types.Network, error) {
	networkList, err := methods.GetNetworkList(ctx, n.RestAPITripper)
	if err != nil {
		return nil, err
	}

	for _, network := range networkList {
		if network.Name == name {
			return &network, nil
		}
	}

	return nil, fmt.Errorf("Network not found by name %s", name)
}

func (n *NetworkService) GetNetworkByID(ctx context.Context, networkID string) (*types.Network, error) {
	networkInfo, err := methods.GetNetworkByID(ctx, n.RestAPITripper, networkID)
	return &networkInfo, err
}

func (n *NetworkService) GetNetworkList(ctx context.Context) ([]types.Network, error) {
	networkList, err := methods.GetNetworkList(ctx, n.RestAPITripper)
	return networkList, err
}

func (n *NetworkService) GetSdnNetworkByName(ctx context.Context, name string) (*types.SdnNetwork, error) {
	sdnNetworkList, err := methods.GetSdnNetworkList(ctx, n.RestAPITripper)
	if err != nil {
		return nil, err
	}

	for _, network := range sdnNetworkList {
		if network.Name == name {
			return &network, nil
		}
	}

	return nil, fmt.Errorf("SDN network not found by name %s", name)
}

func (n *NetworkService) GetSdnNetworkByID(ctx context.Context, networkID string) (*types.SdnNetwork, error) {
	sdnNetworkInfo, err := methods.GetSdnNetworkByID(ctx, n.RestAPITripper, networkID)
	return &sdnNetworkInfo, err
}

func (n *NetworkService) GetSdnNetworkList(ctx context.Context) ([]types.SdnNetwork, error) {
	sdnNetworkList, err := methods.GetSdnNetworkList(ctx, n.RestAPITripper)
	return sdnNetworkList, err
}

func (n *NetworkService) GetExtSdnNetworkByID(ctx context.Context, networkID string) (*types.Network, error) {
	networkInfo, err := methods.GetNetworkByID(ctx, n.RestAPITripper, networkID)
	return &networkInfo, err
}
