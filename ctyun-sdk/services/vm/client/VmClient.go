package client

import (
	"encoding/json"
	"errors"
	"github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/core"
	vm "github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/services/vm/apis"
)

type VmClient struct {
	core.CtyunClient
}

func NewVmClient(credential *core.Credential) *VmClient {
	if credential == nil {
		return nil
	}

	config := core.NewConfig()
	config.SetEndpoint("ctecs-global.ctapi.ctyun.cn")

	return &VmClient{
		core.CtyunClient{
			Credential:  *credential,
			Config:      *config,
			ServiceName: "basic",
			Revision:    "1.0.8",
			Logger:      core.NewDefaultLogger(core.LogInfo),
		}}
}

func (c *VmClient) SetConfig(config *core.Config) {
	c.Config = *config
}

func (c *VmClient) SetLogger(logger core.Logger) {
	c.Logger = logger
}

/*
 * 为云主机创建私有镜像。
 */
func (c *VmClient) CreateImage(request *vm.CreateImageRequest) (*vm.CreateImageResponse, error) {

	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	c.Config.SetEndpoint("ctimage-global.ctapi.ctyun.cn")
	resp, err := c.Send(request)
	if err != nil {
		c.Logger.Log(core.LogError, "create image failed, resp: %s", string(resp))
		return nil, err
	}
	var ctResp vm.CreateImageResponse
	err = json.Unmarshal(resp, &ctResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return &ctResp, err
}

/*
 * 创建一台云主机实例
 */
func (c *VmClient) CreateInstances(request *vm.CreateInstancesRequest) (*vm.CreateInstancesResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	c.Config.Endpoint = "ctecs-global.ctapi.ctyun.cn"
	resp, err := c.Send(request)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	//var dataMap map[string]string
	var ctResp vm.CreateInstancesResponse
	err = json.Unmarshal(resp, &ctResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return &ctResp, err
}

/*
 * 关闭一台云主机实例
 */
func (c *VmClient) StopInstance(request *vm.StopInstanceRequest) (*vm.StopInstanceResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	c.Config.Endpoint = "ctecs-global.ctapi.ctyun.cn"
	resp, err := c.Send(request)
	if err != nil {
		c.Logger.Log(core.LogError, "Stop Instance failed, resp: %s", string(resp))
		return nil, err
	}
	var ctResp vm.StopInstanceResponse
	err = json.Unmarshal(resp, &ctResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return &ctResp, err
}

/*
 * 查询云主机列表实例
 */
func (c *VmClient) QueryInstancesList(request *vm.QueryInstancesRequest) (*vm.QueryInstancesResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	c.Config.Endpoint = "ctecs-global.ctapi.ctyun.cn"
	resp, err := c.Send(request)
	if err != nil {
		c.Logger.Log(core.LogError, "Query Instance List failed, resp: %s", string(resp))
		return nil, err
	}
	//c.Logger.Log(core.LogError, "ctResp=====================================", string(resp))
	ctResp := &vm.QueryInstancesResponse{}
	err = json.Unmarshal(resp, ctResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, err: %v", err)
		return nil, err
	}

	return ctResp, err
}

/*
 * 查询云主机实例详情
 */
func (c *VmClient) DescribeInstance(request *vm.DescribeInstanceRequest) (*vm.DescribeInstanceResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request)
	if err != nil {
		c.Logger.Log(core.LogError, "Query Instance Describe failed, resp: %s", string(resp))
		return nil, err
	}
	var ctResp vm.DescribeInstanceResponse
	err = json.Unmarshal(resp, &ctResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return &ctResp, err
}

/*
 * 查询镜像详情
 */
func (c *VmClient) DescribeImage(request *vm.DescribeImageRequest) (*vm.DescribeImageResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	c.Config.SetEndpoint("ctimage-global.ctapi.ctyun.cn")
	resp, err := c.Send(request)
	if err != nil {
		c.Logger.Log(core.LogError, "Query Image Describe failed, resp: %s", string(resp))
		return nil, err
	}
	var ctResp vm.DescribeImageResponse
	err = json.Unmarshal(resp, &ctResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return &ctResp, err
}

/*
 *创建ssh密钥对。
 */
func (c *VmClient) CreateKeypair(request *vm.CreateKeypairRequest) (*vm.CreateKeypairResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	c.Config.SetEndpoint("ctecs-global.ctapi.ctyun.cn")
	resp, err := c.Send(request)
	if err != nil {
		return nil, err
	}

	ctResp := &vm.CreateKeypairResponse{}
	err = json.Unmarshal(resp, ctResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return ctResp, err
}

type QueryInstancesResponse struct {
	/*返回状态码（800为成功，900为失败）*/
	StatusCode int `json:"statusCode"`
	/*具体错误码标志*/
	ErrorCode string `json:"errorCode"`
	/*失败时的错误信息*/
	Message string `json:"message"`
	/*失败时的错误描述*/
	Description string `json:"description"`
	/*一般情况只有错误，才会返回详细信息*/
	Details string `json:"details"`
	/*成功时返回的数据，参见returnObj对象结构	*/
	ReturnObj QueryInstancesResult `json:"returnObj"`
}

type QueryInstancesResult struct {
	/*当前页记录数目*/
	CurrentCount int `json:"currentCount"`
	/*总记录数*/
	TotalCount int `json:"totalCount"`
	/*失败时的错误信息*/
	TotalPage int `json:"totalPage"`
	/*失败时的错误描述*/
	Results []results `json:"results"`
}
type results struct {
	/*项目id*/
	ProjectID string `json:"projectID"`
	/*az名称*/
	AzName string `json:"azName"`
	/*附加卷*/
	AttachedVolume []string `json:"attachedVolume"`
	/*网络地址信息*/
	Addresses []addresses `json:"addresses"`
	/*资源id	*/
	ResourceID string `json:"resourceID"`
	/*网络地址信息*/
	InstanceID string `json:"instanceID"`
	/*云主机名称*/
	DisplayName string `json:"displayName"`
	/*主机名称*/
	InstanceName string `json:"instanceName"`
	/*操作系统类型，详见操作系统类型说明	*/
	OsType string `json:"osType"`
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
	SecGroupList []secGroupList `json:"secGroupList"`
	/*内网ipv4地址*/
	PrivateIP string `json:"privateIP"`
	/*内网ipv6址*/
	PrivateIPv6 string `json:"privateIPv6"`
	/*网卡信息*/
	NetworkCardList []networkCardList `json:"networkCardList"`
	/*虚拟ip信息列表*/
	VipInfoList []vipInfoList `json:"vipInfoList"`
	/*vip数目*/
	VipCount int `json:"vipCount"`
	/*云主机组信息*/
	AffinityGroup affinityGroup `json:"affinityGroup"`
	/*镜像信息*/
	Image image `json:"image"`
	/*规格信息*/
	Flavor flavor `json:"flavor"`
	/*付费方式 ，true表示按量付费; false为包周期*/
	OnDemand bool `json:"onDemand"`
	/*vpc名称*/
	VpcName string `json:"vpc名称"`
	/*vpc ID*/
	vpcID string `json:"vpcID"`
	/*固定IP*/
	FixedIP []string `json:"fixedIP"`
	/*弹性ip	*/
	floatingIP string `json:"onDemand"`
	/*子网ID列表*/
	SubnetIDList []string `json:"subnetIDList"`
	/*密钥对名称		*/
	KeypairName string `json:"keypairName"`
}

type addresses struct {
	vpcName     string        `json:"vpcName"`
	AddressList []addressList `json:"addressList"`
}
type addressList struct {
	Addr    string `json:"addr"`
	Version int    `json:"version"`
	Type    string `json:"type"`
}
type secGroupList struct {
	imageID   string `json:"imageID"`
	imageName string `json:"imageName"`
}
type networkCardList struct {
	IPv4Address   string   `json:"iPv4Address"`
	IPv6Address   string   `json:"iPv4Address"`
	IsMaster      bool     `json:"isMaster"`
	SubnetCidr    string   `json:"subnetCidr"`
	networkCardID string   `json:"networkCardID"`
	gateway       string   `json:"gateway"`
	securityGroup []string `json:"gateway"`
	subnetID      string   `json:"subnetID"`
}

type vipInfoList struct {
	vipID          string `json:"vipID"`
	vipAddress     string `json:"vipAddress"`
	vipBindNicIP   string `json:"vipBindNicIP"`
	vipBindNicIPv6 string `json:"vipBindNicIPv6"`
	nicID          string `json:"nicID"`
}

type affinityGroup struct {
	affinityGroupPolicy string `json:"affinityGroupPolicy"`
	affinityGroupName   string `json:"affinityGroupName"`
	affinityGroupID     string `json:"affinityGroupID"`
}

type image struct {
	ImageID   string `json:"imageID"`
	ImageName string `json:"imageName"`
}
type flavor struct {
	FlavorID     string `json:"flavorID"`
	FlavorName   string `json:"flavorName"`
	FlavorCPU    int    `json:"flavorCPU"`
	FlavorRAM    int    `json:"flavorRAM"`
	GpuType      string `json:"gpuType"`
	GpuCount     int    `json:"gpuCount"`
	GpuVendor    string `json:"gpuVendor"`
	VideoMemSize int    `json:"videoMemSize"`
}
