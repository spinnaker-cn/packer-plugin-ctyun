package apis

import (
	"github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/core"
	vm "github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/services/vm/models"
)

type DescribeInstanceRequest struct {
	core.CtyunRequest

	/* 地域ID  */
	RegionId string `json:"regionID"`

	/* 云主机ID  */
	InstanceId string `json:"instanceID"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: 云主机ID (Required)
 */
func NewDescribeInstanceRequest(
	regionId string,
	instanceId string,
) *DescribeInstanceRequest {

	return &DescribeInstanceRequest{
		CtyunRequest: core.CtyunRequest{
			URL:     "/v4/ecs/instance-details",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
		RegionId:   regionId,
		InstanceId: instanceId,
	}
}

/* param regionId: 地域ID(Required) */
func (r *DescribeInstanceRequest) SetRegionId(regionId string) {
	r.RegionId = regionId
}

/* param instanceId: 云主机ID(Required) */
func (r *DescribeInstanceRequest) SetInstanceId(instanceId string) {
	r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInstanceRequest) GetRegionId() string {
	return r.RegionId
}

type DescribeInstanceResponse struct {
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
	ReturnObj vm.Instance `json:"returnObj"`
}
