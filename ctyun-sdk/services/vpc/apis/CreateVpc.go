package apis

import (
	"github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/core"
)

type CreateVpcRequest struct {
	core.CtyunRequest

	/* 客户端存根，用于保证订单幂等性。要求单个云平台账户内唯一 */
	ClientToken string `json:"clientToken"`
	/* 资源池 ID	  */
	RegionID string `json:"regionID"`
	/* vpc 名称，只能由数字，字母，-组成不能以数字和-开头，最大长度 32 */
	Name string `json:"name"`
	/* VPC 的网段。建议您使用 192.168.0.0/16、172.16.0.0/12、10.0.0.0/8 三个 RFC 标准私网网段及其子网作为专有网络的主 IPv4 网段，网段掩码有效范围为 8~28 位  */
	CIDR string `json:"CIDR"`
}

/*
 * param regionId: 资源池 ID (Required)
 * param clientToken: 客户端存根 (Required)
 * param name: vpc 名称 (Required)
 * param cidr: VPC 的网段 (Required)
 */
func NewCreateVpcRequest(
	regionId string,
	clientToken string,
	name string,
	cidr string,
) *CreateVpcRequest {

	return &CreateVpcRequest{
		CtyunRequest: core.CtyunRequest{
			URL:     "/v4/vpc/create",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
		RegionID:    regionId,
		ClientToken: clientToken,
		Name:        name,
		CIDR:        cidr,
	}
}

func (r *CreateVpcRequest) SetRegionId(regionId string) {
	r.RegionID = regionId
}

func (r *CreateVpcRequest) SetClientToken(clientToken string) {
	r.ClientToken = clientToken
}

func (r *CreateVpcRequest) SetName(name string) {
	r.Name = name
}

func (r *CreateVpcRequest) SetCIDR(cidr string) {
	r.CIDR = cidr
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateVpcRequest) GetRegionId() string {
	return r.RegionID
}

type CreateVpcResponse struct {
	/*返回状态码（800为成功，900为失败）*/
	StatusCode int `json:"statusCode"`
	/*具体错误码标志*/
	ErrorCode string `json:"errorCode"`
	/*失败时的错误信息*/
	Message string `json:"message"`
	/*失败时的错误描述*/
	Description string `json:"description"`
	/*成功时返回的数据，参见returnObj对象结构	*/
	ReturnObj CreateVpcResult `json:"returnObj"`
}

type CreateVpcResult struct {
	/* vpc 示例 ID	*/
	VpcID string `json:"vpcID"`
}
