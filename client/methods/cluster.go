package methods

import (
	"context"
	"encoding/json"
	"github.com/ics-sigs/ics-go-sdk/client/restful"
	"github.com/ics-sigs/ics-go-sdk/client/types"
)

func GetClusterList(ctx context.Context, r restful.RestAPITripper) ([]types.Cluster, error) {
	var api          types.ICSApi
	var response = types.ClusterListRsp{}
	var reqBody      *types.Common
	api.Api = "/clusters"
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