package models

type Instance struct {
	/*项目id*/
	ProjectID string `json:"projectID"`
	/*az名称*/
	AzName string `json:"azName"`
	/*附加卷*/
	AttachedVolume []string `json:"attachedVolume"`
	/*网络地址信息*/
	Addresses []Addresses `json:"addresses"`
	/*资源id	*/
	ResourceID string `json:"resourceID"`
	/*网络地址信息*/
	InstanceID string `json:"instanceID"`
	/*云主机名称*/
	DisplayName string `json:"displayName"`
	/*主机名称*/
	InstanceName string `json:"instanceName"`
	/*操作系统类型，详见操作系统类型说明	*/
	OsType int `json:"osType"`
	/*主机状态*/
	InstanceStatus string `json:"instanceStatus"`
	/*到期时间*/
	ExpiredTime string `json:"expiredTime"`
	/*可用(天)*/
	AvailableDay int `json:"availableDay"`
	/*更新时间*/
	UpdatedTime string `json:"updatedTime"`
	/*更新时间*/
	CreatedTime string `json:"createdTime"`
	/*监控对象名称*/
	ZabbixName string `json:"zabbixName"`
	/*安全组信息*/
	SecGroupList []SecGroupList `json:"secGroupList"`
	/*内网ipv4地址*/
	PrivateIP string `json:"privateIP"`
	/*内网ipv6址*/
	PrivateIPv6 string `json:"privateIPv6"`
	/*网卡信息*/
	NetworkCardList []networkCardList `json:"networkCardList"`
	/*虚拟ip信息列表*/
	VipInfoList []VipInfoList `json:"vipInfoList"`
	/*vip数目*/
	VipCount int `json:"vipCount"`
	/*云主机组信息*/
	AffinityGroup AffinityGroup `json:"affinityGroup"`
	/*镜像信息*/
	Image image `json:"image"`
	/*规格信息*/
	Flavor Flavor `json:"flavor"`
	/*付费方式 ，true表示按量付费; false为包周期*/
	OnDemand bool `json:"onDemand"`
	/*vpc名称*/
	VpcName string `json:"vpcName"`
	/*vpc ID*/
	VpcID string `json:"vpcID"`
	/*固定IP*/
	FixedIP []string `json:"fixedIP"`
	/*弹性ip	*/
	FloatingIP string `json:"floatingIP"`
	/*子网ID列表*/
	SubnetIDList []string `json:"subnetIDList"`
	/*密钥对名称		*/
	KeypairName string `json:"keypairName"`
}

type Addresses struct {
	VpcName     string        `json:"vpcName"`
	AddressList []AddressList `json:"addressList"`
}
type AddressList struct {
	Addr    string `json:"addr"`
	Version int    `json:"version"`
	Type    string `json:"type"`
}
type SecGroupList struct {
	SecurityGroupName string `json:"securityGroupName"`
	SecurityGroupID   string `json:"securityGroupID"`
}
type networkCardList struct {
	IPv4Address   string   `json:"IPv4Address"`
	IPv6Address   string   `json:"IPv6Address"`
	IsMaster      bool     `json:"isMaster"`
	SubnetCidr    string   `json:"subnetCidr"`
	networkCardID string   `json:"networkCardID"`
	gateway       string   `json:"gateway"`
	securityGroup []string `json:"securityGroup"`
	subnetID      string   `json:"subnetID"`
}

type VipInfoList struct {
	vipID          string `json:"vipID"`
	vipAddress     string `json:"vipAddress"`
	vipBindNicIP   string `json:"vipBindNicIP"`
	vipBindNicIPv6 string `json:"vipBindNicIPv6"`
	nicID          string `json:"nicID"`
}

type AffinityGroup struct {
	AffinityGroupPolicy string `json:"affinityGroupPolicy"`
	AffinityGroupName   string `json:"affinityGroupName"`
	AffinityGroupID     string `json:"affinityGroupID"`
}

type image struct {
	ImageID   string `json:"imageID"`
	ImageName string `json:"imageName"`
}
type Flavor struct {
	FlavorID     string `json:"flavorID"`
	FlavorName   string `json:"flavorName"`
	FlavorCPU    int    `json:"flavorCPU"`
	FlavorRAM    int    `json:"flavorRAM"`
	GpuType      string `json:"gpuType"`
	GpuCount     int    `json:"gpuCount"`
	GpuVendor    string `json:"gpuVendor"`
	VideoMemSize int    `json:"videoMemSize"`
}
