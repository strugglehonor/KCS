package model

import (
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

type Cluster struct {
	ID  string  `gorm:"_id, omitempty" json:"_id"`
	metav1.ObjectMeta  `gorm:"metadata, omitempty" json:"metadata"`  
	KubeConfig   string  `gorm:"kube_config, omitempty" json:"kube_config"`
	Nodes  []Node  `gorm:"nodes, omitempty" json:"nodes"`
	Region string `gorm:"region, omitempty" json:"region"`
	Description string `gorm:"description, omitempty" json:"description"`
	Status  string  `gorm:"status, omitempty" json:"status"`
	MemLimit  int32 `gorm:"mem_limit, omitempty" json:"mem_limit"`
	CpuLimit  int32 `gorm:"cpu_limit, omitempty" json:"cpu_limit"`
	GpuLimit  int32 `gorm:"gpu_limit, omitempty" json:"gpu_limit"`
	LifeLimit int32 `gorm:"life_limit, omitempty" json:"life_limit"`
	Version  string `gorm:"version, omitempty" json:"version"`
}

