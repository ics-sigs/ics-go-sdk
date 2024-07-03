package vm

import (
    "github.com/ics-sigs/ics-go-sdk/client"
    "github.com/ics-sigs/ics-go-sdk/common"
)

type VirtualMachineService struct {
    common.RestAPI
}

// NewVirtualMachineService returns the session's virtual machine service.
func NewVirtualMachineService(c *client.Client) *VirtualMachineService {
    vm := VirtualMachineService{
        common.RestAPI{
            RestAPITripper: c,
        },
    }

    return &vm
}