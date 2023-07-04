package basic

import (
	"context"
	"fmt"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer/packer-plugin-sdk/communicator"
	"github.com/hashicorp/packer/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer/packer-plugin-sdk/multistep/commonsteps"
	packersdk "github.com/hashicorp/packer/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer/packer-plugin-sdk/template/config"
	"github.com/hashicorp/packer/packer-plugin-sdk/template/interpolate"
)

/*
 *  此方法返回一个 hcldec.ObjectSpec，这是将 HCL2 模板与 Packer 一起使用所必需的规范
 */
func (b *Builder) ConfigSpec() hcldec.ObjectSpec { return b.config.FlatMapstructure().HCL2Spec() }

/*
 * 将在构建开始时由 Packer 核心调用。它的目的是解析和验证提供给 Packer 的配置模板
 * packer build your_packer_template.json，而不是执行 API 调用或开始创建任何资源或工件
 */
func (b *Builder) Prepare(raws ...interface{}) ([]string, []string, error) {
	err := config.Decode(&b.config, &config.DecodeOpts{
		PluginType:         BUILDER_ID,
		Interpolate:        true,
		InterpolateContext: &b.config.ctx,
		InterpolateFilter: &interpolate.RenderFilter{
			Exclude: []string{
				"boot_command",
			},
		},
	}, raws...)
	if err != nil {
		return nil, nil, fmt.Errorf("[ERROR] Failed in decoding JSON->mapstructure")
	}

	errs := &packersdk.MultiError{}
	errs = packersdk.MultiErrorAppend(errs, b.config.CTyunCredentialConfig.Prepare(&b.config.ctx)...)
	errs = packersdk.MultiErrorAppend(errs, b.config.CTyunInstanceSpecConfig.Prepare(&b.config.ctx)...)
	if errs != nil && len(errs.Errors) != 0 {
		return nil, nil, errs
	}

	packersdk.LogSecretFilter.Set(b.config.AccessKey, b.config.SecretKey)

	return nil, nil, nil
}

/*
 * 通常为多个构建器并行执行，以实际构建机器、配置它并创建生成的机器映像
 */
func (b *Builder) Run(ctx context.Context, ui packersdk.Ui, hook packersdk.Hook) (packersdk.Artifact, error) {

	state := new(multistep.BasicStateBag)
	state.Put("hook", hook)
	state.Put("ui", ui)
	state.Put("config", &b.config)
	//定义执行步骤的数组接口
	steps := []multistep.Step{

		&stepValidateParameters{
			InstanceSpecConfig: &b.config.CTyunInstanceSpecConfig,
		},

		&stepConfigCredentials{
			InstanceSpecConfig: &b.config.CTyunInstanceSpecConfig,
		},

		&stepCreateCTyunInstance{
			InstanceSpecConfig: &b.config.CTyunInstanceSpecConfig,
			CredentialConfig:   &b.config.CTyunCredentialConfig,
		},

		&communicator.StepConnect{
			Config:    &b.config.CTyunInstanceSpecConfig.Comm,
			SSHConfig: b.config.CTyunInstanceSpecConfig.Comm.SSHConfigFunc(),
			Host:      instanceHost,
		},

		&commonsteps.StepProvision{},

		&stepStopCTCloudInstance{
			InstanceSpecConfig: &b.config.CTyunInstanceSpecConfig,
		},

		&stepCreateCTyunImage{
			InstanceSpecConfig: &b.config.CTyunInstanceSpecConfig,
		},
	}

	b.runner = commonsteps.NewRunnerWithPauseFn(steps, b.config.PackerConfig, ui, state)
	b.runner.Run(ctx, state)

	if rawErr, ok := state.GetOk("error"); ok {
		return nil, rawErr.(error)
	}

	artifact := &Artifact{
		ImageId:   b.config.ArtifactId,
		RegionID:  b.config.RegionId,
		ImageName: b.config.ArtifactName,
		StateData: map[string]interface{}{"generated_data": state.Get("generated_data")},
	}
	return artifact, nil
}
