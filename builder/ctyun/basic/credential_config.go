package basic

import (
	"fmt"
	"os"

	"github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/core"
	vm "github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/services/vm/client"
	vpc "github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/services/vpc/client"
	"github.com/hashicorp/packer/packer-plugin-sdk/template/interpolate"
)

type CTyunCredentialConfig struct {
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	RegionId  string `mapstructure:"region_id"`
	Az        string `mapstructure:"az_name"`
}

func (ct *CTyunCredentialConfig) Prepare(ctx *interpolate.Context) []error {

	errorArray := []error{}

	if ct == nil {
		return append(errorArray, fmt.Errorf("[PRE-FLIGHT] Empty CTyunCredentialConfig detected"))
	}

	if err := ct.ValidateKeyPair(); err != nil {
		errorArray = append(errorArray, err)
	}

	if err := ct.validateAz(); err != nil {
		errorArray = append(errorArray, err)
	}

	if len(errorArray) != 0 {
		return errorArray
	}

	credential := core.NewCredentials(ct.AccessKey, ct.SecretKey)
	VmClient = vm.NewVmClient(credential)
	VpcClient = vpc.NewVpcClient(credential)
	Region = ct.RegionId

	return nil
}

func (ct *CTyunCredentialConfig) ValidateKeyPair() error {

	if ct.AccessKey == "" {
		ct.AccessKey = os.Getenv("CTYUN_ACCESS_KEY")
	}

	if ct.SecretKey == "" {
		ct.SecretKey = os.Getenv("CTYUN_SECRET_KEY")
	}

	if ct.AccessKey == "" || ct.SecretKey == "" {
		return fmt.Errorf("[PRE-FLIGHT] We can't find your key pairs," +
			"write them here {access_key=xxx , secret_key=xxx} ")
	}

	return nil
}

func (ct *CTyunCredentialConfig) validateAz() error {
	if len(ct.Az) == 0 {
		return fmt.Errorf("[PRE-FLIGHT] az info missing")
	}
	return nil
}
