package host

import (
    "github.com/ics-sigs/ics-go-sdk/client"
    "github.com/ics-sigs/ics-go-sdk/common"
)

type HostService struct {
    common.RestAPI
}

// NewDatacenterService returns the session's host service.
func NewHostService(c *client.Client) *HostService {
    ht := HostService{
        common.RestAPI{
            RestAPITripper: c,
        },
    }

    return &ht
}