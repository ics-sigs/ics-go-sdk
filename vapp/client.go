package vapp

import (
	"github.com/ics-sigs/ics-go-sdk/client"
	"github.com/ics-sigs/ics-go-sdk/common"
)

type VappService struct {
	common.RestAPI
}

func NewVappService(c *client.Client) *VappService {
	vs := VappService{
		common.RestAPI{
			RestAPITripper: c,
		},
	}
	return &vs
}
