package basic

import (
	"context"
	"fmt"
	"github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/core"
	apis "github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/services/vm/apis"
	vm "github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/services/vm/models"
	"github.com/hashicorp/packer/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer/packer-plugin-sdk/packer"
	"regexp"
	"time"
)

type stepCreateCTyunInstance struct {
	InstanceSpecConfig *CTyunInstanceSpecConfig
	CredentialConfig   *CTyunCredentialConfig
	ui                 packersdk.Ui
}

func (s *stepCreateCTyunInstance) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {

	s.ui = state.Get("ui").(packersdk.Ui)
	s.ui.Say("Creating instances")

	instanceSpec := vm.InstanceSpec{
		AzName:          &s.CredentialConfig.Az,
		RegionID:        &s.CredentialConfig.RegionId,
		ClientToken:     &s.InstanceSpecConfig.ClientToken,
		ImageID:         &s.InstanceSpecConfig.ImageID,
		InstanceName:    &s.InstanceSpecConfig.InstanceName,
		DisplayName:     &s.InstanceSpecConfig.DisplayName,
		FlavorID:        &s.InstanceSpecConfig.FlavorID,
		ImageType:       &s.InstanceSpecConfig.ImageType,
		BootDiskType:    &s.InstanceSpecConfig.BootDiskType,
		BootDiskSize:    &s.InstanceSpecConfig.BootDiskSize,
		VpcID:           &s.InstanceSpecConfig.VpcID,
		ExtIP:           &s.InstanceSpecConfig.ExtIP,
		OnDemand:        &s.InstanceSpecConfig.OnDemand,
		CycleCount:      &s.InstanceSpecConfig.CycleCount,
		CycleType:       &s.InstanceSpecConfig.CycleType,
		NetworkCardList: s.InstanceSpecConfig.NetworkCardList,
		KeyPairID:       &s.InstanceSpecConfig.KeyPairID,
		RootPassword:    &s.InstanceSpecConfig.RootPassword,
		BandWidth:       &s.InstanceSpecConfig.BandWidth,
		SecGroupList:    &s.InstanceSpecConfig.SecGroupList,
	}

	req := apis.NewCreateInstancesRequest(&instanceSpec)
	resp, err := VmClient.CreateInstances(req)

	if err != nil {
		err := fmt.Errorf("Error creating instance, error-%v ", err)
		state.Put("error", err)
		s.ui.Error(err.Error())
		return multistep.ActionHalt
	}
	if resp.StatusCode != 800 {
		error := fmt.Errorf("Error Creating Instance Code Is Not 800")
		state.Put("error", error)
		s.ui.Error(fmt.Sprintf(
			"created instance error, its errorCode=%v , "+
				"its message=%v, :) ", resp.ErrorCode, resp.Message))
		return multistep.ActionHalt
	}
	insListResponse, err := WaitForExpected(resp.ReturnObj.MasterResourceID)
	if err != nil {
		error := fmt.Errorf("Waiting For Instance Id Error")
		state.Put("error", error)
		s.ui.Error("query Instance Id error: " + err.Error())
		return multistep.ActionHalt
	}
	s.InstanceSpecConfig.InstanceId = insListResponse.ReturnObj.Results[0].InstanceID
	//TODO
	//s.InstanceSpecConfig.InstanceId = "0bd9df10-f7b1-c896-4b1e-32a3cb6c51c8"
	//等待实例状态
	instanceInterface, err := InstanceStatusRefresher(s.InstanceSpecConfig.InstanceId, []string{VM_CREATING, VM_STARTING, VM_MASTER_ORDER_CREATING}, []string{VM_RUNNING})

	if err != nil {
		error := fmt.Errorf("Waiting For Instance Status Error")
		state.Put("error", error)
		s.ui.Error("Waiting For Instance Status Error: " + err.Error())
		return multistep.ActionHalt
	}
	//返回云主机实例详情
	instance := instanceInterface.(vm.Instance)
	if s.InstanceSpecConfig.ExtIP == "1" {
		//使用外网IP
		state.Put("floatingIP", instance.FloatingIP)
	} else {
		//使用内网IP
		state.Put("floatingIP", instance.PrivateIP)
	}
	s.ui.Message(fmt.Sprintf(
		"create the instance success, its name=%v , "+
			"its id=%v: ", instance.InstanceName, s.InstanceSpecConfig.InstanceId))

	return multistep.ActionContinue
}

/*
 * 查询创建云主机实例ID
 */
func queryInstanceId(resourceID string) (*apis.QueryInstancesResponse, error) {

	req := apis.NewQueryInstancesRequest(&Region, &resourceID)
	resp, err := VmClient.QueryInstancesList(req)
	if err != nil || resp.StatusCode != 800 {
		return nil, err
	}
	fmt.Println(resp.StatusCode)
	return resp, nil
}

/*
 * 查询云主机详情
 */
func describeInstance(instanceId string) (*apis.DescribeInstanceResponse, error) {

	req := apis.NewDescribeInstanceRequest(Region, instanceId)
	resp, err := VmClient.DescribeInstance(req)
	if err != nil || resp.StatusCode != 800 {
		return nil, err
	}
	fmt.Println(resp.StatusCode)
	return resp, nil
}

/*
 * 创建云主机实例后刷新云主机状态
 */
func InstanceStatusRefresher(id string, pending, target []string) (instance interface{}, err error) {

	stateConf := &StateChangeConf{
		Pending:    pending,
		Target:     target,
		Refresh:    instanceStatusRefresher(id),
		Delay:      3 * time.Second,
		Timeout:    10 * time.Minute,
		MinTimeout: 1 * time.Second,
	}
	if instance, err = stateConf.WaitForState(); err != nil {
		return nil, fmt.Errorf("[ERROR] Failed in creating instance ,err message:%v", err)
	}
	return instance, nil
}

func instanceStatusRefresher(instanceId string) StateRefreshFunc {

	return func() (instance interface{}, status string, err error) {

		err = Retry(time.Minute, func() *RetryError {

			req := apis.NewDescribeInstanceRequest(Region, instanceId)
			resp, err := VmClient.DescribeInstance(req)

			if err == nil && resp.StatusCode == 800 {
				instance = resp.ReturnObj
				status = resp.ReturnObj.InstanceStatus
				return nil
			}

			instance = nil
			status = ""
			if connectionError(err) {
				return RetryableError(err)
			} else {
				return NonRetryableError(err)
			}
		})
		return instance, status, err
	}
}

/*
 *API创建云主机成功后，等待云主机创建完成
 */
func WaitForExpected(masterResourceID string) (*apis.QueryInstancesResponse, error) {

	args := WaitForExpectArgs{
		RetryInterval: 10 * time.Second,
		RetryTimeout:  10 * time.Minute,
	}

	if args.RetryInterval <= 0 {
		args.RetryInterval = defaultRetryInterval
	}
	if args.RetryTimes <= 0 {
		args.RetryTimes = defaultRetryTimes
	}

	var timeoutPoint time.Time
	if args.RetryTimeout > 0 {
		timeoutPoint = time.Now().Add(args.RetryTimeout)
	}

	var lastResponse *apis.QueryInstancesResponse
	var lastError error

	for i := 0; ; i++ {
		if args.RetryTimeout > 0 && time.Now().After(timeoutPoint) {
			break
		}

		if args.RetryTimeout <= 0 && i >= args.RetryTimes {
			break
		}

		//查询云主机列表，根据云主机列表查询云主机实例ID
		insListResponse, err := queryInstanceId(masterResourceID)
		//TODO 测试
		//insListResponse, err := queryInstanceId(instanceSpec, "4f371c495346483398efdedbe281d2c9")
		lastResponse = insListResponse
		lastError = err

		if len(lastResponse.ReturnObj.Results) > 0 &&
			lastResponse.ReturnObj.Results[0].InstanceStatus == "running" {

			return insListResponse, nil
		}
		time.Sleep(args.RetryInterval)
		VmClient.Logger.Log(core.LogInfo, "[TRACE] Waiting list-instances 10s before next try")
	}

	if lastError == nil {
		lastError = fmt.Errorf("<no error>")
	}

	if args.RetryTimeout > 0 {
		return lastResponse, fmt.Errorf("evaluate failed after %d Minutes timeout with %d seconds retry interval: %s", int(args.RetryTimeout.Minutes()), int(args.RetryInterval.Seconds()), lastError)
	}

	return lastResponse, fmt.Errorf("evaluate failed after %d times retry with %d seconds retry interval: %s", args.RetryTimes, int(args.RetryInterval.Seconds()), lastError)
}

func connectionError(e error) bool {

	if e == nil {
		return false
	}
	ok, _ := regexp.MatchString(CONNECT_FAILED, e.Error())
	return ok
}
func instanceHost(state multistep.StateBag) (string, error) {
	return state.Get("floatingIP").(string), nil
}
func (s *stepCreateCTyunInstance) Cleanup(state multistep.StateBag) {

	ui := state.Get("ui").(packersdk.Ui)
	if state.Get("error") != nil {

		req := apis.DeleteKeypairRequest(Region, s.InstanceSpecConfig.Comm.SSHKeyPairName)
		resp, err := VmClient.DelKeypair(req)
		if err != nil || resp.StatusCode != 800 {
			ui.Error(fmt.Sprintf("[ERROR] Delete KeyPair On Instance Error-%v ,Resp:%v", err, resp))
		} else {
			ui.Message("Delete KeyPair On Instance Success")
		}
		if s.InstanceSpecConfig.InstanceId != "" {
			reqStop := apis.NewStopInstanceRequest(Region, s.InstanceSpecConfig.InstanceId)
			respStop, errStop := VmClient.StopInstance(reqStop)
			if errStop != nil || respStop.StatusCode != 800 {
				if errStop != nil || respStop.StatusCode != 800 {
					ui.Error(fmt.Sprintf("[ERROR] Delete Instance On Instance Stop Error-%v ,Resp:%v", errStop, respStop))
				}
			} else {
				_, err = InstanceStatusRefresher(s.InstanceSpecConfig.InstanceId, []string{VM_RUNNING, VM_STOPPING}, []string{VM_STOPPED})
				if err != nil {
					error := fmt.Errorf("Waiting For Stop Instance Status Error")
					ui.Error(error.Error())
				}

				reqDel := apis.NewDelInstanceRequest(Region, s.InstanceSpecConfig.InstanceId, s.InstanceSpecConfig.ClientToken)
				respDel, errDel := VmClient.DelInstance(reqDel)

				if errDel != nil || respDel.StatusCode != 800 {
					ui.Error(fmt.Sprintf("[ERROR] Delete Instance On Instance Error-%v ,Resp:%v", errDel, respDel))
				} else {
					ui.Message("Delete Instance On Instance Success")
				}
			}

		}
	}

}
