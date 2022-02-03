package mysql

import (
	"context"

	"github.com/strugglehonor/KCS/internal/model"
	"gorm.io/gorm"
)

type Deployment struct {
	db *gorm.DB
}

func newDeployment(ms *mysqlstore) *Deployment {
	return &Deployment{db: ms.db}
}

func (d Deployment) Create(ctx context.Context, depployment *model.Deployment) error {
	return d.db.Create(depployment).Error
}

func (d Deployment) Delete(ctx context.Context, depployment *model.Deployment) error {
	return d.db.Delete(depployment).Error
}

func (d Deployment) InsertDeployment(ctx context.Context, depployment *model.Deployment) error {
	return d.db.Create(depployment).Error
}

func (d Deployment) Update(ctx context.Context, depployment *model.Deployment) error {
	return d.db.Save(depployment).Error
}