package methods

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/ics-sigs/ics-go-sdk/client/restful"
    "github.com/ics-sigs/ics-go-sdk/client/types"
)

func GetTagById(ctx context.Context, r restful.RestAPITripper, id string) (*types.Tag, error) {
    var reqBody      *types.Common
    var api          types.ICSApi
    var response = types.Tag{}

    if len(id) <= 0 {
        id = "anonymous"
    }
    api.Api = fmt.Sprintf("/tags/%s",id)
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

func ListAttachedTags(ctx context.Context, r restful.RestAPITripper, targetType string, ref string) ([]types.TreeItem, error) {
    var reqBody      *types.Common
    var api          types.ICSApi
    var response     []types.TreeItem

    api.Api = fmt.Sprintf("/tags/bindings?format=tree&tagSourceType=%s&sourceIds=%s", targetType, ref)
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