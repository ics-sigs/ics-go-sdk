package vapp

import (
	"context"
	"fmt"
	"github.com/ics-sigs/ics-go-sdk/client"
	"github.com/ics-sigs/ics-go-sdk/client/methods"
	"github.com/ics-sigs/ics-go-sdk/client/types"
)

func (v *VappService) GetVappList(ctx context.Context) ([]types.Vapp, error) {
	vappList, err := methods.GetVappList(ctx, v.RestAPITripper)
	return vappList, err
}

func (v *VappService) GetVappByName(ctx context.Context, name string) (*types.Vapp, error) {
	vappList, err := methods.GetVappList(ctx, v.RestAPITripper)
	if err != nil {
		return nil, err
	}
	for _, vapp := range vappList {
		if vapp.Name == name {
			return &vapp, nil
		}
	}
	return nil, fmt.Errorf("Vapp not found by name %s", name)
}

func (v *VappService) CreateVapp(ctx context.Context, req types.VappCreateReq) (types.Task, error) {
	task, err := methods.CreateVapp(ctx, v.RestAPITripper, req)
	return task, err
}

func (v *VappService) DeleteVapp(ctx context.Context, vappID string) (types.Task, error) {
	task, err := methods.DeleteVapp(ctx, v.RestAPITripper, vappID)
	return task, err
}

func (v *VappService) DeleteVappWithCheckParams(ctx context.Context, vappID string, passwd string) (types.Task, error) {
	checkParams, err := methods.GenerateCheckParams(ctx, v.RestAPITripper, passwd)
	if err != nil {
		return types.Task{}, err
	}

	restClient := v.RestAPITripper.(*client.Client)
	restClient.SetCheckParams(checkParams)
	defer restClient.SetCheckParams("")

	return v.DeleteVapp(ctx, vappID)
}

func (v *VappService) AddVmToVapp(ctx context.Context, vappID string, vmID []string) (types.Task, error) {
	task, err := methods.AddVmToVapp(ctx, v.RestAPITripper, vappID, vmID)
	return task, err
}

func (v *VappService) DeleteVmFromVapp(ctx context.Context, vappID string, vmID []string) (types.Task, error) {
	task, err := methods.DeleteVmFromVapp(ctx, v.RestAPITripper, vappID, vmID)
	return task, err
}

func (v *VappService) PowerOnVapp(ctx context.Context, vappID string) (types.Task, error) {
	task, err := methods.PowerOnVapp(ctx, v.RestAPITripper, vappID)
	return task, err
}

func (v *VappService) PowerOffVapp(ctx context.Context, vappID string) (types.Task, error) {
	task, err := methods.PowerOffVapp(ctx, v.RestAPITripper, vappID)
	return task, err
}

func (v *VappService) PowerOffVappSafely(ctx context.Context, vappID string) (types.Task, error) {
	task, err := methods.PowerOffVappSafely(ctx, v.RestAPITripper, vappID)
	return task, err
}

func (v *VappService) RestartVapp(ctx context.Context, vappID string) (types.Task, error) {
	task, err := methods.RestartVapp(ctx, v.RestAPITripper, vappID)
	return task, err
}
