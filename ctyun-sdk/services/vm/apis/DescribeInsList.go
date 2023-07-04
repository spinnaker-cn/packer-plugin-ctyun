package apis

import (
	"github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/core"
	vm "github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/services/vm/models"
)

/*
 * 查询云主机列表api
 */
type QueryInstancesRequest struct {
	core.CtyunRequest
	/* 资源池ID。您可以调用资源池列表查询获取最新的资源池列表可查询：https://www.ctyun.cn/document/10026730/10040588。 (true) */
	RegionID *string `json:"regionID"`

	ResourceID *string `json:"resourceID"`
}

/*
 * 查询云主机列表
 * param instanceListSpec: 描述云主机配置(Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryInstancesRequest(
	regionId *string,
	resourceID *string,
) *QueryInstancesRequest {

	return &QueryInstancesRequest{
		CtyunRequest: core.CtyunRequest{
			URL:     "/v4/ecs/list-instances",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
		RegionID:   regionId,
		ResourceID: resourceID,
	}
}

func (r *QueryInstancesRequest) SetRegionId(regionID string) {
	r.RegionID = &regionID
}
func (r *QueryInstancesRequest) SetResourceId(resourceID string) {
	r.ResourceID = &resourceID
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
	Results []vm.Instance `json:"results"`
}
