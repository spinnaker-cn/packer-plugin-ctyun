package apis

import (
	"github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/core"
	vm "github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/services/vm/models"
)

type CreateInstancesRequest struct {
	core.CtyunRequest
	/* 描述云主机配置*/
	InstanceSpec *vm.InstanceSpec `json:"instanceSpec"`
	/* 客户端存根，用于保证订单幂等性。要求单个云平台账户内唯一。 (true) */
	ClientToken *string `json:"clientToken"`

	/* 资源池ID。您可以调用资源池列表查询获取最新的资源池列表可查询：https://www.ctyun.cn/document/10026730/10040588。 (true) */
	RegionID *string `json:"regionID"`

	/* 可用区名称，4.0资源池必填。您可以调用资源池可用区查询获取资源池可用区列表：https://www.ctyun.cn/document/10026730/10040590。 (false) */
	AzName *string `json:"azName"`

	/* 云主机名称，只能由数字、字母、-组成，不能以-开头或结尾，不能连续使用-，也不能仅使用数字，且长度为2-15字符。 (false) */
	InstanceName *string `json:"instanceName"`

	/* 云主机显示名称，长度为2-63字符。 (true) */
	DisplayName *string `json:"displayName"`

	/* 规格ID。您可以调用查询一个或多个云主机规格资源获取云主机规格信息。 (true) */
	FlavorID *string `json:"flavorID"`

	/* 本参数表示镜像类型 ，取值范围：0：私有镜像1：公有镜像2：共享镜像3：安全镜像4：甄选镜像。 (true) */
	ImageType *int `json:"imageType"`

	/* 镜像ID。 (true) */
	ImageID *string `json:"imageID"`

	/* 本参数表示系统盘类型 ，取值范围：SATA：普通云盘SAS：SAS云盘SSD-genric：通用SSD云盘SSD：SSD云盘。 (true) */
	BootDiskType *string `json:"bootDiskType"`

	/* 系统盘大小单位为GiB，取值范围[40-2048]。 (true) */
	BootDiskSize *int `json:"bootDiskSize"`

	/* 虚拟私有云ID。 (true) */
	VpcID *string `json:"vpcID"`

	/* 网卡。您可以调用查询网卡列表获取网卡信息及对应的虚拟私有云ID https://www.ctyun.cn/document/10026730/10040207(true) */
	NetworkCardList []vm.NetworkCardList `json:"networkCardList"`

	/* 本参数表示是否使用弹性公网IP ，取值范围：0：不使用1：自动分配2：使用已有 (true) */
	ExtIP *string `json:"extIP"`

	/* 本参数表示购买方式 ，取值范围：false（按周期）true（按需），按周期创建云主机需要同时指定cycleCount和cycleType参数 */
	OnDemand bool `json:"onDemand"`

	/* 订购时长*/
	CycleCount *int `json:"cycleCount"`

	/* 本参数表示订购周期类型 ，取值范围：MONTH：按月YEAR：按年最长订购周期为5年*/
	CycleType *string `json:"cycleType"`

	KeyPairID    *string `json:"keyPairID"`
	RootPassword *string `json:"rootPassword"`
	BandWidth    *int    `json:"bandwidth"`
}

/*
 * 创建云主机实例
 * param instanceSpec: 描述云主机配置(Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateInstancesRequest(
	instanceSpec *vm.InstanceSpec,
) *CreateInstancesRequest {

	cycleCount := 1
	cycleType := "MONTH"
	return &CreateInstancesRequest{
		CtyunRequest: core.CtyunRequest{
			URL:     "/v4/ecs/create-instance",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
		RegionID:        instanceSpec.RegionID,
		AzName:          instanceSpec.AzName,
		ClientToken:     instanceSpec.ClientToken,
		ImageID:         instanceSpec.ImageID,
		InstanceName:    instanceSpec.InstanceName,
		DisplayName:     instanceSpec.DisplayName,
		FlavorID:        instanceSpec.FlavorID,
		ImageType:       instanceSpec.ImageType,
		BootDiskType:    instanceSpec.BootDiskType,
		BootDiskSize:    instanceSpec.BootDiskSize,
		VpcID:           instanceSpec.VpcID,
		ExtIP:           instanceSpec.ExtIP,
		OnDemand:        false,
		CycleCount:      &cycleCount,
		CycleType:       &cycleType,
		NetworkCardList: instanceSpec.NetworkCardList,
		KeyPairID:       instanceSpec.KeyPairID,
		RootPassword:    instanceSpec.RootPassword,
		BandWidth:       instanceSpec.BandWidth,
	}
}

/*
	param instanceSpec: 描述云主机配置

(Required)
*/
func (r *CreateInstancesRequest) SetInstanceSpec(instanceSpec *vm.InstanceSpec) {
	r.InstanceSpec = instanceSpec
}

/*
	param clientToken: 用于保证请求的幂等性。由客户端生成，长度不能超过64个字符。

(Optional)
*/
func (r *CreateInstancesRequest) SetClientToken(clientToken string) {
	r.ClientToken = &clientToken
}

type CreateInstancesResponse struct {
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
	ReturnObj CreateInstancesResult `json:"returnObj"`
}

type CreateInstancesResult struct {
	/*订单ID。调用方在拿到masterOrderID之后，在若干错误情况下，可以使用materOrderID进一步确认订单状态及资源状态*/
	MasterOrderID string `json:"masterOrderID"`
	/*订单号*/
	MasterOrderNO string `json:"masterOrderNO"`
	/*主资源ID*/
	MasterResourceID string `json:"masterResourceID"`
	/*资源所属资源池ID*/
	RegionID string `json:"regionID"`
}
