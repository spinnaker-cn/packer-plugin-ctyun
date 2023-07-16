package basic

import (
	"context"
	"fmt"
	"github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/services/vm/apis"
	"github.com/hashicorp/packer/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer/packer-plugin-sdk/packer"
	"io/ioutil"
	"time"
)

type stepConfigCredentials struct {
	InstanceSpecConfig *CTyunInstanceSpecConfig
	ui                 packersdk.Ui
}

func (s *stepConfigCredentials) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {

	s.ui = state.Get("ui").(packersdk.Ui)
	privateKeyPath := s.InstanceSpecConfig.Comm.SSHPrivateKeyFile
	//privateKeyName := s.InstanceSpecConfig.Comm.SSHKeyPairName
	//newKeyName := s.InstanceSpecConfig.Comm.SSHTemporaryKeyPairName

	//如果有私钥文件则使用私钥登录
	if len(privateKeyPath) > 0 {
		s.ui.Message("\t Private key detected, we are going to login with this private key :)")
		return s.ReadExistingPair(state)
	} else {
		privateKeyName := "packer-" + time.Now().Format("20060102150405")
		s.InstanceSpecConfig.Comm.SSHKeyPairName = privateKeyName
		s.ui.Message("\t We are going to create a new key pair with its name=" + privateKeyName)
		return s.CreateRandomKeyPair(privateKeyName, state)
	}

	//使用用户名密码登录
	//if len(password) > 0 {
	//	s.ui.Message("\t Password detected, we are going to login with this password :)")
	//	return multistep.ActionContinue
	//}

	s.ui.Error("[ERROR] Didn't detect any credentials, you have to specify either " +
		"{password} or " +
		"{key_name+local_private_key_path} or " +
		"{temporary_key_pair_name} cheers :)")
	return multistep.ActionHalt
}

func (s *stepConfigCredentials) ReadExistingPair(state multistep.StateBag) multistep.StepAction {
	privateKeyBytes, err := ioutil.ReadFile(s.InstanceSpecConfig.Comm.SSHPrivateKeyFile)
	if err != nil {
		error := fmt.Errorf("Error Read KeyPair File")
		state.Put("error", error)
		s.ui.Error("Cannot read local private-key, were they correctly placed? Here's the error" + err.Error())
		return multistep.ActionHalt
	}
	s.ui.Message("\t\t Keys read successfully :)")
	s.InstanceSpecConfig.Comm.SSHPrivateKey = privateKeyBytes
	return multistep.ActionContinue
}

func (s *stepConfigCredentials) CreateRandomKeyPair(keyName string, state multistep.StateBag) multistep.StepAction {
	req := apis.NewCreateKeypairRequest(Region, keyName)
	resp, err := VmClient.CreateKeypair(req)
	if err != nil || resp.StatusCode != 800 {
		error := fmt.Errorf("Error creating KeyPair")
		state.Put("error", error)
		s.ui.Error(fmt.Sprintf("[ERROR] Cannot create a new key pair for you, \n error=%v \n response=%v", err, resp))
		return multistep.ActionHalt
	}
	s.ui.Message("\t\t Keys created successfully :)")
	s.InstanceSpecConfig.Comm.SSHPrivateKey = []byte(resp.ReturnObj.PrivateKey)
	s.InstanceSpecConfig.KeyPairID = resp.ReturnObj.KeyPairID
	return multistep.ActionContinue
}

func (s *stepConfigCredentials) Cleanup(state multistep.StateBag) {}
