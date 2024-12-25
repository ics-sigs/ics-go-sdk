package icsgo

import (
    "context"
    "net/url"

    "github.com/ics-sigs/ics-go-sdk/client"
    "github.com/ics-sigs/ics-go-sdk/client/restful"
    "github.com/ics-sigs/ics-go-sdk/session"
)

type Client struct {
    *client.Client

    SessionManager *session.Manager
}

// NewClient creates a new client from a URL. The client authenticates with the
// server with username/password before returning if the URL contains user information.
func NewClient(ctx context.Context, u *url.URL, insecure bool, locale string) (*Client, error) {
    restfulClient := restful.NewClient(u, insecure, locale)
    sdkClient, err := client.NewClient(ctx, restfulClient)
    if err != nil {
        return nil, err
    }

    c := &Client{
        Client:         sdkClient,
        SessionManager: session.NewManager(sdkClient),
    }

    // Only login if the URL contains user information.
    if u.User != nil {
        err = c.Login(ctx, u.User)
        if err != nil {
            return nil, err
        }
    }

    return c, nil
}

// Login dispatches to the SessionManager.
func (c *Client) Login(ctx context.Context, u *url.Userinfo) error {
    return c.SessionManager.Login(ctx, u)
}

// Logout dispatches to the SessionManager.
func (c *Client) Logout(ctx context.Context) error {
   return c.SessionManager.Logout(ctx)
}

//// PropertyCollector returns the session's default property collector.
//func (c *Client) PropertyCollector() *property.Collector {
//    return property.DefaultCollector(c.Client)
//}
//
//// RetrieveOne dispatches to the Retrieve function on the default property collector.
//func (c *Client) RetrieveOne(ctx context.Context, obj types.ManagedObjectReference, p []string, dst interface{}) error {
//    return c.PropertyCollector().RetrieveOne(ctx, obj, p, dst)
//}
//
//// Retrieve dispatches to the Retrieve function on the default property collector.
//func (c *Client) Retrieve(ctx context.Context, objs []types.ManagedObjectReference, p []string, dst interface{}) error {
//    return c.PropertyCollector().Retrieve(ctx, objs, p, dst)
//}
//
//// Wait dispatches to property.Wait.
//func (c *Client) Wait(ctx context.Context, obj types.ManagedObjectReference, ps []string, f func([]types.PropertyChange) bool) error {
//    return property.Wait(ctx, c.PropertyCollector(), obj, ps, f)
//}
//
//// IsVC returns true if we are connected to a vCenter
//func (c *Client) IsVC() bool {
//    return c.Client.IsVC()
//}
