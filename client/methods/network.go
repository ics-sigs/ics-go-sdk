package methods

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/inspur-ics/ics-go-sdk/client/restful"
	"github.com/inspur-ics/ics-go-sdk/client/types"
)

func GetNetworkList(ctx context.Context, r restful.RestAPITripper) ([]types.Network, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response types.NetworkPageResponse

	api.Api = "/networks"
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

func GetNetworkByID(ctx context.Context, r restful.RestAPITripper, networkID string) (types.Network, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response types.Network

	api.Api = fmt.Sprintf("/networks/%s", networkID)
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
