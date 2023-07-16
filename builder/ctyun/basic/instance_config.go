package basic

import (
	"fmt"
	vm "github.com/ctyun/packer-plugin-ctyun/ctyun-sdk/services/vm/models"
	"github.com/hashicorp/packer/packer-plugin-sdk/communicator"
	"github.com/hashicorp/packer/packer-plugin-sdk/template/interpolate"
)

type CTyunInstanceSpecConfig struct {
	/* 客户端存根，用于保证订单幂等性。要求单个云平台账户内唯一。 (true) */
	ClientToken string `mapstructure:"client_token"`

	/* 资源池ID。您可以调用资源池列表查询获取最新的资源池列表可查询：https://www.ctyun.cn/document/10026730/10040588。 (true) */
	RegionID string `mapstructure:"region_id"`

	/* 可用区名称，4.0资源池必填。您可以调用资源池可用区查询获取资源池可用区列表：https://www.ctyun.cn/document/10026730/10040590。 (false) */
	AzName string `mapstructure:"az_name"`

	/* 云主机名称，只能由数字、字母、-组成，不能以-开头或结尾，不能连续使用-，也不能仅使用数字，且长度为2-15字符。 (false) */
	InstanceName string `mapstructure:"instance_name"`
	InstanceId   string
	/* 云主机显示名称，长度为2-63字符。 (true) */
	DisplayName string `mapstructure:"display_name"`

	/* 规格ID。您可以调用查询一个或多个云主机规格资源获取云主机规格信息。 (true) */
	FlavorID string `mapstructure:"flavor_id"`

	/* 本参数表示镜像类型 ，取值范围：0：私有镜像1：公有镜像2：共享镜像3：安全镜像4：甄选镜像。 (true) */
	ImageType int `mapstructure:"image_type"`

	/* 镜像ID。 (true) */
	ImageID string `mapstructure:"image_id"`

	/* 本参数表示系统盘类型 ，取值范围：SATA：普通云盘SAS：SAS云盘SSD-genric：通用SSD云盘SSD：SSD云盘。 (true) */
	BootDiskType string `mapstructure:"boot_disk_type"`

	/* 系统盘大小单位为GiB，取值范围[40-2048]。 (true) */
	BootDiskSize int `mapstructure:"boot_disk_size"`

	/* 虚拟私有云ID。 (true) */
	VpcID string `mapstructure:"vpc_id"`

	/* 网卡。您可以调用查询网卡列表获取网卡信息及对应的虚拟私有云ID https://www.ctyun.cn/document/10026730/10040207(true) */
	NetworkCardList []vm.NetworkCardList `mapstructure:"network_card_list"`

	/* 本参数表示是否使用弹性公网IP ，取值范围：0：不使用1：自动分配2：使用已有 (true) */
	ExtIP string `mapstructure:"ext_ip"`

	/* 本参数表示购买方式 ，取值范围：false（按周期）true（按需），按周期创建云主机需要同时指定cycleCount和cycleType参数 */
	OnDemand bool `mapstructure:"on_demand"`

	/* 订购时长*/
	CycleCount int `mapstructure:"cycle_count"`

	/* 本参数表示订购周期类型 ，取值范围：MONTH：按月YEAR：按年最长订购周期为5年*/
	CycleType string `mapstructure:"cycle_type"`

	KeyPairID    string `mapstructure:"key_pair_id"`
	RootPassword string `mapstructure:"root_password"`
	BandWidth    int    `mapstructure:"band_width"`

	Comm         communicator.Config `mapstructure:",squash"`
	ArtifactId   string
	ArtifactName string
}

func (ct *CTyunInstanceSpecConfig) Prepare(ctx *interpolate.Context) []error {

	errs := ct.Comm.Prepare(ctx)

	if ct == nil {
		return append(errs, fmt.Errorf("[PRE-FLIGHT] Configuration appears to be empty"))
	}

	if len(ct.ImageID) == 0 {
		errs = append(errs, fmt.Errorf("[PRE-FLIGHT] 'image_id' empty"))
	}

	if len(ct.ClientToken) == 0 {
		errs = append(errs, fmt.Errorf("[PRE-FLIGHT] 'client_token' empty"))
	}

	if len(ct.RegionID) == 0 {
		errs = append(errs, fmt.Errorf("[PRE-FLIGHT] 'region_id' empty"))
	}

	if len(ct.AzName) == 0 {
		errs = append(errs, fmt.Errorf("[PRE-FLIGHT] 'az_name' empty"))
	}

	if len(ct.InstanceName) == 0 {
		errs = append(errs, fmt.Errorf("[PRE-FLIGHT] 'vm_name' empty"))
	}

	if len(ct.DisplayName) == 0 {
		errs = append(errs, fmt.Errorf("[PRE-FLIGHT] 'display_name' empty"))
	}

	if len(ct.FlavorID) == 0 {
		errs = append(errs, fmt.Errorf("[PRE-FLIGHT] 'flavor_id' empty"))
	}

	if len(ct.BootDiskType) == 0 {
		errs = append(errs, fmt.Errorf("[PRE-FLIGHT] 'boot_disk_type' empty"))
	}

	if ct.BootDiskSize < 40 || ct.BootDiskSize > 2048 {
		errs = append(errs, fmt.Errorf("[PRE-FLIGHT] 'syshd' size error"))
	}

	if len(ct.VpcID) == 0 {
		errs = append(errs, fmt.Errorf("[PRE-FLIGHT] 'vpcID' empty"))
	}

	//noPassword := len(ct.Comm.SSHPassword) == 0
	//noKeys := len(ct.Comm.SSHKeyPairName) == 0 && len(ct.Comm.SSHPrivateKeyFile) == 0
	//noTempKey := len(ct.Comm.SSHTemporaryKeyPairName) == 0
	//if noPassword && noKeys && noTempKey {
	//	errs = append(errs, fmt.Errorf("[PRE-FLIGHT] Didn't detect any credentials, you have to specify either "+
	//		"{password} or "+
	//		"{key_name+local_private_key_path} or "+
	//		"{temporary_key_pair_name} cheers :)"))
	//}

	return errs
}
