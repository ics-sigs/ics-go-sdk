package tag

import (
	"context"
	"encoding/json"
	"fmt"
	icsgo "github.com/ics-sigs/ics-go-sdk"
	"testing"
)

var (
	icsConnection = icsgo.ICSConnection{
		Username: "admin",
		Password: "Cloud@s1",
		Hostname: "10.49.34.161",
		Port:     "443",
		Insecure: true,
	}
)

func TestGetTag(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	tagClient := NewTagsService(icsConnection.Client)
	tag, err := tagClient.GetTag(ctx, "ba6d1a64-0515-45db-8b80-be5d3f932256")
	if tag != nil {
		tagJson, _ := json.MarshalIndent(tag, "", "\t")
		t.Logf("Tag Info: %s\n", string(tagJson))
	}
}

func TestListAttachedTags(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	tagClient := NewTagsService(icsConnection.Client)
	tags, err := tagClient.ListAttachedTags(ctx, "HOST", "56c940ba-f62d-423f-aaf3-f91dfc22d7ea")
	if err != nil {
		t.Fatalf("Get attached tags failed! Error:%v", err)
	} else {
		fmt.Printf("tags:%+v", tags)
	}
}
