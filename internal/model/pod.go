package model

import "gorm.io/gorm"

type Pod struct {
	gorm.Model
	Ip  string  `gorm:"not null; type:varchar(20)" json:"ip"`
	Name string `gorm:"not null; unique; size:32" json:"name"`
	NameSpace string `gorm:"not null; type:varchar(20)" josn:"namespace"`
	Node Node `gorm:"embedded; not null;" json:"node"`
	Status string `gorm:"not null; size:16" json:"status"`
	Cluster Cluster `gorm:"not null; embedded" json:"cluster"`
	Volumes []Volume `json:"volumes"`
	Creator string `gorm:"not null; size:16" json:"creator"`
	RestartPolicy string `gorm:"default:Always; not null" json:"restart_policy"`
	ImagePullPolicy string `gorm:"default:IfNotPresent; not null" json:"image_pull_policy"`
	MemLimit  int32 `json:"mem_limit"`
	CpuLimit  int32 `json:"cpu_limit"`
	GpuLimit  int32 `json:"gpu_limit"`
	LifeLimit int32 `json:"life_limit"`
	Env  map[interface{}]interface{} `json:"env"`
	HostIPC bool `gorm:"default:false" json:"host_ipc"`
	HostNetwork bool `gorm:"default:false" json:"host_network"`
	HostPID bool `gorm:"default:false" json:"host_pid"`
	DeploymentID  uint `json:"deployment_id"`
}