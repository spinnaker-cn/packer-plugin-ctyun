package apis

import (
	"github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/core"
)

type DescribeVpcRequest struct {
	core.CtyunRequest

	/* 客户端存根，用于保证订单幂等性。要求单个云平台账户内唯一  (Optional) */
	ClientToken string `json:"clientToken"`
	/* 资源池 ID	  (Required)*/
	RegionID string `json:"regionID"`
	/* 可用区名称  (Optional) */
	AzName string `json:"azName"`
	/* 企业项目 ID，默认为"0"   (Optional) */
	ProjectID string `json:"projectID"`
	/* VPC 的 ID  (Required)*/
	VpcID string `json:"vpcID"`
}

/*
 * param regionId: Region ID (Required)
 * param vpcId: Vpc ID (Required)
 */
func NewDescribeVpcRequest(
	regionId string,
	vpcId string,
) *DescribeVpcRequest {

	return &DescribeVpcRequest{
		CtyunRequest: core.CtyunRequest{
			URL:     "/v4/vpc/query",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
		RegionID: regionId,
		VpcID:    vpcId,
	}
}

/* param regionId: Region ID(Required) */
func (r *DescribeVpcRequest) SetRegionId(regionId string) {
	r.RegionID = regionId
}

/* param vpcId: Vpc ID(Required) */
func (r *DescribeVpcRequest) SetVpcId(vpcId string) {
	r.VpcID = vpcId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeVpcRequest) GetRegionId() string {
	return r.RegionID
}

type DescribeVpcResponse struct {
	/*返回状态码（800为成功，900为失败）*/
	StatusCode int `json:"statusCode"`
	/*具体错误码标志*/
	ErrorCode string `json:"errorCode"`
	/*失败时的错误信息*/
	Message string `json:"message"`
	/*失败时的错误描述*/
	Description string `json:"description"`
	/*成功时返回的数据，参见returnObj对象结构	*/
	ReturnObj DescribeVpcResult `json:"returnObj"`
}

type DescribeVpcResult struct {
	/* vpcID	*/
	VpcID string `json:"vpcID"`
	/* 名称 */
	Name string `json:"name"`
	/* 描述 */
	Description string `json:"description"`
	/* 子网 */
	CIDR string `json:"CIDR"`
	/* 是否开启 ipv6 */
	Ipv6Enabled bool `json:"ipv6Enabled"`
	/* ipv6 子网列表 */
	Ipv6CIDRs []string `json:"ipv6CIDRs"`
	/* 子网 id 列表 */
	SubnetIDs []string `json:"subnetIDs"`
	/* 网关 id 列表 */
	NatGatewayIDs []string `json:"natGatewayIDs"`
	/* 附加网段 */
	SecondaryCIDRS []string `json:"secondaryCIDRS"`
}
