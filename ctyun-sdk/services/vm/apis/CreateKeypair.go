package apis

import (
	"github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/core"
)

type CreateKeypairRequest struct {
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
func NewCreateKeypairRequest(
	regionId string,
	keyName string,
) *CreateKeypairRequest {

	return &CreateKeypairRequest{
		CtyunRequest: core.CtyunRequest{
			URL:     "/v4/ecs/keypair/create-keypair",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
		RegionId: regionId,
		KeyName:  keyName,
	}
}

func (r *CreateKeypairRequest) SetRegionId(regionId string) {
	r.RegionId = regionId
}

func (r *CreateKeypairRequest) SetKeyName(keyName string) {
	r.KeyName = keyName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateKeypairRequest) GetRegionId() string {
	return r.RegionId
}

type CreateKeypairResponse struct {
	/*返回状态码（800为成功，900为失败）*/
	StatusCode int `json:"statusCode"`
	/*具体错误码标志*/
	ErrorCode string `json:"errorCode"`
	/*失败时的错误信息*/
	Message string `json:"message"`
	/*失败时的错误描述*/
	Description string `json:"description"`
	/*成功时返回的数据，参见returnObj对象结构	*/
	ReturnObj CreateKeypairResult `json:"returnObj"`
}

type CreateKeypairResult struct {
	/*密钥对的公钥*/
	PublicKey string `json:"publicKey"`

	/*密钥对的私钥*/
	PrivateKey string `json:"privateKey"`

	/*密钥对名称*/
	KeyPairName string `json:"keyPairName"`

	/*密钥对的指纹，采用MD5信息摘要算法*/
	FingerPrint string `json:"privateKey"`

	/*密钥对的ID*/
	KeyPairID string `json:"keyPairName"`
}
