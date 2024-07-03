package storage

import (
    "github.com/ics-sigs/ics-go-sdk/client"
    "github.com/ics-sigs/ics-go-sdk/common"
)

type StorageService struct {
    common.RestAPI
}

// NewStorageService returns the session's storage service.
func NewStorageService(c *client.Client) *StorageService {
    ss := StorageService{
        common.RestAPI{
            RestAPITripper: c,
        },
    }
    return &ss
}
