package methods

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ics-sigs/ics-go-sdk/client/restful"
	"github.com/ics-sigs/ics-go-sdk/client/types"
)

func GetVappList(ctx context.Context, r restful.RestAPITripper) ([]types.Vapp, error) {
	var api types.ICSApi
	var response = types.VappListRsp{}
	var reqBody *types.Common
	api.Api = "/vclusters"
	api.Token = true
	resp, err := r.GetTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}
	return response.Items, err
}

func CreateVapp(ctx context.Context, r restful.RestAPITripper, req types.VappCreateReq) (types.Task, error) {
	var api types.ICSApi
	var response = types.Task{}

	api.Api = "/vclusters"
	api.Token = true

	resp, err := r.PostTrip(ctx, api, req)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return response, err
}

func DeleteVapp(ctx context.Context, r restful.RestAPITripper, vappID string) (types.Task, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.Task{}

	api.Api = fmt.Sprintf("/vclusters/%s", vappID)
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

func AddVmToVapp(ctx context.Context, r restful.RestAPITripper, vappID string, vmID []string) (types.Task, error) {
	var api types.ICSApi
	var response = types.Task{}

	api.Api = fmt.Sprintf("/vclusters/%s/vms?action=shiftIn", vappID)
	api.Token = true

	resp, err := r.PutTrip(ctx, api, vmID)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return response, err
}

func DeleteVmFromVapp(ctx context.Context, r restful.RestAPITripper, vappID string, vmID []string) (types.Task, error) {
	var api types.ICSApi
	var response = types.Task{}

	api.Api = fmt.Sprintf("/vclusters/%s/vms?action=shiftOut", vappID)
	api.Token = true

	resp, err := r.PutTrip(ctx, api, vmID)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return response, err
}

func PowerOnVapp(ctx context.Context, r restful.RestAPITripper, vappID string) (types.Task, error) {
	var api types.ICSApi
	var reqBody *types.Common
	var response = types.Task{}

	api.Api = fmt.Sprintf("/vclusters/%s?action=powerOn", vappID)
	api.Token = true

	resp, err := r.PutTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return response, err
}

func PowerOffVapp(ctx context.Context, r restful.RestAPITripper, vappID string) (types.Task, error) {
	var api types.ICSApi
	var reqBody *types.Common
	var response = types.Task{}

	api.Api = fmt.Sprintf("/vclusters/%s?action=powerOff", vappID)
	api.Token = true

	resp, err := r.PutTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return response, err
}

func PowerOffVappSafely(ctx context.Context, r restful.RestAPITripper, vappID string) (types.Task, error) {
	var api types.ICSApi
	var reqBody *types.Common
	var response = types.Task{}

	api.Api = fmt.Sprintf("/vclusters/%s?action=safelypowerOff", vappID)
	api.Token = true

	resp, err := r.PutTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return response, err
}

func RestartVapp(ctx context.Context, r restful.RestAPITripper, vappID string) (types.Task, error) {
	var api types.ICSApi
	var reqBody *types.Common
	var response = types.Task{}

	api.Api = fmt.Sprintf("/vclusters/%s?action=restart", vappID)
	api.Token = true

	resp, err := r.PutTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return response, err
}
