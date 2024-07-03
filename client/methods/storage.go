package methods

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ics-sigs/ics-go-sdk/client/restful"
	"github.com/ics-sigs/ics-go-sdk/client/types"
)

func GetStoragePageList(ctx context.Context, r restful.RestAPITripper, req *types.StoragePageReq) (*types.StoragePageResponse, error) {
	var api types.ICSApi
	var response = types.StoragePageResponse{}

	api.Api = "/storages/"
	api.Token = true

	resp, err := r.GetTrip(ctx, api, req)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return &response, err
}

func GetStorageInfo(ctx context.Context, r restful.RestAPITripper, storageId string) (*types.Storage, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.Storage{}

	if len(storageId) <= 0 {
		storageId = "anonymous"
	}
	api.Api = fmt.Sprintf("/storages/%s", storageId)
	api.Token = true

	resp, err := r.GetTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return &response, err
}
func GetStorageList(ctx context.Context, r restful.RestAPITripper) (*types.StoragePageResponse, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.StoragePageResponse{}

	api.Api = "/storages/"
	api.Token = true

	resp, err := r.GetTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return &response, err
}

func GetImageStorageList(ctx context.Context, r restful.RestAPITripper) (*types.StoragePageResponse, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.StoragePageResponse{}

	api.Api = "/storages?extype=IMAGE_STORAGE"
	api.Token = true

	resp, err := r.GetTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return &response, err
}

func GetImageFileList(ctx context.Context, r restful.RestAPITripper, storageId string) (*types.ImageFilePageResponse, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.ImageFilePageResponse{}

	api.Api = fmt.Sprintf("/storages/%s/files", storageId)
	api.Token = true

	resp, err := r.GetTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return &response, err
}
