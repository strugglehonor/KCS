package model

import "gorm.io/gorm"

type Deployment struct {
	gorm.DB
	Pod  []Pod `gorm:"embedded" json:"pod"`
}