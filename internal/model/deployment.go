package model

import "gorm.io/gorm"

type Deployment struct {
	gorm.Model
	Pods  []*Pod `json:"pods"`
}