package v1

import "github.com/strugglehonor/KCS/internal/apiserver/store"

type Service interface {
	Cluster()  ClusterSrv
	Deployment() DeploymentSrv
	Pod() PodSrv
}

type service struct {
	store store.Factory
}

// return service interface
func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}

func (s *service) Cluster() ClusterSrv {
	return newCluster(s)
}

func (s *service) Deployment() DeploymentSrv {
	return newDeployment(s)
}

func (s *service) Pod() PodSrv {
	return newPod(s)
}