package volume

import (
    "github.com/ics-sigs/ics-go-sdk/client"
    "github.com/ics-sigs/ics-go-sdk/common"
)

type VolumeService struct {
    common.RestAPI
}

// NewStorageService returns the session's storage service.
func NewVolumeService(c *client.Client) *VolumeService {
    ss := VolumeService{
        common.RestAPI{
            RestAPITripper: c,
        },
    }
    return &ss
}
