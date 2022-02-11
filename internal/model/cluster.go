package model

import (
	"gorm.io/gorm"
)

type Cluster struct {
	gorm.Model  
	KubeConfig   string  `gorm:"not null; type:varchar(32)" json:"kube_config"`
	Nodes  []Node  `gorm:"not null;" json:"nodes"`
	Region string `gorm:"not null; size:32" json:"region"`
	Description string `gorm:"type: varchar(1024)" json:"description"`
	Status  string  `gorm:"not null; size:16" json:"status"`
	MemLimit  int32 `json:"mem_limit"`
	CpuLimit  int32 `json:"cpu_limit"`
	GpuLimit  int32 `json:"gpu_limit"`
	LifeLimit int32 `json:"life_limit"`
	Version  string `gorm:"not null" json:"version"`
	PodID  uint `json:"pod_id"`
}

