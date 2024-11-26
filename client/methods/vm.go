package methods

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ics-sigs/ics-go-sdk/client/restful"
	"github.com/ics-sigs/ics-go-sdk/client/types"
)

func GetVMById(ctx context.Context, r restful.RestAPITripper, vmId string) (*types.VirtualMachine, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.VirtualMachine{}

	if len(vmId) <= 0 {
		vmId = "anonymous"
	}
	api.Api = fmt.Sprintf("/vms/%s", vmId)
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

func GetVMByIP(ctx context.Context, r restful.RestAPITripper, ipAddy string) (*types.VirtualMachine, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var vmPages = types.VMPageResponse{}
	var response = types.VirtualMachine{}

	if len(ipAddy) <= 6 {
		ipAddy = "255.255.255.255"
	}
	api.Api = fmt.Sprintf("/vms?pageSize=1&currentPage=1&sortField=&sort=desc&vnicIp=%s", ipAddy)
	api.Token = true

	resp, err := r.GetTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &vmPages)
		err = JsonError(jsonErr)
	}
	if vmPages.TotalSize == 1 {
		response = vmPages.Items[0]
	}

	return &response, err
}

func GetVMByName(ctx context.Context, r restful.RestAPITripper, name string) (*types.VirtualMachine, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var vmPages = types.VMPageResponse{}
	var response = types.VirtualMachine{}

	if len(name) <= 0 {
		name = "vm@anonymous"
	}
	api.Api = fmt.Sprintf("/vms?pageSize=1&currentPage=1&sortField=&sort=desc&name=%s", name)
	api.Token = true

	resp, err := r.GetTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &vmPages)
		err = JsonError(jsonErr)
	}
	if vmPages.TotalSize == 1 {
		response = vmPages.Items[0]
	}

	return &response, err
}

func GetVMPageList(ctx context.Context, r restful.RestAPITripper, req *types.VMPageReq) (*types.VMPageResponse, error) {
	var api types.ICSApi
	var response = types.VMPageResponse{}

	api.Api = "/vms"
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

func PowerOnVMById(ctx context.Context, r restful.RestAPITripper, vmId string) (*types.Task, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.Task{}

	if len(vmId) <= 0 {
		vmId = "anonymous"
	}
	api.Api = fmt.Sprintf("/vms/%s?action=poweron", vmId)
	api.Token = true

	resp, err := r.PutTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return &response, err
}

func PowerOffVMById(ctx context.Context, r restful.RestAPITripper, vmId string) (*types.Task, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.Task{}

	if len(vmId) <= 0 {
		vmId = "anonymous"
	}
	api.Api = fmt.Sprintf("/vms/%s?action=poweroff", vmId)
	api.Token = true

	resp, err := r.PutTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return &response, err
}

func ShutdownVMById(ctx context.Context, r restful.RestAPITripper, vmId string) (*types.Task, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.Task{}

	if len(vmId) <= 0 {
		vmId = "anonymous"
	}
	api.Api = fmt.Sprintf("/vms/%s?action=shutdown", vmId)
	api.Token = true

	resp, err := r.PutTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return &response, err
}

func RestartVMById(ctx context.Context, r restful.RestAPITripper, vmId string) (*types.Task, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.Task{}

	if len(vmId) <= 0 {
		vmId = "anonymous"
	}
	api.Api = fmt.Sprintf("/vms/%s?action=restart", vmId)
	api.Token = true

	resp, err := r.PutTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return &response, err
}

func DeleteVMById(ctx context.Context, r restful.RestAPITripper, vmId string, deleteFile bool, removeData bool) (*types.Task, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.Task{}

	if len(vmId) <= 0 {
		vmId = "anonymous"
	}
	api.Api = fmt.Sprintf("/vms/%s?deleteFile=%t&removeData=%t", vmId, deleteFile, removeData)
	api.Token = true

	resp, err := r.DeleteTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return &response, err
}

func GetVMList(ctx context.Context, r restful.RestAPITripper) (*types.VMPageResponse, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.VMPageResponse{}

	api.Api = "/vms"
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

func SetVM(ctx context.Context, r restful.RestAPITripper, vmInfo types.VirtualMachine) (*types.Task, error) {
	var api types.ICSApi
	var response = types.Task{}

	api.Api = fmt.Sprintf("/vms/%s", vmInfo.ID)
	api.Token = true

	resp, err := r.PutTrip(ctx, api, vmInfo)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return &response, err
}

func CreateVMByTemplate(ctx context.Context, r restful.RestAPITripper, vmSpec types.VirtualMachine, quickClone bool) (*types.Task, error) {
	var api types.ICSApi
	var response = types.Task{}

	api.Api = fmt.Sprintf("/vms?action=createByTemplate&quickClone=%t", quickClone)
	api.Token = true

	resp, err := r.PostTrip(ctx, api, vmSpec)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return &response, err
}

func ImportVM(ctx context.Context, r restful.RestAPITripper, vmSpec types.VirtualMachine,
	ovaFilePath string, hostUUID string, rateLimit int) (*types.Task, error) {
	var api types.ICSApi
	var response = types.Task{}

	if rateLimit < 20 || rateLimit > 100 {
		rateLimit = 40
	}
	api.Api = fmt.Sprintf("/vms/ovfs?action=import&ovaFile=%s&hostId=%s&rateLimit=%d", ovaFilePath, hostUUID, rateLimit)
	api.Token = true

	resp, err := r.PutTrip(ctx, api, vmSpec)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return &response, err
}

func GetOvaConfig(ctx context.Context, r restful.RestAPITripper, ovaFilePath string, hostUUID string,
	imageHostUUID string) (*types.VirtualMachine, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.VirtualMachine{}

	api.Api = fmt.Sprintf("/vms/ovfs?action=preimport&hostUuid=%s&filePath=%s&imageHostId=%s", hostUUID, ovaFilePath, imageHostUUID)
	api.Token = true

	resp, err := r.PutTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return &response, err
}

func GetVMTemplateList(ctx context.Context, r restful.RestAPITripper) (*types.VMPageResponse, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.VMPageResponse{}

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

func GetVMTemplateById(ctx context.Context, r restful.RestAPITripper, vmtId string) (*types.VirtualMachine, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.VirtualMachine{}

	if len(vmtId) <= 0 {
		vmtId = "anonymous"
	}
	api.Api = fmt.Sprintf("/vmtemplates/%s", vmtId)
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
