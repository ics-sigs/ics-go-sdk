package methods

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ics-sigs/ics-go-sdk/client/restful"
	"github.com/ics-sigs/ics-go-sdk/client/types"
)

func CreateVolume(ctx context.Context, r restful.RestAPITripper, volume types.VolumeReq) (types.Task, error) {
	//var reqBody      *types.Common
	var api types.ICSApi
	var response = types.Task{}

	api.Api = fmt.Sprintf("/volumes")
	api.Token = true

	resp, err := r.PostTrip(ctx, api, volume)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return response, err
}

func DeleteVolume(ctx context.Context, r restful.RestAPITripper, volumeId string, deleteVolume bool) (types.Task, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.Task{}

	api.Api = fmt.Sprintf("volumes/%s/?removeData=%v", volumeId, deleteVolume)
	api.Token = true

	resp, err := r.DeleteTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return response, err
}

func SetVolume(ctx context.Context, r restful.RestAPITripper, volumeId string, volume types.Volume) (types.Task, error) {
	var api types.ICSApi
	var response = types.Task{}

	api.Api = fmt.Sprintf("/volumes/%s", volumeId)
	api.Token = true

	resp, err := r.PutTrip(ctx, api, volume)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return response, err
}

func GetVolumeInfoById(ctx context.Context, r restful.RestAPITripper, volumeId string) (types.Volume, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.Volume{}

	if len(volumeId) <= 0 {
		volumeId = "anonymous"
	}
	api.Api = fmt.Sprintf("/volumes/%s", volumeId)
	api.Token = true

	resp, err := r.GetTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return response, err
}

func GetVolumesInDatastore(ctx context.Context, r restful.RestAPITripper, datastoreId string) (types.VolumeListRsp, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.VolumeListRsp{}
	if len(datastoreId) <= 0 {
		datastoreId = "anonymous"
	}
	api.Api = fmt.Sprintf("storages/%s/volumes", datastoreId)
	api.Token = true

	resp, err := r.GetTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return response, err
}
