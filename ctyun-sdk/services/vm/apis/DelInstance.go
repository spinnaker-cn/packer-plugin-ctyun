package apis

import (
	"github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/core"
)

type DelInstanceRequest struct {
	core.CtyunRequest

	/* 地域ID  */
	RegionId string `json:"regionID"`

	/* 云主机ID  */
	InstanceId string `json:"instanceID"`

	ClientToken string `json:"clientToken"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: 云主机ID (Required)
 * param clientToken: clientToken (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDelInstanceRequest(
	regionId string,
	instanceId string,
	clientToken string,
) *DelInstanceRequest {

	return &DelInstanceRequest{
		CtyunRequest: core.CtyunRequest{
			URL:     "/v4/ecs/unsubscribe-instance",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
		RegionId:    regionId,
		InstanceId:  instanceId,
		ClientToken: clientToken,
	}
}

/* param regionId: 地域ID(Required) */
func (r *DelInstanceRequest) SetRegionId(regionId string) {
	r.RegionId = regionId
}

/* param instanceId: 云主机ID(Required) */
func (r *DelInstanceRequest) SetInstanceId(instanceId string) {
	r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DelInstanceRequest) GetRegionId() string {
	return r.RegionId
}

type DelInstanceResponse struct {
	/*返回状态码（800为成功，900为失败）*/
	StatusCode int `json:"statusCode"`
	/*具体错误码标志*/
	ErrorCode string `json:"errorCode"`
	/*失败时的错误信息*/
	Message string `json:"message"`
	/*失败时的错误描述*/
	Description string `json:"description"`
	/*成功时返回的数据，参见returnObj对象结构	*/
	ReturnObj DelInstancesResult `json:"result"`
}
type DelInstancesResult struct {
	/*主订单ID*/
	MasterOrderID string `json:"masterOrderID"`
	/*订单号*/
	MasterOrderNO string `json:"masterOrderNO"`
	/*资源池ID*/
	RegionID string `json:"regionID"`
}
