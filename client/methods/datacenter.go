package methods

import (
	"context"
	"encoding/json"
	"fmt"

	//	"fmt"
	"github.com/ics-sigs/ics-go-sdk/client/restful"
	"github.com/ics-sigs/ics-go-sdk/client/types"
)

func GetAllDatacenterList(ctx context.Context, r restful.RestAPITripper) (*types.DatacenterPageResponse, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.DatacenterPageResponse{}

	api.Api = "/datacenters"
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

func GetDatacenterById(ctx context.Context, r restful.RestAPITripper, datacenterId string) (*types.Datacenter, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.Datacenter{}

	if len(datacenterId) <= 0 {
		datacenterId = "anonymous"
	}
	api.Api = fmt.Sprintf("/datacenters/%s", datacenterId)
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

func GetDatacenterByName(ctx context.Context, r restful.RestAPITripper, datacenterName string) (*types.Datacenter, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.Datacenter{}

	if len(datacenterName) <= 0 {
		datacenterName = "anonymous"
	}
	api.Api = fmt.Sprintf("/datacenters?datacenterName=%s", datacenterName)
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

func GetDatacenterVMById(ctx context.Context, r restful.RestAPITripper, datacenterId string) (*types.VMPageResponse, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.VMPageResponse{}

	if len(datacenterId) <= 0 {
		datacenterId = "anonymous"
	}
	api.Api = fmt.Sprintf("/datacenters/%s/vms", datacenterId)
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
