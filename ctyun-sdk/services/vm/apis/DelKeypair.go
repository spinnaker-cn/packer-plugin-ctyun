package apis

import (
	"github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/core"
)

type DelKeypairRequest struct {
	core.CtyunRequest

	/* 地域ID  */
	RegionId string `json:"regionID"`

	/* 密钥对名称，需要全局唯一。只允许数字、大小写字母、下划线“_”及中划线“-”，不超过32个字符。
	 */
	KeyName string `json:"keyPairName"`
}

/*
 * param regionId: 地域ID (Required)
 * param keyName: 密钥对名称。只能由数字、字母、-组成,不能以数字和-开头、以-结尾,且长度为2-63字符。(Required)
 *
 */
func DeleteKeypairRequest(
	regionId string,
	keyName string,
) *DelKeypairRequest {

	return &DelKeypairRequest{
		CtyunRequest: core.CtyunRequest{
			URL:     "/v4/ecs/keypair/delete",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
		RegionId: regionId,
		KeyName:  keyName,
	}
}

func (r *DelKeypairRequest) SetRegionId(regionId string) {
	r.RegionId = regionId
}

func (r *DelKeypairRequest) SetKeyName(keyName string) {
	r.KeyName = keyName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DelKeypairRequest) GetRegionId() string {
	return r.RegionId
}

type DelKeypairResponse struct {
	/*返回状态码（800为成功，900为失败）*/
	StatusCode int `json:"statusCode"`
	/*具体错误码标志*/
	ErrorCode string `json:"errorCode"`
	/*失败时的错误信息*/
	Message string `json:"message"`
	/*失败时的错误描述*/
	Description string `json:"description"`
	/*成功时返回的数据，参见returnObj对象结构	*/
	ReturnObj DelKeypairResult `json:"returnObj"`
}

type DelKeypairResult struct{}
