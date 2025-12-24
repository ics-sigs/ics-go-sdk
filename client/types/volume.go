package types

type Volume struct {
	ID                  string          `json:"id"`
	UUID                string          `json:"uuid"`
	Size                float64         `json:"size"`
	SizeInByte          int             `json:"sizeInByte"`
	RealSize            float64         `json:"realSize"`
	RealSizeInByte      float64         `json:"realSizeInByte"`
	Name                string          `json:"name"`
	FileName            string          `json:"fileName"`
	Offset              int             `json:"offset"`
	Shared              bool            `json:"shared"`
	DeleteModel         string          `json:"deleteModel"`
	VolumePolicy        string          `json:"volumePolicy"`
	Format              string          `json:"format"`
	BlockDeviceID       string          `json:"blockDeviceId,omitempty"`
	DiskType            string          `json:"diskType,omitempty"`
	DataStoreID         string          `json:"dataStoreId"`
	DataStoreName       string          `json:"dataStoreName"`
	DataStoreSize       float64         `json:"dataStoreSize"`
	DataStoreSizeInByte int             `json:"dataStoreSizeInByte,omitempty"`
	FreeStorage         float64         `json:"freeStorage"`
	DataStoreType       string          `json:"dataStoreType"`
	DataStoreReplicate  int             `json:"dataStoreReplicate"`
	VMName              string          `json:"vmName,omitempty"`
	VMStatus            string          `json:"vmStatus,omitempty"`
	Type                string          `json:"type,omitempty"`
	Description         string          `json:"description,omitempty"`
	Bootable            bool            `json:"bootable"`
	VolumeStatus        string          `json:"volumeStatus"`
	MountedHostIds      []string        `json:"mountedHostIds"`
	Md5                 string          `json:"md5,omitempty"`
	DataSize            int             `json:"dataSize"`
	OpenStackID         string          `json:"openStackId,omitempty"`
	VvSourceDto         interface{}     `json:"vvSourceDto"`
	FormatDisk          bool            `json:"formatDisk"`
	ToBeConverted       bool            `json:"toBeConverted"`
	RelatedVms          []RelatedVmInfo `json:"relatedVms"`
	XactiveDataStoreId  string          `json:"xactiveDataStoreId"`
	ClusterSize         int             `json:"clusterSize"`
	ScsiID              string          `json:"scsiId"`
	SecondaryUuid       string          `json:"secondaryUuid"`
	Usage               string          `json:"usage"`
	SecondaryVolumes    interface{}     `json:"secondaryVolumes"`
	MountDir            interface{}     `json:"mountDir"`
	OpenMode            interface{}     `json:"openMode"`
	Encrypted           bool            `json:"encrypted"`
	DatastoreStrategy   interface{}     `json:"datastoreStrategy"`
	Replicate           int             `json:"replicate"`
	WorkMode            interface{}     `json:"workMode"`
	SerialNumber        string          `json:"serialNumber"`
	AutoSerialNumber    bool            `json:"autoSerialNumber"`
	ProtectType         string          `json:"protectType"`
	CdpShareMemory      int             `json:"cdpShareMemory"`
	HostIds             interface{}     `json:"hostIds"`
	SnapShotCount       int             `json:"snapShotCount"`
	Unmap               bool            `json:"unmap"`
}

type RelatedVmInfo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Text  string `json:"text"`
	State string `json:"state"`
}

type VolumeReq struct {
	Name          string `json:"name"`
	Size          string `json:"size"`
	DataStoreType string `json:"dataStoreType"`
	DataStoreId   string `json:"dataStoreId"`
	VolumePolicy  string `json:"volumePolicy"`
	Description   string `json:"description"`
	Bootable      bool   `json:"bootable"`
	Shared        bool   `json:"shared"`
	Format        string `json:"format"`
	Usage         string `json:"usage"`
}

type VolumeListRsp struct {
	Items []Volume `json:"items"`
}
