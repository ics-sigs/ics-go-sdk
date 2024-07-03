package cluster

import (
	"context"
	"encoding/json"
	//"fmt"
	icsgo "github.com/ics-sigs/ics-go-sdk"
	"testing"
)

var (
	icsConnection = icsgo.ICSConnection{
		Username: "admin",
		Password: "Cloud@s1",
		Hostname: "10.49.34.162",
		Port:     "443",
		Insecure: true,
	}
)

func TestGetClusterList(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}
	clusterClient := NewClusterService(icsConnection.Client)
	clusterList, err := clusterClient.GetClusterList(ctx)
	if err == nil {
		for i := range clusterList {
			clusterJson, _ := json.MarshalIndent(clusterList[i], "", "\t")
			t.Logf("ClusterInfo: %v\n", string(clusterJson))
		}
	} else {
		t.Errorf("Failed to get cluster list. Error: %v\n", err.Error())
	}
}

func TestGetClusterByName(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	clusterName := "Cluster"
	clusterClient := NewClusterService(icsConnection.Client)
	clusterInfo, err := clusterClient.GetClusterByName(ctx, clusterName)
	if clusterInfo != nil {
		clusterJson, _ := json.MarshalIndent(clusterInfo, "", "\t")
		t.Logf("ClusterInfo: %v\n", string(clusterJson))
	} else {
		t.Errorf("Failed to get cluster info by name. Error: %v\n", err.Error())
	}
}
