package volume

import (
	"context"
	"encoding/json"
	"fmt"
	icsgo "github.com/ics-sigs/ics-go-sdk"
	"github.com/ics-sigs/ics-go-sdk/client/types"
	"testing"
)

var volumeClient *VolumeService
var icsConnection icsgo.ICSConnection
var ctx context.Context
var volumeid string

func TestCreateVolume(t *testing.T) {
	fmt.Println("********************CreateVolume**************")
	icsConnection = icsgo.ICSConnection{
		Username: "admin",
		Password: "Cloud@s1",
		//AccessKeyID:     "09cb1VN5xu5T7469d1Yo",
		//AccessKeySecret: "6e4711WY1pY310qPG8ntv7RdIuM8vDO07XKM2GeX",
		Hostname: "10.68.22.57",
		Port:     "443",
		Locale:   "en",
		Insecure: true,
	}
	ctx = context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}
	createVolumeReq := types.VolumeReq{
		Name: "vol_test001",
		Size: "20",
		//DataStoreId:   "8ab1a2968145ef35018145fc98ee0097", // 10.49.34.22
		//DataStoreId:   "8ab0b34979c154880179c207ef05004d", // 10.49.34.23
		//DataStoreId:   "8ab1a21f8e11b0f9018e11b4c19f0016", // 10.49.34.159
		//DataStoreId:   "8ab1a2218d55e067018d55ec81d10042", // 10.49.34.161
		DataStoreId:   "8ac496b993b4ea2e0193dc9eaafb5d07", // 10.49.34.162
		DataStoreType: "LOCAL",
		VolumePolicy:  "THIN",
		Format:        "RAW",
		Description:   "k8s",
		Usage:         "POD",
		Bootable:      false,
		Shared:        false,
	}
	volumeClient = NewVolumeService(icsConnection.Client)
	task, err := volumeClient.CreateVolume(ctx, createVolumeReq)
	if err != nil {
		t.Fatalf("Failed to Create volume. Error: %v", err)
	}
	t.Logf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err := volumeClient.TraceTaskProcess(&task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Status: %v\n", string(taskJson))
	}
}

func TestGetVolumesInDatastore(t *testing.T) {
	icsConnection = icsgo.ICSConnection{
		Username: "admin",
		Password: "Cloud@s1",
		//AccessKeyID:     "09cb1VN5xu5T7469d1Yo",
		//AccessKeySecret: "6e4711WY1pY310qPG8ntv7RdIuM8vDO07XKM2GeX",
		Hostname: "10.49.34.159",
		Port:     "443",
		Insecure: true,
	}

	//datastoreID := "8ab1a2968145ef35018145fc98ee0097"  // 10.49.34.22
	//datastoreID := "8ab0b34979c154880179c207ef05004d" // 10.49.34.23
	datastoreID := "8ab1a21f8e11b0f9018e11b4c19f0016" // 10.49.34.159
	//datastoreID := "8ab1a2218d55e067018d55ec81d10042" // 10.49.34.161
	//datastoreID := "8ab1a2228e07312e018e0736db0e0016" // 10.49.34.162
	ctx = context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	if icsConnection.AccessKeyID == "" {
		fmt.Println("access key id emtpy")
	} else {
		fmt.Printf("access key id %s\n", icsConnection.AccessKeyID)
	}

	volumeClient = NewVolumeService(icsConnection.Client)
	fmt.Println("********************GetVolumesInDatastore**************")
	volumes, err := volumeClient.GetVolumesInDatastore(ctx, datastoreID)
	if volumes != nil {
		for i := 0; i < len(volumes); i++ {
			volumeinfo := volumes[i]
			volumeJson, _ := json.MarshalIndent(volumeinfo, "", "\t")
			t.Logf("Volume Info: %s\n", string(volumeJson))
		}
	} else {
		fmt.Println(err.Error())
	}
}

func TestGetVolumeInfoById(t *testing.T) {
	icsConnection = icsgo.ICSConnection{
		Username: "admin",
		Password: "Cloud@s1",
		Hostname: "10.49.34.162",
		Port:     "443",
		Insecure: true,
	}
	ctx = context.Background()
	icsConnection.Connect(ctx)
	//volumeID := "8ab1a21f8e11b0f9018e12522ea60027" //10.49.34.159
	//volumeID := "8ab1a2218dbb7668018dc22cd9060030" //10.49.34.161
	volumeID := "8ab1a2228e07312e018e0e0e83f50021" //10.49.34.162
	volumeClient = NewVolumeService(icsConnection.Client)
	fmt.Println("********************TestGetVolumeInfoById**************")
	volumeinfo, err := volumeClient.GetVolumeInfoById(ctx, volumeID)
	if err == nil {
		volumeJson, _ := json.MarshalIndent(volumeinfo, "", "\t")
		t.Logf("Volume Info: %s\n", string(volumeJson))
	} else {
		fmt.Println(err.Error())
	}
}

func TestDeleteVolume(t *testing.T) {
	fmt.Println("********************TestDeleteVolume**************")
	icsConnection = icsgo.ICSConnection{
		Username: "admin",
		Password: "Cloud@s1",
		Hostname: "10.49.34.161",
		Port:     "443",
		Insecure: true,
	}
	ctx = context.Background()
	icsConnection.Connect(ctx)
	volumeClient = NewVolumeService(icsConnection.Client)
	isNeedAuth, err := volumeClient.IsDeleteNeedIdentityAuth(ctx)
	if err != nil {
		t.Fatalf("Failed to get login policy. Error: %v", err)
	}

	var task types.Task
	volumeid = "8ab1a2218dc239da018dc25772a00021"
	if !isNeedAuth {
		task, err = volumeClient.DeleteVolume(ctx, volumeid, true)
	} else {
		task, err = volumeClient.DeleteVolumeWithCheckParams(ctx, volumeid, true, icsConnection.Password)
	}

	if err != nil {
		t.Fatalf("Failed to delete volume. Error: %v", err)
	} else {
		t.Logf("Waiting task %v finish.....\n", task.TaskId)
		taskInfo, err := volumeClient.TraceTaskProcess(&task)
		if err != nil {
			t.Fatalf("Failed to trace task. Error: %v", err)
		} else {
			taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
			t.Logf("Task Status: %v\n", string(taskJson))
		}
	}
}

func TestSetVolume(t *testing.T) {
	fmt.Println("********************TestSetVolume**************")

	icsConnection = icsgo.ICSConnection{
		Username: "admin",
		Password: "Cloud@s1",
		Hostname: "10.49.34.162",
		Port:     "443",
		Insecure: true,
	}
	ctx = context.Background()
	icsConnection.Connect(ctx)
	//volumeId := "8ab1a296868272db0186ce6476fa0073" //10.49.34.22
	//volumeId := "8ab1a297860c01270186ce646ec701fa" //10.49.34.23
	//volumeId := "8ab1a21f8e11b0f9018e12522ea60027" //10.49.34.159
	//volumeId := "8ab1a2218dbb7668018dc22cd9060030" //10.49.34.161
	volumeId := "8ab1a2228e07312e018e0e0e83f50021" //10.49.34.162
	volumeClient = NewVolumeService(icsConnection.Client)
	volumeInfo, err := volumeClient.GetVolumeInfoById(ctx, volumeId)
	if err != nil {
		t.Fatalf("Failed to get volume info. Error: %v", err)
	}

	volumeInfo.Size += 10
	task, err := volumeClient.SetVolume(ctx, volumeId, volumeInfo)
	if err != nil {
		t.Fatalf("Failed to set volume. Error: %v", err)
	}

	t.Logf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err := volumeClient.TraceTaskProcess(&task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Status: %v\n", string(taskJson))
	}
}
