package datacenter

import (
    "github.com/ics-sigs/ics-go-sdk/client"
    "github.com/ics-sigs/ics-go-sdk/common"
)

type DatacenterService struct {
    common.RestAPI
}

// NewDatacenterService returns the session's datacenter service.
func NewDatacenterService(c *client.Client) *DatacenterService {
    dc := DatacenterService{
        common.RestAPI{
            RestAPITripper: c,
        },
    }
    return &dc
}