package apis

import (
	"github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/core"
	vm "github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/services/vm/models"
)

type CreateImageRequest struct {
	core.CtyunRequest

	ImageSpec vm.ImageSpec `json:"imageSpec"`
	/* 地域ID  */
	RegionId string `json:"regionID"`

	/* 云主机ID  */
	InstanceId string `json:"instanceID"`

	/* 镜像名称 */
	ImageName string `json:"imageName"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: 云主机ID (Required)
 * param imageName: 镜像名称(Required)
 */
func NewCreateImageRequest(
	imageSpec vm.ImageSpec,
) *CreateImageRequest {

	return &CreateImageRequest{
		CtyunRequest: core.CtyunRequest{
			URL:     "/v4/image/create",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
		RegionId:   imageSpec.RegionId,
		InstanceId: imageSpec.InstanceId,
		ImageName:  imageSpec.ImageName,
	}
}

func (r *CreateImageRequest) SetImageSpec(imageSpec vm.ImageSpec) {
	r.ImageSpec = imageSpec
}

func (r *CreateImageRequest) SetInstanceId(instanceId string) {
	r.InstanceId = instanceId
}

func (r *CreateImageRequest) SetName(name string) {
	r.ImageName = name
}

func (r *CreateImageRequest) setRegionId(regionId string) {
	r.RegionId = regionId
}

type CreateImageResponse struct {
	/*返回状态码（800为成功，900为失败）*/
	StatusCode int `json:"statusCode"`
	/*具体错误码标志*/
	ErrorCode string `json:"errorCode"`
	/*失败时的错误信息*/
	Message string `json:"message"`
	/*失败时的错误描述*/
	Description string `json:"description"`
	/*成功时返回的数据，参见returnObj对象结构	*/
	ReturnObj CreateImageResult `json:"returnObj"`
}

type CreateImageResult struct {
	TotalCount string     `json:"totalCount"`
	Images     []vm.Image `json:"images"`
}
