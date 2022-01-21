package mysql

import (
	"context"

	"github.com/strugglehonor/KCS/internal/model"
	"gorm.io/gorm"
)

type Volume struct {
	db *gorm.DB
}

func newVolume(ms *mysqlstore) *Volume {
	return &Volume{db: ms.db}
}

func (v Volume) Create(ctx context.Context, volume *model.Volume) error {
	return v.db.Create(volume).Error
}