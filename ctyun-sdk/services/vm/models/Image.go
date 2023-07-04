package models

type Image struct {

	/*镜像系统架构。*/
	Architecture string `json:"architecture"`
	/*启动方式*/
	BootMode string `json:"bootMode"`
	/*容器格式。*/
	ContainerFormat string `json:"containerFormat"`
	/*镜像系统架构。*/
	CreatedTime int `json:"createdTime"`
	/*镜像描述信息*/
	Description string `json:"description"`
	/*共享镜像的接受人*/
	DestinationUser string `json:"destinationUser"`
	/*磁盘格式*/
	DiskFormat string `json:"diskFormat"`
	/*私有镜像来源的系统盘/数据盘 ID*/
	DiskID string `json:"diskID"`
	/*磁盘容量，单位为 GB*/
	DiskSize int `json:"diskSize"`
	/*镜像ID*/
	ImageClass string `json:"imageClass"`
	/*镜像ID*/
	ImageID string `json:"imageID"`
	/*镜像名称*/
	ImageName string `json:"imageName"`
	/*镜像类型*/
	ImageType string `json:"imageType"`
	/*最大内存*/
	MaximumRAM int `json:"maximumRAM"`
	/*最小内存*/
	MinimumRAM int `json:"minimumRAM"`
	/*操作系统的发行版名称*/
	OsDistro string `json:"osDistro"`
	/*操作系统类型*/
	OsType string `json:"osType"`
	/*操作系统版本*/
	OsVersion string `json:"osVersion"`
	/*项目 ID*/
	ProjectID string `json:"projectID"`
	/*私有镜像的共享列表的总记录数*/
	SharedListLength int `json:"sharedListLength"`
	/*镜像大小，单位为 byte*/
	Size int `json:"size"`
	/*私有镜像来源的云主机/物理机 ID*/
	SourceServerID string `json:"sharedListLength"`
	/*共享镜像的发起人*/
	SourceUser string `json:"sourceUser"`
	/*镜像状态*/
	Status string `json:"status"`
	/*标签。一种场景是标记公共 GPU 镜像适用规格*/
	Tags string `json:"tags"`
	/*镜像更新时间，epoch 秒数，即从 1970-01-01 00:00:00 UTC 到当前时间的秒数*/
	UpdatedTime int `json:"updatedTime"`
	/*镜像可见类型，应始终为“private”（私有镜像）*/
	Visibility string `json:"visibility"`
}
