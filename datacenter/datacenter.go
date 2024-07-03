package datacenter

import (
	"context"
	"github.com/ics-sigs/ics-go-sdk/client/methods"
	"github.com/ics-sigs/ics-go-sdk/client/types"
)

func (d *DatacenterService) GetDatacenter(ctx context.Context, datacenterid string) (*types.Datacenter, error) {
	dc, err := methods.GetDatacenterById(ctx, d.RestAPITripper, datacenterid)
	return dc, err
}

func (d *DatacenterService) GetDatacenterByName(ctx context.Context, datacenterName string) (*types.Datacenter, error) {
	dc, err := methods.GetDatacenterByName(ctx, d.RestAPITripper, datacenterName)
	return dc, err
}

func (d *DatacenterService) GetAllDatacenters(ctx context.Context) ([]*types.Datacenter, error) {
	var dclist []*types.Datacenter
	datacenters, err := methods.GetAllDatacenterList(ctx, d.RestAPITripper)
	for idx := range datacenters.Items {
		dclist = append(dclist, &datacenters.Items[idx])
	}
	return dclist, err
}

func (d *DatacenterService) GetDatacenterVMList(ctx context.Context, datacenterid string) ([]*types.VirtualMachine, error) {

	var vmlist []*types.VirtualMachine
	dcvms, err := methods.GetDatacenterVMById(ctx, d.RestAPITripper, datacenterid)

	for idx := range dcvms.Items {
		vmlist = append(vmlist, &dcvms.Items[idx])
	}

	return vmlist, err
}
