package methods

import (
    "context"
    "time"

    "github.com/ics-sigs/ics-go-sdk/client/restful"
    "github.com/ics-sigs/ics-go-sdk/client/types"
)

var serviceInstance = types.ManagedObjectReference{
    Type:  "ServiceInstance",
    Value: "ServiceInstance",
}

func GetServiceContent(ctx context.Context, r restful.RestAPITripper) (types.ServiceContent, error) {
    //req := types.RetrieveServiceContent{
    //    This: serviceInstance,
    //}
    //
    //res, err := RetrieveServiceContent(ctx, r, &req)
    //if err != nil {
    //    return types.ServiceContent{}, err
    //}
    serviceContent := types.ServiceContent {}

    return serviceContent, nil
}

func GetCurrentTime(ctx context.Context, r restful.RestAPITripper) (*time.Time, error) {
    //req := types.CurrentTime{
    //    This: serviceInstance,
    //}
    //
    //res, err := CurrentTime(ctx, r, &req)
    //if err != nil {
    //    return nil, err
    //}
    var currentTime    time.Time

    return &currentTime, nil
}
