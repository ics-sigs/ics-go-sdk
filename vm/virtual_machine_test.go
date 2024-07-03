package vm

import (
	"context"
	"encoding/json"
	"fmt"
	icsgo "github.com/ics-sigs/ics-go-sdk"
	"github.com/ics-sigs/ics-go-sdk/client/types"
	"testing"
)

var (
	icsConnection = &icsgo.ICSConnection{
		Username: "admin",
		Password: "Cloud@s1",
		Hostname: "10.49.34.162",
		Port:     "443",
		Insecure: true,
	}
)

func TestGetVMInfo(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmId := "8ab1a297881ce9fc01884d15f13500a4" //10.49.34.23 on
	//vmId := "8ab1a2968df9271d018e1140871501a2" //10.49.34.22 on
	//vmId := "8ab1a21f8e11b0f9018e1254eb8d002d" //10.49.34.159 on
	//vmId := "8ab1a21f8e11b0f9018e1d305183006c" //10.49.34.159 off
	//vmId := "8ab1a2218dc27ce1018dc710c7380081" //10.49.34.161 on
	//vmId := "8ab1a2218de95b48018deecf92be00ea" //10.49.34.161 off
	//vmId := "8ab1a2228e07312e018e0e10d39f0027" //10.49.34.162 on
	//vmId := "8ab1a2228e07312e018e1d1f9cd10075" //10.49.34.162 off
	vmClient := NewVirtualMachineService(icsConnection.Client)
	vm, err := vmClient.GetVM(ctx, vmId)
	if err != nil {
		t.Errorf("Failed to get vm info by specified id. Error: %v", err)
	} else {
		vmJson, _ := json.MarshalIndent(vm, "", "\t")
		t.Logf("VM Info: %s\n", string(vmJson))
	}
}

func TestSetVM(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmId := "8ab1a21f8e11b0f9018e1d8c261f0099" //10.49.34.159 on
	//vmId := "8ab1a21f8e11b0f9018e1d305183006c" //10.49.34.159 off
	//vmId := "8ab1a2218dc27ce1018dc710c7380081" //10.49.34.161 on
	//vmId := "8ab1a2218de95b48018deecf92be00ea" //10.49.34.161 off
	//vmId := "8ab1a2228e07312e018e0e10d39f0027" //10.49.34.162 on
	//vmId := "8ab1a2228e07312e018e1d1f9cd10075" //10.49.34.162 off
	vmClient := NewVirtualMachineService(icsConnection.Client)
	vm, err := vmClient.GetVM(ctx, vmId)
	if err != nil {
		t.Fatalf("Failed to get vm info by specified id. Error: %v", err)
	}

	//vm.Name = vm.Name + "x"
	//vm.Disks[0].Volume.Size += 10
	if len(vm.Disks) > 1 {
		//remove the last disk
		vm.Disks = vm.Disks[:len(vm.Disks)-1]
	}
	//vm.VncPasswd = "12345678"
	task, err := vmClient.SetVM(ctx, *vm)
	if err != nil {
		t.Fatalf("Failed to set vm. Error: %v", err)
	} else {
		t.Logf("Set VM Task: %+v\n", task)
	}

	taskInfo, err := vmClient.GetTaskInfo(task)
	if err != nil {
		t.Fatalf("Failed to get task info. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Info: %v\n", string(taskJson))
	}

	t.Logf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err = vmClient.TraceTaskProcess(task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Status: %v\n", string(taskJson))
	}
}

func TestGetVMByUUID(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)

	vm, err := vmClient.GetVMByUUID(ctx, "71909586-f83b-45de-b620-4650aceb7865")
	if err != nil || vm == nil {
		t.Errorf("Failed to get vm info by specified uuid. Error: %v", err)
	} else {
		vmJson, _ := json.MarshalIndent(vm, "", "\t")
		t.Logf("VM Info: %s\n", string(vmJson))
	}
}

func TestGetVMByIP(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)

	vm, err := vmClient.GetVMByIP(ctx, "10.7.11.81")
	if vm != nil {
		fmt.Printf("%+v", vm)
	} else {
		fmt.Println("No VM be found by specified id.")
	}
}

func TestGetVMByName(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)

	vm, err := vmClient.GetVMByName(ctx, "master")
	if vm != nil {
		fmt.Printf("%+v", vm)
	} else {
		fmt.Println("No VM be found by specified id.")
	}
}

func TestVMPageList(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)

	pakg := types.PageReq{
		1000,
		1,
		"",
		"desc",
	}

	req := &types.VMPageReq{
		pakg,
	}
	vmpagelist, err := vmClient.VMPageList(req)
	if vmpagelist != nil {
		fmt.Printf("%+v", vmpagelist)
	}
}

func TestPowerOnVM(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmId := "8ab0b28d77be994a0178110f28f110ce"
	vmClient := NewVirtualMachineService(icsConnection.Client)
	task, err := vmClient.PowerOnVM(ctx, vmId)
	if err != nil {
		t.Fatalf("Failed to poweron vm. Error: %v", err)
	} else {
		t.Logf("PowerOn VM Task: %+v\n", task)
	}

	t.Logf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err := vmClient.TraceTaskProcess(task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Status: %v\n", string(taskJson))
	}
}

func TestPowerOffVM(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmId := "8ab0b28d77be994a0178110f28f110ce"
	vmClient := NewVirtualMachineService(icsConnection.Client)
	task, err := vmClient.PowerOffVM(ctx, vmId)
	if err != nil {
		t.Fatalf("Failed to poweroff vm. Error: %v", err)
	} else {
		t.Logf("PowerOff VM Task: %+v\n", task)
	}

	t.Logf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err := vmClient.TraceTaskProcess(task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Status: %v\n", string(taskJson))
	}
}

func TestShutdownVM(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmId := "8ab0b28d77be994a0178110f28f110ce"
	vmClient := NewVirtualMachineService(icsConnection.Client)
	task, err := vmClient.ShutdownVM(ctx, vmId)
	if err != nil {
		t.Fatalf("Failed to shutdown vm. Error: %v", err)
	} else {
		t.Logf("Shutdown VM Task: %+v\n", task)
	}

	t.Logf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err := vmClient.TraceTaskProcess(task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Status: %v\n", string(taskJson))
	}
}

func TestRestartVM(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmId := "8a878bda6f6f3ca4016f6f6eb8d300bb"
	vmClient := NewVirtualMachineService(icsConnection.Client)
	task, err := vmClient.RestartVM(ctx, vmId)
	if err != nil {
		t.Fatalf("Failed to restart vm. Error: %v", err)
	} else {
		t.Logf("Restart VM Task: %+v\n", task)
	}

	t.Logf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err := vmClient.TraceTaskProcess(task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Status: %v\n", string(taskJson))
	}
}

func TestDeleteVM(t *testing.T) {
	var task *types.Task
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmId := "8ab1a2218dc27ce1018dc6e0a218004f"
	deleteFile := true
	removeData := true

	vmClient := NewVirtualMachineService(icsConnection.Client)
	needAuth, err := vmClient.IsDeleteNeedIdentityAuth(ctx)
	if needAuth {
		t.Logf("Delete VM %s with check params", vmId)
		task, err = vmClient.DeleteVMWithCheckParams(ctx, vmId, deleteFile, removeData, icsConnection.Password)
	} else {
		task, err = vmClient.DeleteVM(ctx, vmId, deleteFile, removeData)
	}
	if err != nil {
		t.Fatalf("Failed to delete vm. Error: %v", err)
	} else {
		t.Logf("Delete VM Task: %+v\n", task)
	}

	t.Logf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err := vmClient.TraceTaskProcess(task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Status: %v\n", string(taskJson))
	}
}

func TestGetVMTemplateList(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)
	vmtList, err := vmClient.GetVMTemplateList(ctx)
	if err != nil {
		t.Errorf("Failed to get vm template list. Error: %v", err)
	} else {
		t.Logf("VMTemplateList: %+v", vmtList)
	}
}

func TestGetVMTemplate(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmtID := "8ab0b28d77be994a0177fc16f5b20ab0"
	vmClient := NewVirtualMachineService(icsConnection.Client)
	vmt, err := vmClient.GetVMTemplate(ctx, vmtID)
	if err != nil {
		t.Errorf("Failed to get vm template by id. Error: %v", err)
	} else {
		vmtJson, _ := json.MarshalIndent(vmt, "", "\t")
		t.Logf("VMTemplate: %v", string(vmtJson))
	}
}

func TestGetVMTemplateByUUID(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	uuid := "b9a6059a-5057-4efa-b501-78e2cb8d3e64"
	vmClient := NewVirtualMachineService(icsConnection.Client)
	vmt, err := vmClient.GetVMTemplateByUUID(ctx, uuid)
	if err != nil {
		t.Errorf("Failed to get vm template by uuid. Error: %v", err)
	} else {
		vmtJson, _ := json.MarshalIndent(vmt, "", "\t")
		t.Logf("VMTemplate: %v", string(vmtJson))
	}
}

func TestGetVMTemplateByName(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmtName := "template_test"
	vmClient := NewVirtualMachineService(icsConnection.Client)
	vmt, err := vmClient.GetVMTemplateByName(ctx, vmtName)
	if err != nil {
		t.Errorf("Failed to get vm template by name. Error: %v", err)
	} else {
		vmtJson, _ := json.MarshalIndent(vmt, "", "\t")
		t.Logf("VMTemplate: %v", string(vmtJson))
	}
}

func TestCreateVMByTemplate(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmtName := "template_test"
	vmClient := NewVirtualMachineService(icsConnection.Client)
	vmt, err := vmClient.GetVMTemplateByName(ctx, vmtName)
	if err != nil {
		t.Fatalf("Failed to get vm template by name. Error: %v", err)
	} else {
		vmtJson, _ := json.MarshalIndent(vmt, "", "\t")
		t.Logf("VMTemplate: %v\n", string(vmtJson))
	}

	quickClone := false
	vmt.Name = "vm_create_by_template_001"
	task, err := vmClient.CreateVMByTemplate(ctx, *vmt, quickClone)
	if err != nil {
		t.Fatalf("Failed to create vm by template. Error: %v", err)
	} else {
		t.Logf("Clone VM Task: %+v\n", task)
	}

	t.Logf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err := vmClient.TraceTaskProcess(task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Status: %v\n", string(taskJson))
	}

}

func TestGetVMPowerStateByID(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)
	status, err := vmClient.GetVMPowerStateByID(ctx, "8a878bda781f145e0178456029a8016c")
	if status != nil {
		fmt.Println(*status)
	} else {
		fmt.Println("No VM be found by you point id.")
	}
}

func TestGetVMNetState(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)
	nic, err := vmClient.GetVMNetState(ctx, "8a878bda6f7012c7016f70b40ed000a1")
	//nic, err := vmClient.GetVMNetState(ctx, "8a878bda781f145e0178456029a8016c")

	if err != nil {
		t.Fatalf("No VM be found by you point id.")
	} else {
		for _, data := range nic {
			fmt.Println("-----------------------")
			fmt.Println(data)
			fmt.Println("-----------------------")
		}
	}
}

func TestGetOvaConfig(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)
	ovaFilePath := "/datastore/2afee96f-4242-4841-926d-e2553b1a3ca5/vm_rh7.6min.ova"
	hostUUID := "56c940ba-f62d-423f-aaf3-f91dfc22d7ea"
	imageHostUUID := "56c940ba-f62d-423f-aaf3-f91dfc22d7ea"
	ovaConfig, err := vmClient.GetOvaConfig(ctx, ovaFilePath, hostUUID, imageHostUUID)
	if err != nil {
		t.Errorf("Failed to get ova config info. Error: %v", err)
	} else {
		ovaConfigJson, _ := json.MarshalIndent(ovaConfig, "", "\t")
		t.Logf("VM Info: %s\n", string(ovaConfigJson))
	}
}

func TestImportVM(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)
	ovaFilePath := "/datastore/57f089e7-0165-4868-be63-9bbafebc7225/vm_rh7.6_x86.ova"
	hostUUID := "2cf4bb24-cc1d-46fb-862e-e408da0898ce"
	imageHostUUID := "2cf4bb24-cc1d-46fb-862e-e408da0898ce"
	vmConfig, err := vmClient.GetOvaConfig(ctx, ovaFilePath, hostUUID, imageHostUUID)
	if err != nil {
		t.Fatalf("Failed to get ova config info. Error: %v", err)
	}

	metadata := `
{
    "hostname": "web06.cloud.org",
    "launch_index": 0,
    "name":"web06",
    "uuid":"282181d8-4eea-4438-a928-5a7976283898"
}
`
	userdata := `
#cloud-config
write_files:
- path: /etc/sysconfig/selinux
  encoding: b64
  content: IyBteSBuYW1lIGlzIGNsb3VkLWluaXQ=
  owner: root:root
  permissions: '0644'
- path: /etc/sysconfig/samba
  content: |
    # My new /etc/sysconfig/samba file
    SMBDOPTIONS="-D"
  path: /etc/sysconfig/samba
runcmd:
- echo "123456789" > /tmp/test_file
- ip address show  > /tmp/ip_addr
`
	vmConfig.Name = "vm_create_by_ova_004"
	vmConfig.HostID = hostUUID
	vmConfig.Template = false
	vmConfig.CPUSocket = vmConfig.CPUNum / vmConfig.CPUCore
	vmConfig.DataStoreID = "8ab1a2228e07312e018e0736db0e0016"
	vmConfig.Disks[0].Volume.DataStoreID = "8ab1a2228e07312e018e0736db0e0016"
	vmConfig.Nics[0].DeviceName = "manageNetwork0"
	vmConfig.Nics[0].DeviceID = "8ab1a2228e071ead018e072d9f970037"
	vmConfig.CloudInit = types.CloudInit{
		MetaData:       metadata,
		UserData:       userdata,
		DataSourceType: "OPENSTACK",
	}
	rateLimit := 100
	task, err := vmClient.ImportVM(ctx, *vmConfig, ovaFilePath, imageHostUUID, rateLimit)
	if err != nil {
		t.Fatalf("Failed to import vm by ova. Error: %v", err)
	} else {
		t.Logf("Import VM %s Task: %+v\n", vmConfig.Name, task)
	}

	t.Logf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err := vmClient.TraceTaskProcess(task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Status: %v\n", string(taskJson))
	}

}
