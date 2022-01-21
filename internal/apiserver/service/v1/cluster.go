package v1

import (
	"context"

	"github.com/strugglehonor/KCS/internal/model"
	"github.com/strugglehonor/KCS/internal/apiserver/store"
)

type ClusterSrv interface {
	Create(ctx context.Context, cluster *model.Cluster) error
	GetAllCluster(ctx context.Context)
	GetClusterByID(ctx context.Context)
}

type ClusterService struct {
	store  store.Factory
}

var _ ClusterSrv = (*ClusterService)(nil)

func newCluster(s *service) ClusterSrv {
	return &ClusterService{store: s.store}
} 

func (clusterSrv *ClusterService) Create(ctx context.Context, cluster *model.Cluster) error {
	if err := clusterSrv.store.Cluster().Create(ctx, cluster); err != nil {
		return err
	}

	return nil
 }

func (clusterSrv *ClusterService) GetAllCluster(ctx context.Context) {

}

func (ClusterSrv *ClusterService) GetClusterByID(ctx context.Context) {
	
}