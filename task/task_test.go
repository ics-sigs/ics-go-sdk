package task

import (
	"context"
	"encoding/json"
	icsgo "github.com/ics-sigs/ics-go-sdk"
	icsvm "github.com/ics-sigs/ics-go-sdk/vm"
	"testing"
)

var (
	icsConnection = &icsgo.ICSConnection{
		Username: "admin",
		Password: "admin@inspur", // "Cloud@s1"
		Hostname: "10.7.11.90",   // "10.48.50.13"
		Port:     "443",
		Insecure: true,
	}
)

func TestGetTaskInfo(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := icsvm.NewVirtualMachineService(icsConnection.Client)
	vm, err := vmClient.GetVM(ctx, "8a878bda6f6f3ca4016f6f6eb8d300bb")
	if err != nil {
		t.Fatalf("Failed to get vm info by specified id. Error: %v", err)
	}

	vm.Name = "11.82_zhanghuijian00"
	vm.VncPasswd = "12345678"
	task, err := vmClient.SetVM(ctx, *vm)
	if err != nil {
		t.Fatalf("Failed to set vm. Error: %v", err)
	} else {
		t.Logf("Set VM Task: %+v\n", task)
	}

	taskClient := NewTaskService(icsConnection.Client)
	taskInfo, err := taskClient.GetTaskInfo(ctx, task)
	if err != nil {
		t.Fatalf("Failed to get task info. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Info: %v\n", string(taskJson))
	}
}

func TestWaitForResult(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := icsvm.NewVirtualMachineService(icsConnection.Client)
	vm, err := vmClient.GetVM(ctx, "8a878bda6f6f3ca4016f6f6eb8d300bb")
	if err != nil {
		t.Fatalf("Failed to get vm info by specified id. Error: %v", err)
	}

	vm.Name = "11.82_zhanghuijian00"
	vm.VncPasswd = "12345678"
	task, err := vmClient.SetVM(ctx, *vm)
	if err != nil {
		t.Fatalf("Failed to set vm. Error: %v", err)
	} else {
		t.Logf("Set VM Task: %+v\n", task)
	}

	taskClient := NewTaskService(icsConnection.Client)
	taskInfo, err := taskClient.WaitForResult(ctx, task)
	if err != nil {
		t.Fatalf("Failed to get task info. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Info: %v\n", string(taskJson))
	}
}
