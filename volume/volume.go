package volume

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"github.com/ics-sigs/ics-go-sdk/client/methods"
	"github.com/ics-sigs/ics-go-sdk/client/types"
	"github.com/ics-sigs/ics-go-sdk/client"
)

func (volume *VolumeService) CreateVolume(ctx context.Context, volumereq types.VolumeReq) (types.Task, error) {
	task, err := methods.CreateVolume(ctx, volume.RestAPITripper, volumereq)
	return task, err
}

func (vs *VolumeService) DeleteVolume(ctx context.Context, volumeid string, deleteVolume bool) (types.Task, error) {
	task, err := methods.DeleteVolume(ctx, vs.RestAPITripper, volumeid, deleteVolume)
	return task, err
}

func (vs *VolumeService) DeleteVolumeWithCheckParams(ctx context.Context, volumeid string, deleteVolume bool,
	passwd string) (types.Task, error) {
	checkParams, err := vs.generateCheckParams(ctx, passwd)
	if err != nil {
		return types.Task{}, err
	}

	client := vs.RestAPITripper.(*client.Client)
	client.SetCheckParams(checkParams)
	defer client.SetCheckParams("")

	return vs.DeleteVolume(ctx, volumeid, deleteVolume)
}

func (vs *VolumeService) generateCheckParams(ctx context.Context, params string) (string, error) {
	publicKey, err := methods.GetPublicKey(ctx, vs.RestAPITripper)
	if err != nil {
		return "", fmt.Errorf("Failed to get public key. Err: %v", err)
	}

	decPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return "", fmt.Errorf("Failed to decode public key. Err: %v", err)
	}

	pubInterface, err := x509.ParsePKIXPublicKey(decPublicKey)
	if err != nil {
		return "", fmt.Errorf("Failed to parse public key. Err: %v", err)
	}
	pub := pubInterface.(*rsa.PublicKey)
	encParams, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(params))
	if err != nil {
		return "", fmt.Errorf("Failed to encrypt params %q. Err: %v", params, err)
	}

	return base64.StdEncoding.EncodeToString(encParams), nil
}

func (vs *VolumeService) SetVolume(ctx context.Context, volumeid string, volume types.Volume) (types.Task, error) {
	task, err := methods.SetVolume(ctx, vs.RestAPITripper, volumeid, volume)
	return task, err
}

func (vs *VolumeService) GetVolumeInfoById(ctx context.Context, volumeid string) (types.Volume, error) {
	volumeinfo, err := methods.GetVolumeInfoById(ctx, vs.RestAPITripper, volumeid)
	return volumeinfo, err
}

func (vs *VolumeService) GetVolumesInDatastore(ctx context.Context, storeid string) ([]types.Volume, error) {
	volumes, err := methods.GetVolumesInDatastore(ctx, vs.RestAPITripper, storeid)
	return volumes.Items, err
}
