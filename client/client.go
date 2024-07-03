package client

import (
    "context"
    "encoding/json"
    "github.com/ics-sigs/ics-go-sdk/client/methods"
    "github.com/ics-sigs/ics-go-sdk/client/restful"
)

const (
    Namespace = "client"
    Version   = "5.8"
    Path      = "/"
)

type Client struct {
    *restful.Client

    RestAPITripper restful.RestAPITripper
}

type marshaledClient struct {
    RestAPIClient     *restful.Client
}
// NewClient creates and returns a new client with the ServiceContent field
// filled in.
func NewClient(ctx context.Context, rt restful.RestAPITripper) (*Client, error) {
    c := Client{
        RestAPITripper: rt,
    }

    // Set client if it happens to be a soap.Client
    if sc, ok := rt.(*restful.Client); ok {
        c.Client = sc
    }

    var err error
    _, err = methods.GetServiceContent(ctx, rt)
    if err != nil {
        return nil, err
    }

    return &c, nil
}

func (c *Client) MarshalJSON() ([]byte, error) {
    m := marshaledClient{
        RestAPIClient:     c.Client,
    }

    return json.Marshal(m)
}

func (c *Client) UnmarshalJSON(b []byte) error {
    var m marshaledClient

    err := json.Unmarshal(b, &m)
    if err != nil {
        return err
    }

    *c = Client{
        Client:         m.RestAPIClient,
        RestAPITripper:   m.RestAPIClient,
    }

    return nil
}

// Valid returns whether or not the client is valid and ready for use.
// This should be called after unmarshalling the client.
func (c *Client) Valid() bool {
    if c == nil {
        return false
    }

    if c.Client == nil {
        return false
    }

    return true
}
