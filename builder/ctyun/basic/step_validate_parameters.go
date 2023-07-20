package basic

import (
	"context"
	"fmt"
	vm "github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/services/vm/apis"
	vpc "github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/services/vpc/apis"
	"github.com/hashicorp/packer/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer/packer-plugin-sdk/packer"
)

type stepValidateParameters struct {
	InstanceSpecConfig *CTyunInstanceSpecConfig
	ui                 packersdk.Ui
	state              multistep.StateBag
}

func (s *stepValidateParameters) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {

	s.ui = state.Get("ui").(packersdk.Ui)
	s.state = state
	s.ui.Say("Validating parameters...")

	if err := s.ValidateVpcFunc(state); err != nil {
		s.ui.Error(err.Error())
		return multistep.ActionHalt
	}

	if err := s.ValidateImageFunc(state); err != nil {
		s.ui.Error(err.Error())
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (s *stepValidateParameters) ValidateVpcFunc(state multistep.StateBag) error {

	vpcId := s.InstanceSpecConfig.VpcID
	//if len(vpcId) == 0 {
	//	s.ui.Message("\t 'vpc' is not specified, we will create a new one for you :) ")
	//	return s.CreateRandomVpc()
	//}

	s.ui.Message("\t validating your vpc:" + s.InstanceSpecConfig.VpcID)
	req := vpc.NewDescribeVpcRequest(Region, vpcId)
	resp, err := VpcClient.DescribeVpc(req)
	if err != nil {
		error := fmt.Errorf("Validate Vpc Status Error")
		state.Put("error", error)
		return fmt.Errorf("[ERROR] Failed in validating vpc->%s, reasons:%v", vpcId, err)
	}
	if resp != nil && resp.StatusCode != 800 {
		error := fmt.Errorf("Validate Vpc Status Is Not 800")
		state.Put("error", error)
		return fmt.Errorf("[ERROR] Something wrong with your vpc->%s, reasons:%v", vpcId, resp.ErrorCode)
	}

	s.ui.Message("\t vpc found:" + resp.ReturnObj.Name)
	return nil

}

func (s *stepValidateParameters) ValidateImageFunc(state multistep.StateBag) error {

	s.ui.Message("\t validating your base image:" + s.InstanceSpecConfig.ImageID)
	imageId := s.InstanceSpecConfig.ImageID
	req := vm.NewDescribeImageRequest(s.InstanceSpecConfig.RegionID, imageId)
	resp, err := VmClient.DescribeImage(req)
	if err != nil {
		error := fmt.Errorf("Validate Image Status Error")
		state.Put("error", error)
		return fmt.Errorf("[ERROR] Failed in validating your image->%s, reasons:%v", imageId, err)
	}
	if resp != nil && resp.StatusCode != 800 {
		error := fmt.Errorf("Validate Image Status Is Not 800")
		state.Put("error", error)
		return fmt.Errorf("[ERROR] Something wrong with your image->%s, reasons:%v", imageId, resp.ErrorCode)
	}

	s.ui.Message("\t image found:" + resp.ReturnObj.Image[0].ImageName)
	s.state.Put("source_image", &resp.ReturnObj.Image[0].ImageName)
	return nil
}

func (s *stepValidateParameters) CreateRandomVpc() error {

	resp, err := s.CreateVpc()
	if err != nil || resp.StatusCode != 800 {
		errorMessage := fmt.Sprintf("[ERROR] Failed in creating new vpc :( \n error:%v \n response:%v", err, resp)
		s.ui.Error(errorMessage)
		return fmt.Errorf(errorMessage)
	}

	s.InstanceSpecConfig.VpcID = resp.ReturnObj.VpcID
	s.ui.Message("\t\t Hi, we have created a new vpc for you :) its name is 'created_by_packer' and its id=" + resp.ReturnObj.VpcID)
	return nil
}

/*
 * 创建VPC
 */
func (s *stepValidateParameters) CreateVpc() (*vpc.CreateVpcResponse, error) {

	req := vpc.NewCreateVpcRequest(s.InstanceSpecConfig.RegionID, s.InstanceSpecConfig.ClientToken, "create-packer-vpc", "192.168.0.0/16")
	resp, err := VpcClient.CreateVpc(req)

	return resp, err
}

func (s *stepValidateParameters) Cleanup(state multistep.StateBag) {}
