package types

type NetworkPageResponse struct {
	TotalPage   int       `json:"totalPage"`
	CurrentPage int       `json:"currentPage"`
	TotalSize   int       `json:"totalSize"`
	Items       []Network `json:"items"`
}

type Switch struct {
	ID               string      `json:"id"`
	Name             string      `json:"name"`
	ResourceID       string      `json:"resourceId"`
	ControllerIP     string      `json:"controllerIP"`
	DataCenterDto    Datacenter  `json:"dataCenterDto"`
	HostDtos         []Host      `json:"hostDtos"`
	SwitchType       string      `json:"switchType"`
	AppType          string      `json:"appType"`
	Description      string      `json:"description"`
	NetworkDtos      interface{} `json:"networkDtos"`
	SdnNetworkDtos   interface{} `json:"sdnNetworkDtos"`
	VMDtos           interface{} `json:"vmDtos"`
	HostNum          int         `json:"hostNum"`
	PnicNum          int         `json:"pnicNum"`
	NetworkNum       int         `json:"networkNum"`
	VMNum            int         `json:"vmNum"`
	ProdNum          int         `json:"prodNum"`
	MaxVfs           int         `json:"maxvfs"`
	ThirdPartySDN    bool        `json:"thirdPartySDN"`
	Hierarchy        bool        `json:"hierarchy"`
	ConnectStorage   bool        `json:"connectStorage"`
	ConnectManage    bool        `json:"connectManage"`
	ConnectSwitches  interface{} `json:"connectSwitches"`
	DhcpProtection   bool        `json:"dhcpProtection"`
	NeutronName      interface{} `json:"neutronName"`
	NeutronPassword  interface{} `json:"neutronPassword"`
	ConnectScvm      bool        `json:"connectScvm"`
	SwitchUplinkType string      `json:"switchUplinkType"`
	ComputerNetNum   int         `json:"computerNetNum"`
	DataNetNum       int         `json:"dataNetNum"`
	MigrateNetNum    int         `json:"migrateNetNum"`
	VmMigBandWidth   string      `json:"vmMigBandWidth"`
	EnableDpdk       bool        `json:"enableDpdk"`
	SflowStatus      bool        `json:"sflowStatus"`
	NetflowStatus    bool        `json:"netflowStatus"`
	MulticastStatus  bool        `json:"multicastStatus"`
	MirrorStatus     bool        `json:"mirrorStatus"`
	BrLimitStatus    bool        `json:"brLimitStatus"`
	Hidden           bool        `json:"hidden"`
	NetworkTopoly    bool        `json:"networkTopoly"`
	ArbitrativeIp    string      `json:"arbitrativeIp"`
	HbAutoCreate     bool        `json:"hbAutoCreate"`
	EnableFcoe       bool        `json:"enableFcoe"`
	EnableSc         bool        `json:"enableSc"`
	EnableTrust      bool        `json:"enableTrust"`
	MTU              int         `json:"mtu"`
	ConnectInCloudOs bool        `json:"connectInCloudOs"`
	IamServerDto     interface{} `json:"iamServerDto"`
}

type Network struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	ResourceID      string        `json:"resourceId"`
	Vlan            int           `json:"vlan"`
	VlanFlag        bool          `json:"vlanFlag"`
	Mtu             int           `json:"mtu"`
	Type            string        `json:"type"`
	VswitchDto      Switch        `json:"vswitchDto"`
	PnicDto         interface{}   `json:"pnicDto"`
	PortDtos        []interface{} `json:"portDtos"`
	VMDtos          interface{}   `json:"vmDtos"`
	VnicDtos        interface{}   `json:"vnicDtos"`
	VmCount         int           `json:"vmcount"`
	VnicCount       int           `json:"vniccount"`
	ConnectMode     string        `json:"connectMode"`
	Description     string        `json:"description"`
	UplinkRate      int           `json:"uplinkRate"`
	UplinkBurst     int           `json:"uplinkBurst"`
	DownlinkRate    int           `json:"downlinkRate"`
	DownlinkBurst   int           `json:"downlinkBurst"`
	QosEnabled      bool          `json:"qosEnabled"`
	DataServiceType interface{}   `json:"dataServiceType"`
	UserVlan        interface{}   `json:"userVlan"`
	TpidType        interface{}   `json:"tpidType"`
	PermitDel       bool          `json:"permitDel"`
	Cidr            string        `json:"cidr"`
	CidrType        int           `json:"cidrType"`
	Gateway         string        `json:"gateway"`
	DhcpEnabled     bool          `json:"dhcpEnabled"`
	GatewayEnabled  bool          `json:"gatewayEnabled"`
	DNS             string        `json:"dns"`
	DataCenterDto   interface{}   `json:"dataCenterDto"`
	NetworkTopoly   bool          `json:"networkTopoly"`
	UseTypes        interface{}   `json:"useTypes"`
	StartIp         string        `json:"startIp"`
	EndIp           string        `json:"endIp"`
	Pools           []interface{} `json:"pools"`
	UsedByHbLink    bool          `json:"usedByHbLink"`
}

type SdnNetworkPageResponse struct {
	TotalPage   int          `json:"totalPage"`
	CurrentPage int          `json:"currentPage"`
	TotalSize   int          `json:"totalSize"`
	Items       []SdnNetwork `json:"items"`
}

type SdnNetwork struct {
	ID                string      `json:"id"`
	Name              string      `json:"name"`
	DataCenterDto     Datacenter  `json:"dataCenterDto"`
	SubnetKeys        []SdnSubnet `json:"subnetKeys"`
	Status            string      `json:"status"`
	RouterExternal    bool        `json:"routerExternal"`
	NetworkTopoly     bool        `json:"networkTopoly"`
	ProviderSegID     string      `json:"providerSegID"`
	NetworkType       string      `json:"networkType"`
	ProviderPhyNet    string      `json:"providerPhyNet"`
	SubnetDto         interface{} `json:"subnetDto"`
	SubnetCount       int         `json:"subnetCount"`
	PortCount         int         `json:"portCount"`
	VSwitchID         string      `json:"vswitchId,omitempty"`
	Type              string      `json:"type"`
	RouterDtos        interface{} `json:"routerDtos"`
	FirewallDto       interface{} `json:"firewallDto"`
	UsedIp            int         `json:"usedIp"`
	TotalIp           int         `json:"totalIp"`
	MiredirectStatus  bool        `json:"miredirectStatus"`
	MulticastStatus   bool        `json:"multicastStatus"`
	PortForwardingCnt int         `json:"portForwardingCnt"`
	L2GatewayStatus   bool        `json:"l2GatewayStatus"`
	QosEnabled        bool        `json:"qosEnabled"`
	UplinkRate        int         `json:"uplinkRate"`
	DownlinkRate      int         `json:"downlinkRate"`
}

type SdnSubnet struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Cidr       string `json:"cidr"`
	EnableDHCP bool   `json:"enableDHCP"`
	Gateway    string `json:"gateway"`
	RouterID   string `json:"routerId,omitempty"`
	RouterName string `json:"routerName,omitempty"`
	SubnetID   string `json:"subnetId"`
	SubnetName string `json:"subnetName"`
	Pools      []struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"pools"`
	NetworkTopoly bool `json:"networkTopoly"`
}
