package cluster

import (
	"context"
	"fmt"
	"github.com/ics-sigs/ics-go-sdk/client/methods"
	"github.com/ics-sigs/ics-go-sdk/client/types"
	//    "k8s.io/klog"
)

func (c *ClusterService) GetClusterList(ctx context.Context) ([]types.Cluster, error) {
	clusterList, err := methods.GetClusterList(ctx, c.RestAPITripper)
	return clusterList, err
}

func (c *ClusterService) GetClusterByName(ctx context.Context, name string) (*types.Cluster, error) {
	clusterList, err := methods.GetClusterList(ctx, c.RestAPITripper)
	if err != nil {
		return nil, err
	}
	for _, cluster := range clusterList {
		if cluster.Name == name {
			return &cluster, nil
		}
	}
	return nil, fmt.Errorf("Cluster not found by name %s", name)
}
