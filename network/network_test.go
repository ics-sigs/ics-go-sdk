package network

import (
	"context"
	"encoding/json"
	icsgo "github.com/ics-sigs/ics-go-sdk"
	//"github.com/ics-sigs/ics-go-sdk/client/types"
	"testing"
)

var (
	icsConnection = &icsgo.ICSConnection{
		Username: "admin",
		Password: "Cloud@s1",
		Hostname: "10.49.34.161",
		Port:     "443",
		Insecure: true,
	}
)

func TestGetNetworkByName(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	networkName := "manageNetwork0"
	networkClient := NewNetworkService(icsConnection.Client)
	network, err := networkClient.GetNetworkByName(ctx, networkName)
	if err != nil {
		t.Errorf("Failed to get network by name %s. Error: %v\n", networkName, err)
	} else {
		networkJson, _ := json.MarshalIndent(network, "", "\t")
		t.Logf("Network Info: %s\n", string(networkJson))
	}
}

func TestGetNetworkList(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	networkClient := NewNetworkService(icsConnection.Client)
	networkList, err := networkClient.GetNetworkList(ctx)
	if err != nil {
		t.Errorf("Failed to get network list. Error: %v\n", err)
	} else {
		for _, network := range networkList {
			networkJson, _ := json.MarshalIndent(network, "", "\t")
			t.Logf("Network Info: %s\n", string(networkJson))
		}
	}
}

func TestGetNetworByID(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	//networkID := "8ab1a2968145ef350181460bb71200ac" // 10.49.34.22
	//networkID := "8ab0b34979c154880179c2070cbf0044" // 10.49.34.23
	//networkID := "8ab1a21f8e11a48e018e11a9d630002e" // 10.49.34.159
	//networkID := "8ab1a2218d55e067018d55e716d60039" // 10.49.34.161
	networkID := "8ab1a2228e071ead018e072d9f970037" // 10.49.34.162
	networkClient := NewNetworkService(icsConnection.Client)
	network, err := networkClient.GetNetworkByID(ctx, networkID)
	if err != nil {
		t.Errorf("Failed to get network by ID %s. Error: %v\n", networkID, err)
	} else {
		networkJson, _ := json.MarshalIndent(network, "", "\t")
		t.Logf("Network Info: %s\n", string(networkJson))
	}
}

func TestGetSdnNetworkList(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	networkClient := NewNetworkService(icsConnection.Client)
	sdnNetworkList, err := networkClient.GetSdnNetworkList(ctx)
	if err != nil {
		t.Errorf("Failed to get SDN network list. Error: %v\n", err)
	} else {
		for _, network := range sdnNetworkList {
			networkJson, _ := json.MarshalIndent(network, "", "\t")
			t.Logf("SDN Network Info: %s\n", string(networkJson))
		}
	}
}

func TestGetSdnNetworkByName(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	sdnNetworkName := "k3s"
	networkClient := NewNetworkService(icsConnection.Client)
	sdnNetwork, err := networkClient.GetSdnNetworkByName(ctx, sdnNetworkName)
	if err != nil {
		t.Errorf("Failed to get SDN network by name %s. Error: %v\n", sdnNetworkName, err)
	} else {
		sdnNetworkJson, _ := json.MarshalIndent(sdnNetwork, "", "\t")
		t.Logf("SDN Network Info: %s\n", string(sdnNetworkJson))
	}
}

func TestGetSdnNetworByID(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	sdnNetworkID := "d3398ee0-31ed-4f9e-b398-cf2f5b992e73" // 10.49.34.161
	networkClient := NewNetworkService(icsConnection.Client)
	sdnNetwork, err := networkClient.GetSdnNetworkByID(ctx, sdnNetworkID)
	if err != nil {
		t.Errorf("Failed to get SDN network by ID %s. Error: %v\n", sdnNetworkID, err)
	} else {
		sdnNetworkJson, _ := json.MarshalIndent(sdnNetwork, "", "\t")
		t.Logf("SDN Network Info: %s\n", string(sdnNetworkJson))
	}
}
