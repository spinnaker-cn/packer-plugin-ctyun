package apis

import (
	"github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/core"
	vm "github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/services/vm/models"
)

type DescribeImageRequest struct {
	core.CtyunRequest

	/* 地域ID  */
	RegionId string `json:"regionID"`

	/* 镜像ID  */
	ImageId string `json:"imageID"`
}

/*
 * param regionId: 地域ID (Required)
 * param imageId: 镜像ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeImageRequest(
	regionId string,
	imageId string,
) *DescribeImageRequest {

	return &DescribeImageRequest{
		CtyunRequest: core.CtyunRequest{
			URL:     "/v4/image/detail",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
		RegionId: regionId,
		ImageId:  imageId,
	}
}

/* param regionId: 地域ID(Required) */
func (r *DescribeImageRequest) SetRegionId(regionId string) {
	r.RegionId = regionId
}

/* param imageId: 镜像ID(Required) */
func (r *DescribeImageRequest) SetImageId(imageId string) {
	r.ImageId = imageId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeImageRequest) GetRegionId() string {
	return r.RegionId
}

type DescribeImageResponse struct {
	/*返回状态码（800为成功，900为失败）*/
	StatusCode int `json:"statusCode"`
	/*具体错误码标志*/
	ErrorCode string `json:"errorCode"`
	/*失败时的错误信息*/
	Message string `json:"message"`
	/*失败时的错误描述*/
	Description string `json:"description"`
	/*成功时返回的数据，参见returnObj对象结构	*/
	ReturnObj DescribeImageResult `json:"returnObj"`
}

type DescribeImageResult struct {
	Image []vm.Image `json:"images"`
}
