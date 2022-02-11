package model

import "gorm.io/gorm"

type Node struct {
	gorm.Model
	IP   string  `json:"ip" gorm:"not null; type: varchar(20)"`
	Instance  string  `json:"instance" gorm:"not null; type: varchar(32)"`
	ClusterID uint `json:"cluster_id"`
	PodID uint `json:"pod_id"`
}