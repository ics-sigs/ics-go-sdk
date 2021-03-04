package methods

import (
	"context"
	"encoding/json"
	"github.com/inspur-ics/ics-go-sdk/client/restful"
	"github.com/inspur-ics/ics-go-sdk/client/types"
)

func GetVMTemplateList(ctx context.Context, r restful.RestAPITripper) (*types.VMTemplatePageResponse, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.VMTemplatePageResponse{}

	api.Api = "/vmtemplates"
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
