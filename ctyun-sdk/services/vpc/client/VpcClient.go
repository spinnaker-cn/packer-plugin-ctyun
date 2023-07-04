package client

import (
	"encoding/json"
	"errors"
	"github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/core"
	vpc "github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/services/vpc/apis"
)

type VpcClient struct {
	core.CtyunClient
}

func NewVpcClient(credential *core.Credential) *VpcClient {
	if credential == nil {
		return nil
	}

	config := core.NewConfig()
	config.SetEndpoint("ctvpc-global.ctapi.ctyun.cn")

	return &VpcClient{
		core.CtyunClient{
			Credential:  *credential,
			Config:      *config,
			ServiceName: "vpc",
			Revision:    "0.5.1",
			Logger:      core.NewDefaultLogger(core.LogInfo),
		}}
}

func (c *VpcClient) SetConfig(config *core.Config) {
	c.Config = *config
}

func (c *VpcClient) SetLogger(logger core.Logger) {
	c.Logger = logger
}

/* 查询Vpc信息详情 */
func (c *VpcClient) DescribeVpc(request *vpc.DescribeVpcRequest) (*vpc.DescribeVpcResponse, error) {

	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request)
	if err != nil {
		return nil, err
	}

	ctResp := &vpc.DescribeVpcResponse{}
	err = json.Unmarshal(resp, ctResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return ctResp, err
}

/* 创建VPC私有网络 */

func (c *VpcClient) CreateVpc(request *vpc.CreateVpcRequest) (*vpc.CreateVpcResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request)
	if err != nil {
		c.Logger.Log(core.LogError, "Create Vpc failed, resp: %s", string(resp))
		return nil, err
	}
	var ctResp vpc.CreateVpcResponse
	err = json.Unmarshal(resp, &ctResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return &ctResp, err
}
