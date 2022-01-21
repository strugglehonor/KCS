package mysql

import (
	"context"

	"github.com/strugglehonor/KCS/internal/model"
	"gorm.io/gorm"
)

type Pod struct {
	db *gorm.DB
}

func newPod(ms *mysqlstore) *Pod {
	return &Pod{db: ms.db}
}

func (p Pod) Create(ctx context.Context, pod *model.Pod) error {
	return p.db.Create(pod).Error
}