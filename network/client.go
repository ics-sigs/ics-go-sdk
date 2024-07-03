package network

import (
	"github.com/ics-sigs/ics-go-sdk/client"
	"github.com/ics-sigs/ics-go-sdk/common"
)

type NetworkService struct {
	common.RestAPI
}

func NewNetworkService(c *client.Client) *NetworkService {
	ns := NetworkService{
		common.RestAPI{
			RestAPITripper: c,
		},
	}
	return &ns
}
