package restful

import (
    "encoding/json"
    "fmt"

    "context"
    "crypto/tls"
    "net/url"
    "regexp"
    "sync"
    "time"

    "github.com/ics-sigs/ics-go-sdk/client/types"
    "github.com/go-resty/resty/v2"
    "k8s.io/klog"
)

const (
    SessionCookieName = "JSESSOINID"
)

type Error struct {
    /* variables */
}

type RestAPITripper interface {
    GetTrip(ctx context.Context, api types.ICSApi, req interface{}) (*Response, error)
    PostTrip(ctx context.Context, api types.ICSApi, req interface{}) (*Response, error)
    PutTrip(ctx context.Context, api types.ICSApi, req interface{}) (*Response, error)
    DeleteTrip(ctx context.Context, api types.ICSApi, req interface{}) (*Response, error)
}

type Client struct {
    HttpClient *resty.Client

    u *url.URL
    k bool // Named after curl's -k flag

    hostsMu sync.Mutex
    hosts   map[string]string

    Namespace string // ics internal
    Version   string //  ics api version

    Authorization    string
    AccessKeyID      string
    AccessKeySecret  string
    CheckParams      string
}

var schemeMatch = regexp.MustCompile(`^\w+://`)

// ParseURL is wrapper around url.Parse, where Scheme defaults to "https" and Path defaults to "/"
func ParseURL(s string) (*url.URL, error) {
    var err error
    var u *url.URL

    if s != "" {
        // Default the scheme to https
        if !schemeMatch.MatchString(s) {
            s = "https://" + s
        }

        u, err = url.Parse(s)
        if err != nil {
            return nil, err
        }

        // Default the path to /sdk
        if u.Path == "" {
            u.Path = "/"
        }

        if u.User == nil {
            u.User = url.UserPassword("", "")
        }
    }

    return u, nil
}

func ParseURI(url *url.URL, uri string) string {
    apiPath := url.String()
    apiPath += uri
    return apiPath
}

func NewClient(u *url.URL, insecure bool) *Client {
    c := Client{
        u: u,
        k: insecure,
    }

    c.HttpClient = resty.New()

    c.HttpClient.SetHeader("Content-Type", "application/json; charset=utf-8").
        SetHeader("Accept", "application/json; charset=utf-8").
        SetHeader("version", "5.8")

    c.HttpClient.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: insecure })

    c.HttpClient.SetTimeout(1 * time.Minute)

    c.hosts = make(map[string]string)

    // Remove user information from a copy of the URL
    c.u = c.URL()
    c.u.User = nil

    return &c
}

func (c *Client) URL() *url.URL {
    urlCopy := *c.u
    return &urlCopy
}

func (c *Client) GetToken() string {
    return c.Authorization
}

func (c *Client) SetToken(token string) {
    c.Authorization = token
}

func (c *Client) GetAccessKey() string {
    return fmt.Sprintf("ICS %s:%s", c.AccessKeyID, c.AccessKeySecret)
}

func (c *Client) SetAccessKey(keyID string, keySecret string) {
    c.AccessKeyID = keyID
    c.AccessKeySecret = keySecret
}

func (c *Client) SetCheckParams(checkParams string) {
    c.CheckParams = checkParams
}

func (c *Client) generateRequest(api types.ICSApi) *resty.Request {
    client := c.HttpClient
    request := client.R()
    if api.Token {
        if len(c.GetToken()) == 0 {
            request.SetHeader("Authorization", c.GetAccessKey())
        } else {
            request.SetHeader("Authorization", c.GetToken())
        }
    }
    if c.CheckParams != "" {
        request.SetHeader("CheckParams", c.CheckParams)
    }
    return request
}

func (c *Client) GetTrip(ctx context.Context, api types.ICSApi, req interface{}) (*Response, error) {
    var errorParam    Response

    apiPath := ParseURI(c.URL(), api.Api)

    reqBody, err := json.Marshal(req)
    if err != nil {
        return &errorParam, err
    }

    var getReq map[string]interface{}
    err = json.Unmarshal([]byte(reqBody), &getReq)
    if err != nil {
        return &errorParam, err
    }
    klog.V(10).Infof("Get Request API path:%s, parameters:%+v", apiPath, getReq)

    resp, err := c.generateRequest(api).
        SetQueryParams(formatReqParams(getReq)).
        Get(apiPath)

    response := Response{
        resp.RawResponse,
        resp.Body(),
    }
    klog.V(10).Infof("Get Request API path:%s, Response:%+v", apiPath, response)
    return &response, err
}

func formatReqParams(reqbody map[string]interface{}) map[string]string {
    formatParams := make(map[string]string)
    for key, value := range reqbody {
        strValue := fmt.Sprintf("%v", value)
        formatParams[key] = strValue
    }
    return formatParams
}

func (c *Client) PostTrip(ctx context.Context, api types.ICSApi, req interface{}) (*Response, error) {
    var errorParam    Response

    apiPath := ParseURI(c.URL(), api.Api)

    reqBody, err := json.Marshal(req)
    if err != nil {
        return &errorParam, err
    }
    klog.V(10).Infof("Post Request API path:%s, body:%+v", apiPath, reqBody)

    // POST JSON string
    resp, err := c.generateRequest(api).
        SetBody(reqBody).
        Post(apiPath)

    response := Response{
        resp.RawResponse,
        resp.Body(),
    }
    klog.V(10).Infof("Post Request API path:%s, Response:%+v", apiPath, response)
    return &response, err
}

func (c *Client) PutTrip(ctx context.Context, api types.ICSApi, req interface{}) (*Response, error) {
    var errorParam    Response

    apiPath := ParseURI(c.URL(), api.Api)

    reqBody, err := json.Marshal(req)
    if err != nil {
        return &errorParam, err
    }

    klog.V(10).Infof("Put Request API path:%s, body:%+v", apiPath, reqBody)

    // Just one sample of PUT, refer POST for more combination
    resp, err := c.generateRequest(api).
        SetBody(reqBody).
        SetError(&Error{}).
        Put(apiPath)

    response := Response{
        resp.RawResponse,
        resp.Body(),
    }
    klog.V(10).Infof("Put Request API path:%s, Response:%+v", apiPath, response)
    return &response, err
}

func (c *Client) DeleteTrip(ctx context.Context, api types.ICSApi, req interface{}) (*Response, error) {
    var errorParam    Response

    apiPath := ParseURI(c.URL(), api.Api)

    reqBody, err := json.Marshal(req)
    if err != nil {
        return &errorParam, err
    }

    klog.V(10).Infof("Delete Request API path:%s, body:%+v", apiPath, reqBody)

    // DELETE a articles with payload/body as a JSON string
    resp, err := c.generateRequest(api).
        SetError(&Error{}).
        SetBody(reqBody).
        Delete(apiPath)

    response := Response{
        resp.RawResponse,
        resp.Body(),
    }
    klog.V(10).Infof("Delete Request API path:%s, Response:%+v", apiPath, response)
    return &response, err
}
