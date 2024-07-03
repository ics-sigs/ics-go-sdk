package methods

import (
	"context"
	"fmt"
	"github.com/ics-sigs/ics-go-sdk/client/restful"
	"github.com/ics-sigs/ics-go-sdk/client/types"
	"testing"
)

func TestLogin(t *testing.T) {
	url, _ := restful.ParseURL("https://192.168.220.164")
	client := restful.NewClient(url, true)
	ctx := context.Background()

	req := types.Login{
		Domain: "internal",
		Locale: "cn",
		Password: "admin@inspur",
		Username: "admin",
	}
	resp, err := Login(ctx, client, &req)
	fmt.Println("response:", resp)
	if err == nil {
		println(err)
	}

	println(resp)
}
