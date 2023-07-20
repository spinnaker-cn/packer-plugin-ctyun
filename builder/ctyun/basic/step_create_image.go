package basic

import (
	"context"
	"fmt"
	"github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/services/vm/apis"
	vm "github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/services/vm/models"
	"github.com/hashicorp/packer/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer/packer-plugin-sdk/packer"
	"time"
)

type stepCreateCTyunImage struct {
	InstanceSpecConfig *CTyunInstanceSpecConfig
}

func (s *stepCreateCTyunImage) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {

	ui := state.Get("ui").(packersdk.Ui)
	ui.Say("Creating images")

	imageSpec := vm.ImageSpec{
		RegionId:   s.InstanceSpecConfig.RegionID,
		InstanceId: s.InstanceSpecConfig.InstanceId,
		ImageName:  s.InstanceSpecConfig.InstanceName,
	}

	req := apis.NewCreateImageRequest(imageSpec)
	resp, err := VmClient.CreateImage(req)
	if err != nil || resp.StatusCode != 800 {
		error := fmt.Errorf("Creating image: Error")
		state.Put("error", error)
		ui.Error(fmt.Sprintf("[ERROR] Creating image: Error-%v ,Resp:%v", err, resp))

		return multistep.ActionHalt
	}
	s.InstanceSpecConfig.ArtifactId = resp.ReturnObj.Images[0].ImageID
	s.InstanceSpecConfig.ArtifactName = resp.ReturnObj.Images[0].ImageName
	_, err = ImageStatusWaiter(resp.ReturnObj.Images[0].ImageID, []string{IM_QUEUED}, []string{READY})

	if err != nil {
		error := fmt.Errorf("Waiting For Image Status Error")
		state.Put("error", error)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	ui.Message("Image has been created :)")
	return multistep.ActionContinue
}

//func ImageStatusWaiter(imageId string) error {
//	req := apis.NewDescribeImageRequest(Region, imageId)
//
//	return Retry(5*time.Minute, func() *RetryError {
//		resp, err := VmClient.DescribeImage(req)
//		if err == nil && resp.ReturnObj.Image[0].Status == READY {
//			return nil
//		}
//		if connectionError(err) {
//			return RetryableError(err)
//		} else {
//			return NonRetryableError(err)
//		}
//	})
//
//}

/*
 * 创建私有镜像后刷新镜像状态
 */
func ImageStatusWaiter(id string, pending, target []string) (instance interface{}, err error) {

	stateConf := &StateChangeConf{
		Pending:    pending,
		Target:     target,
		Refresh:    imageStatusRefresher(id),
		Delay:      3 * time.Second,
		Timeout:    30 * time.Minute,
		MinTimeout: 1 * time.Second,
	}
	if instance, err = stateConf.WaitForState(); err != nil {
		return nil, fmt.Errorf("[ERROR] Failed in creating image ,err message:%v", err)
	}
	return instance, nil
}

func imageStatusRefresher(imageId string) StateRefreshFunc {

	return func() (image interface{}, status string, err error) {

		err = Retry(time.Minute, func() *RetryError {
			req := apis.NewDescribeImageRequest(Region, imageId)
			resp, err := VmClient.DescribeImage(req)

			if err == nil && resp.StatusCode == 800 {
				image = resp.ReturnObj
				status = resp.ReturnObj.Image[0].Status
				return nil
			}

			image = nil
			status = ""
			if connectionError(err) {
				return RetryableError(err)
			} else {
				return NonRetryableError(err)
			}
		})
		return image, status, err
	}
}

func (s *stepCreateCTyunImage) Cleanup(state multistep.StateBag) {

	ui := state.Get("ui").(packersdk.Ui)
	req := apis.DeleteKeypairRequest(Region, s.InstanceSpecConfig.Comm.SSHKeyPairName)
	resp, err := VmClient.DelKeypair(req)

	if err != nil || resp.StatusCode != 800 {
		ui.Error(fmt.Sprintf("[ERROR] Delete KeyPair On Image Error-%v ,Resp:%v", err, resp))
	} else {
		ui.Message("Delete KeyPair On Image Success")
	}

}
