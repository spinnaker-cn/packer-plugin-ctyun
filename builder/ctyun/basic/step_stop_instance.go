package basic

import (
	"context"
	"fmt"
	"github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/services/vm/apis"
	"github.com/hashicorp/packer/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer/packer-plugin-sdk/packer"
)

/*
 * 停止主机实例
 */
type stepStopCTCloudInstance struct {
	InstanceSpecConfig *CTyunInstanceSpecConfig
}

func (s *stepStopCTCloudInstance) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {

	ui := state.Get("ui").(packersdk.Ui)
	ui.Say("Stopping this instance")

	req := apis.NewStopInstanceRequest(s.InstanceSpecConfig.RegionID, s.InstanceSpecConfig.InstanceId)
	resp, err := VmClient.StopInstance(req)
	if err != nil || resp.StatusCode != 800 {
		ui.Error(fmt.Sprintf("[ERROR] Failed in trying to stop this basic: Error-%v ,Resp:%v", err, resp))
		return multistep.ActionHalt
	}

	_, err = InstanceStatusRefresher(s.InstanceSpecConfig.InstanceId, []string{VM_RUNNING, VM_STOPPING}, []string{VM_STOPPED})

	if err != nil {
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	ui.Message("Instance has been stopped :)")
	return multistep.ActionContinue
}
func (s *stepStopCTCloudInstance) Cleanup(state multistep.StateBag) {

}
