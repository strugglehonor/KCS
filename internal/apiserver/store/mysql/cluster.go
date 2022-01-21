package mysql

import (
	"context"

	"github.com/strugglehonor/KCS/internal/model"
	"gorm.io/gorm"
)

type Cluster struct {
	db *gorm.DB
}

func newCluster(ms *mysqlstore) *Cluster {	
	return &Cluster{db: ms.db}
}
 
func (c Cluster) Create(ctx context.Context, cluster *model.Cluster) error {
	return c.db.Create(cluster).Error
}