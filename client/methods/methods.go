package methods

import (
    "context"
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "strconv"
    "github.com/ics-sigs/ics-go-sdk/client/restful"
    "github.com/ics-sigs/ics-go-sdk/client/types"
)

var TypeFunc types.Func

func init() {
    TypeFunc = types.TypeFunc()
}

func JsonError(err error) error {
    var code, msg    string
    if err != nil {
        code = "400"
        msg = "SDK unmarshal data error"
        err = &types.SDKError{
            Code: code,
            Message: msg,
        }
    }
    return err
}

func HandleResponse(resp *restful.Response, errHttp error) ([]byte, error) {
    var err    error
    var code, msg    string
    var response    []byte
    if errHttp != nil {
        code = "404"
        msg = errHttp.Error()
        err = &types.SDKError{
            Code: code,
            Message: msg,
        }
        return response, err
    }
    if resp.StatusCode() == 200 {
        response = resp.Body()
    } else if resp.StatusCode() == 202 || resp.StatusCode() == 401 || resp.StatusCode() == 403 {
        var e interface{}
        err = json.Unmarshal([]byte(resp.Body()), &e)
        if err != nil {
            code = "501"
            msg = fmt.Sprintf("Service response error: %+v", err)
        } else {
            code = strconv.Itoa(resp.StatusCode())
            msg = fmt.Sprintf("Service response error: %+v", e)
        }
        err = &types.SDKError{
            Code: code,
            Message: msg,
        }
    } else {
        code = "500"
        msg = fmt.Sprintf("Service unknown error. StatusCode:%d", resp.StatusCode())
        err = &types.SDKError{
            Code: code,
            Message: msg,
        }
    }
    return response, err
}

func ValidUserSession(ctx context.Context, r restful.RestAPITripper, user *types.UserSession) error {
    var reqBody    *interface{}
    var api        types.ICSApi
    var err    error
    var code, msg    string

    userId := "anonymous"
    if len(user.UserId) > 0 {
        userId = user.UserId
    }
    api.Api = fmt.Sprintf("/users/%s/themes",userId)
    api.Token = true

    resp, errs := r.GetTrip(ctx, api, reqBody)
    if err != nil {
        code = "404"
        msg = errs.Error()
        err = &types.SDKError{
            Code: code,
            Message: msg,
        }
    } else {
        if resp.StatusCode() == 200 {
            err = nil
        } else if resp.StatusCode() == 401 {
            code = "401"
            msg = "the session is not authenticated"
            err = &types.SDKError{
                Code: code,
                Message: msg,
            }
        } else {
            code = "400"
            msg = "Client Request Error"
            err = &types.SDKError{
                Code: code,
                Message: msg,
            }
        }
    }
    return err
}

func Login(ctx context.Context, r restful.RestAPITripper, req *types.Login) (*types.LoginResponse, error) {
    var reqBody    *types.Login
    var api        types.ICSApi

    var response = types.LoginResponse{}

    reqBody = req
    api.Api = "/system/user/sdklogin"
    api.Token = false

    resp, err := r.PostTrip(ctx, api, reqBody)
    respBody, err1 := HandleResponse(resp, err)
    if err1 != nil {
        err = err1
    } else if respBody != nil {
        jsonErr := json.Unmarshal([]byte(respBody), &response)
        err = JsonError(jsonErr)
    }

    return &response, err
}

func Logout(ctx context.Context, r restful.RestAPITripper) (*types.Task, error) {
    var reqBody      *types.Common
    var api        types.ICSApi

    var response = types.Task{}

    api.Api = "/logout"
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

func GetTaskInfo(ctx context.Context, r restful.RestAPITripper, req *types.Task) (*types.TaskInfo, error) {
    var reqBody      *types.Common
    var api          types.ICSApi
    var response = types.TaskInfo{}

    if req == nil || req.TaskId == "" {
        return nil, nil
    }
    api.Api = fmt.Sprintf("/tasks/%s", req.TaskId)
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

func GetPublicKey(ctx context.Context, r restful.RestAPITripper) (string, error) {
    var reqBody      *types.Common
    var api          types.ICSApi
    var publicKey    string

    api.Api = "/system/publickey"
    api.Token = true

    resp, err := r.GetTrip(ctx, api, reqBody)
    respBody, err := HandleResponse(resp, err)
    if err != nil {
        publicKey = ""
    } else {
        publicKey = string(respBody)
    }

    return publicKey, err
}

func GetLoginPolicy(ctx context.Context, r restful.RestAPITripper) (*types.LoginPolicy, error) {
    var reqBody      *types.Common
    var api          types.ICSApi
    var response = types.LoginPolicy{}

    api.Api = "/system/loginpolicy"
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

func GenerateCheckParams(ctx context.Context, r restful.RestAPITripper, params string) (string, error) {
    publicKey, err := GetPublicKey(ctx, r)
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
