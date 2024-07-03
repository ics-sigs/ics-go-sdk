package host

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
		Hostname: "10.49.34.161",
		Port:     "443",
		Insecure: true,
	}
)

func TestGetHostList(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	hostClient := NewHostService(icsConnection.Client)
	hostList, err := hostClient.GetHostList(ctx)
	if err != nil {
		t.Errorf("Failed to get host list. Error: %v\n", err.Error())
	} else {
		for _, hostInfo := range hostList {
			hostJson, _ := json.MarshalIndent(hostInfo, "", "\t")
			t.Logf("Host Info: %s\n", string(hostJson))
		}
	}
}

func TestGetHostInfo(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	hostUUID := "2cf4bb24-cc1d-46fb-862e-e408da0898ce"
	hostClient := NewHostService(icsConnection.Client)
	hostInfo, err := hostClient.GetHost(ctx, hostUUID)
	if err != nil {
		t.Errorf("Failed to get host info. Error: %v\n", err.Error())
	} else {
		hostJson, _ := json.MarshalIndent(hostInfo, "", "\t")
		t.Logf("Host Info: %s\n", string(hostJson))
	}
}

func TestGetHostListByClusterID(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	clusterID := "8a878bda6f7012c7016f87e979120798"
	hostClient := NewHostService(icsConnection.Client)
	hostList, err := hostClient.GetHostListByClusterID(ctx, clusterID)
	if err != nil {
		t.Errorf("Failed to get host list by clusterID. Error: %v\n", err.Error())
	} else {
		for _, hostInfo := range hostList {
			hostJson, _ := json.MarshalIndent(hostInfo, "", "\t")
			t.Logf("Host Info: %s\n", string(hostJson))
		}
	}
}

func TestGetHostListByClusterName(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	clusterName := "cluster"
	hostClient := NewHostService(icsConnection.Client)
	hostList, err := hostClient.GetHostListByClusterName(ctx, clusterName)
	if err != nil {
		t.Errorf("Failed to get host list by clusterName. Error: %v\n", err.Error())
	} else {
		for _, hostInfo := range hostList {
			hostJson, _ := json.MarshalIndent(hostInfo, "", "\t")
			t.Logf("Host Info: %s\n", string(hostJson))
		}
	}
}

func TestGetHostListByStorageID(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	storageID := "8a878bda6f6f3ca4016f6f458b50003c"
	hostClient := NewHostService(icsConnection.Client)
	hostList, err := hostClient.GetHostListByStorageID(ctx, storageID)
	if err != nil {
		t.Errorf("Failed to get host list by storageID. Error: %v\n", err.Error())
	} else {
		for _, hostInfo := range hostList {
			hostJson, _ := json.MarshalIndent(hostInfo, "", "\t")
			t.Logf("Host Info: %s\n", string(hostJson))
		}
	}
}

func TestGetAvailHostListByStorageID(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	storageID := "8a878bda6f6f3ca4016f6f458b50003c"
	hostClient := NewHostService(icsConnection.Client)
	hostList, err := hostClient.GetAvailHostListByStorageID(ctx, storageID)
	if err != nil {
		t.Errorf("Failed to get avail host list by storageID. Error: %v\n", err.Error())
	} else {
		for _, hostInfo := range hostList {
			hostJson, _ := json.MarshalIndent(hostInfo, "", "\t")
			t.Logf("Host Info: %s\n", string(hostJson))
		}
	}
}

func TestGetHostListBySwitchID(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	switchID := "3f27df9d2ebb11eaa2691a130ac12531"
	hostClient := NewHostService(icsConnection.Client)
	hostList, err := hostClient.GetHostListBySwitchID(ctx, switchID)
	if err != nil {
		t.Errorf("Failed to get host list by switchID. Error: %v\n", err.Error())
	} else {
		for _, hostInfo := range hostList {
			hostJson, _ := json.MarshalIndent(hostInfo, "", "\t")
			t.Logf("Host Info: %s\n", string(hostJson))
		}
	}
}

func TestGetHostListByNetworkID(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	networkID := "8a878bda6f6f3ca4016f6f4409bf002e"
	hostClient := NewHostService(icsConnection.Client)
	hostList, err := hostClient.GetHostListByNetworkID(ctx, networkID)
	if err != nil {
		t.Errorf("Failed to get host list by networkID. Error: %v\n", err.Error())
	} else {
		for _, hostInfo := range hostList {
			hostJson, _ := json.MarshalIndent(hostInfo, "", "\t")
			t.Logf("Host Info: %s\n", string(hostJson))
		}
	}
}

func TestGetHostListByNetworkName(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	networkName := "manageNetwork0"
	hostClient := NewHostService(icsConnection.Client)
	hostList, err := hostClient.GetHostListByNetworkName(ctx, networkName)
	if err != nil {
		t.Errorf("Failed to get host list by networkName. Error: %v\n", err.Error())
	} else {
		for _, hostInfo := range hostList {
			hostJson, _ := json.MarshalIndent(hostInfo, "", "\t")
			t.Logf("Host Info: %s\n", string(hostJson))
		}
	}
}

func TestGetHostAccessibleDatastoreList(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	hostid := "792b1e3f-8be6-43de-b8c7-27a0327dcd97"
	hostClient := NewHostService(icsConnection.Client)
	storagelist, err := hostClient.GetHostAccessibleDatastoreList(ctx, hostid)
	if storagelist != nil {
		for i := range storagelist {
			t.Logf("StorageName: %s\n", storagelist[i].Name)
		}
	} else {
		t.Logf("Failed to get datastore list. Error: %v\n", err.Error())
	}
}

func TestGetHostListByDC(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	hostClient := NewHostService(icsConnection.Client)
	hostlist, err := hostClient.GetHostListByDC(ctx, "3f0094542ebb11eaa2691a130ac12531")
	if hostlist != nil && len(hostlist) >= 1 {
		for _, hostInfo := range hostlist {
			hostJson, _ := json.MarshalIndent(hostInfo, "", "\t")
			t.Logf("Host Info: %s\n", string(hostJson))
		}
	} else if err != nil {
		t.Errorf("Failed to get host list by dc. Error: %v\n", err.Error())
	}
}

func TestGetHostHealthInfo(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	hostUUID := "2cf4bb24-cc1d-46fb-862e-e408da0898ce"
	hostClient := NewHostService(icsConnection.Client)
	healthInfo, err := hostClient.GetHostHealthInfo(ctx, hostUUID)

	if err != nil {
		t.Errorf("Failed to get host health info. Error: %v\n", err.Error())
	} else {
		healthJson, _ := json.MarshalIndent(healthInfo, "", "\t")
		t.Logf("HostHealth Info: %s\n", string(healthJson))
	}
}
