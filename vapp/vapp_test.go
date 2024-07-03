package vapp

import (
	"context"
	"encoding/json"
	//"fmt"
	icsgo "github.com/ics-sigs/ics-go-sdk"
	"github.com/ics-sigs/ics-go-sdk/client/types"
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

func TestGetVappList(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}
	vappClient := NewVappService(icsConnection.Client)
	vappList, err := vappClient.GetVappList(ctx)
	if err == nil {
		for i := range vappList {
			vappJson, _ := json.MarshalIndent(vappList[i], "", "\t")
			t.Logf("VappInfo: %v\n", string(vappJson))
		}
	} else {
		t.Errorf("Failed to get vapp list. Error: %v\n", err.Error())
	}
}

func TestGetVappByName(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vappName := "vapp-test"
	vappClient := NewVappService(icsConnection.Client)
	vappInfo, err := vappClient.GetVappByName(ctx, vappName)
	if vappInfo != nil {
		vappJson, _ := json.MarshalIndent(vappInfo, "", "\t")
		t.Logf("VappInfo: %v\n", string(vappJson))
	} else {
		t.Errorf("Failed to get vapp info by name. Error: %v\n", err.Error())
	}
}

func TestCreateVapp(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vappReq := types.VappCreateReq{
		Name:        "test_1",
		Description: "go-sdk-test",
		//DataCenterID: "970c1a57e72711ec8e00aef2d98d5652", //10.49.34.22
		//DataCenterID: "04b07c92c1e211ebb21adec2e1510cb9", //10.49.34.23
		//DataCenterID: "62567b0f4a0a11eda2e896b6da1bc6aa", //10.49.34.159
		//DataCenterID: "892506b8bebc11eeb7309aecff09a95a", //10.49.34.161
		DataCenterID: "302b95a1d9c811eea1c2bada90245722", //10.49.34.162
	}
	vappClient := NewVappService(icsConnection.Client)
	task, err := vappClient.CreateVapp(ctx, vappReq)
	if err != nil {
		t.Errorf("Failed to create vapp. Error: %v\n", err)
	} else {
		taskJson, _ := json.MarshalIndent(task, "", "\t")
		t.Logf("taskInfo: %v\n", string(taskJson))
	}
}

func TestDeleteVapp(t *testing.T) {
	var task types.Task
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vappID := "8ab1a2218dc239da018dc267a5970047"
	vappClient := NewVappService(icsConnection.Client)
	needAuth, err := vappClient.IsDeleteNeedIdentityAuth(ctx)
	if needAuth {
		t.Logf("Delete vapp %s with check params", vappID)
		task, err = vappClient.DeleteVappWithCheckParams(ctx, vappID, icsConnection.Password)
	} else {
		task, err = vappClient.DeleteVapp(ctx, vappID)
	}
	if err != nil {
		t.Errorf("Failed to delete vapp. Error: %v\n", err)
	} else {
		t.Logf("Waiting task %v finish.....\n", task.TaskId)
		taskInfo, err := vappClient.TraceTaskProcess(&task)
		if err != nil {
			t.Fatalf("Failed to trace task. Error: %v", err)
		} else {
			taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
			t.Logf("Task Status: %v\n", string(taskJson))
		}
	}
}

func TestAddVmToVapp(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vappID := "8a878bda781f145e01784eed976501f7"
	vmID := []string{"8a878bda6f7012c7016f70b40ed000a1"}
	//vmID := []string{"8a878bda6f7012c7016f70b40ed000a1", "8a878bda6f6f3ca4016f6f6eb8d300bb"}
	vappClient := NewVappService(icsConnection.Client)
	task, err := vappClient.AddVmToVapp(ctx, vappID, vmID)
	if err != nil {
		t.Errorf("Failed to add vm to vapp. Error: %v\n", err)
	} else {
		taskJson, _ := json.MarshalIndent(task, "", "\t")
		t.Logf("taskInfo: %v\n", string(taskJson))
	}

}

func TestDeleteVmFromVapp(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vappID := "8a878bda781f145e01784eed976501f7"
	vmID := []string{"8a878bda6f7012c7016f70b40ed000a1"}
	//vmID := []string{"8a878bda6f7012c7016f70b40ed000a1", "8a878bda6f6f3ca4016f6f6eb8d300bb"}
	vappClient := NewVappService(icsConnection.Client)
	task, err := vappClient.DeleteVmFromVapp(ctx, vappID, vmID)
	if err != nil {
		t.Errorf("Failed to delete vm from vapp. Error: %v\n", err)
	} else {
		taskJson, _ := json.MarshalIndent(task, "", "\t")
		t.Logf("taskInfo: %v\n", string(taskJson))
	}

}

func TestPowerOnVapp(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vappID := "8a878bda781f145e01784eed976501f7"
	vappClient := NewVappService(icsConnection.Client)
	task, err := vappClient.PowerOnVapp(ctx, vappID)
	if err != nil {
		t.Errorf("Failed to poweron vapp. Error: %v\n", err)
	} else {
		taskJson, _ := json.MarshalIndent(task, "", "\t")
		t.Logf("taskInfo: %v\n", string(taskJson))
	}
}

func TestPowerOffVapp(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vappID := "8a878bda781f145e01784eed976501f7"
	vappClient := NewVappService(icsConnection.Client)
	task, err := vappClient.PowerOffVapp(ctx, vappID)
	if err != nil {
		t.Errorf("Failed to poweroff vapp. Error: %v\n", err)
	} else {
		taskJson, _ := json.MarshalIndent(task, "", "\t")
		t.Logf("taskInfo: %v\n", string(taskJson))
	}
}

func TestPowerOffVappSafely(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vappID := "8a878bda781f145e01784eed976501f7"
	vappClient := NewVappService(icsConnection.Client)
	task, err := vappClient.PowerOffVappSafely(ctx, vappID)
	if err != nil {
		t.Errorf("Failed to poweroff vapp safely. Error: %v\n", err)
	} else {
		taskJson, _ := json.MarshalIndent(task, "", "\t")
		t.Logf("taskInfo: %v\n", string(taskJson))
	}
}

func TestRestartVapp(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vappID := "8a878bda781f145e01784eed976501f7"
	vappClient := NewVappService(icsConnection.Client)
	task, err := vappClient.RestartVapp(ctx, vappID)
	if err != nil {
		t.Errorf("Failed to restart vapp. Error: %v\n", err)
	} else {
		taskJson, _ := json.MarshalIndent(task, "", "\t")
		t.Logf("taskInfo: %v\n", string(taskJson))
	}
}
