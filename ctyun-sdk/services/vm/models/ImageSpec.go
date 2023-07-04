package models

type ImageSpec struct {

	/* 地域ID  */
	RegionId string `json:"regionID"`

	/* 云主机ID  */
	InstanceId string `json:"instanceID"`

	/* 镜像名称 */
	ImageName string `json:"imageName"`

	/* 镜像描述*/
	Description string `json:"description"`

	/* 企业项目 ID 默认值 0 */
	ProjectID *string `json:"projectID"`
}
