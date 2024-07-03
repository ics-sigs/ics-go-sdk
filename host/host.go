package host

import (
	"context"
	"github.com/ics-sigs/ics-go-sdk/client/methods"
	"github.com/ics-sigs/ics-go-sdk/client/types"
	//    "k8s.io/klog"
)

func (h *HostService) GetHost(ctx context.Context, hostUUID string) (*types.Host, error) {
	host, err := methods.GetHostById(ctx, h.RestAPITripper, hostUUID)
	return host, err
}

func (h *HostService) GetHostHealthInfo(ctx context.Context, hostUUID string) (*types.HostHealthInfo, error) {
	healthInfo, err := methods.GetHostHealthInfoById(ctx, h.RestAPITripper, hostUUID)
	return healthInfo, err
}

func (h *HostService) GetHostListByDC(ctx context.Context, datacenterPath string) ([]*types.Host, error) {
	var hostslist []*types.Host
	hosts, err := methods.GetHostListByDC(ctx, h.RestAPITripper, datacenterPath)
	if err == nil && len(hosts.Items) > 0 {
		for index := range hosts.Items {
			hostslist = append(hostslist, &hosts.Items[index])
		}
	}
	return hostslist, err
}

func (h *HostService) GetHostListByClusterID(ctx context.Context, clusterID string) ([]*types.Host, error) {
	var hostslist []*types.Host
	hostResp, err := methods.GetHostList(ctx, h.RestAPITripper)
	if err != nil {
		return hostslist, err
	}
	for i := range hostResp.Items {
		if hostResp.Items[i].ClusterID == clusterID {
			hostslist = append(hostslist, &hostResp.Items[i])
		}
	}
	return hostslist, err
}

func (h *HostService) GetHostListByClusterName(ctx context.Context, clusterName string) ([]*types.Host, error) {
	var hostslist []*types.Host
	hostResp, err := methods.GetHostList(ctx, h.RestAPITripper)
	if err != nil {
		return hostslist, err
	}
	for i := range hostResp.Items {
		if hostResp.Items[i].ClusterName == clusterName {
			hostslist = append(hostslist, &hostResp.Items[i])
		}
	}
	return hostslist, err
}

func (h *HostService) GetHostAvailStorages(ctx context.Context, hostUUID string) ([]*types.Storage, error) {
	var storageList []*types.Storage
	storages, err := methods.GetHostAvailStorages(ctx, h.RestAPITripper, hostUUID)

	for idx := range storages {
		storageList = append(storageList, &storages[idx])
	}

	return storageList, err
}

func (h *HostService) GetHostList(ctx context.Context) ([]types.Host, error) {
	hostlist, err := methods.GetHostList(ctx, h.RestAPITripper)
	return hostlist.Items, err
}

func (h *HostService) GetHostListByStorageID(ctx context.Context, storageID string) ([]types.Host, error) {
	hostlist, err := methods.GetHostListByStorageID(ctx, h.RestAPITripper, storageID)
	return hostlist.Items, err
}

func (h *HostService) GetAvailHostListByStorageID(ctx context.Context, storageID string) ([]types.Host, error) {
	hostlist, err := methods.GetAvailHostListByStorageID(ctx, h.RestAPITripper, storageID)
	return hostlist, err
}

func (h *HostService) GetHostListBySwitchID(ctx context.Context, switchID string) ([]types.Host, error) {
	hostlist, err := methods.GetHostListBySwitchID(ctx, h.RestAPITripper, switchID)
	return hostlist.Items, err
}

func (h *HostService) GetHostListByNetworkID(ctx context.Context, networkID string) ([]types.Host, error) {
	var hostList []types.Host
	network, err := methods.GetNetworkByID(ctx, h.RestAPITripper, networkID)
	if err != nil {
		return hostList, err
	}

	hostPageList, err := methods.GetHostListBySwitchID(ctx, h.RestAPITripper, network.VswitchDto.ID)
	return hostPageList.Items, err
}

func (h *HostService) GetHostListByNetworkName(ctx context.Context, networkName string) ([]types.Host, error) {
	var network types.Network
	var hostList []types.Host
	networkList, err := methods.GetNetworkList(ctx, h.RestAPITripper)
	if err != nil {
		return hostList, err
	}

	found := false
	for _, network = range networkList {
		if network.Name == networkName {
			found = true
			break
		}
	}
	if !found {
		return hostList, err
	}

	hostPageList, err := methods.GetHostListBySwitchID(ctx, h.RestAPITripper, network.VswitchDto.ID)
	return hostPageList.Items, err
}

func (h *HostService) GetHostAccessibleDatastoreList(ctx context.Context, hostid string) ([]types.Storage, error) {
	storagelist, err := methods.GetHostAccessibleDatastoreList(ctx, h.RestAPITripper, hostid)
	return storagelist, err
}

func (h *HostService) GetHostListByExtSdnNetworkID(ctx context.Context, networkID string) ([]types.Host, error) {
	var hostList []types.Host
	network, err := methods.GetNetworkByID(ctx, h.RestAPITripper, networkID)
	if err != nil {
		return hostList, err
	}
	if len(network.VswitchDto.HostDtos) >= 1 {
		hostList = network.VswitchDto.HostDtos
	}

	return hostList, nil
}