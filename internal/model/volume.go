package model

type Volume struct {
	ID  string `gorm:"id, omitempty" json:"id"`
	Name string `gorm:"name, omitempty" json:"name"`
	VolumeType string `gorm:"volume_type, omitempty" json:"volume_type"`
	MountPath string `gorm:"mount_path, omitempty" json:"mount_path"`
}