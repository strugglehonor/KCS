package model

import "gorm.io/gorm"

type Volume struct {
	gorm.Model
	PodID  uint  `json:"pod_id"`
	Name string `gorm:"not null; type:varchar(20)" json:"name"`
	VolumeType string `gorm:"not null; size:16" json:"volume_type"`
	MountPath string `gorm:"not null; type:varchar(32)" json:"mount_path"`
}