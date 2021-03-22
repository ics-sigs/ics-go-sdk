package cluster

import (
	"github.com/inspur-ics/ics-go-sdk/client"
	"github.com/inspur-ics/ics-go-sdk/common"
)

type ClusterService struct {
	common.RestAPI
}

func NewClusterService(c *client.Client) *ClusterService {
	ht := ClusterService{
		common.RestAPI{
			RestAPITripper: c,
		},
	}
	return &ht
}
