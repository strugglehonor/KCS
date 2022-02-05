package v1

import (
	"context"

	"github.com/strugglehonor/KCS/internal/apiserver/store"
	"github.com/strugglehonor/KCS/internal/model"
)

type DeploymentSrv interface {
	Create(ctx context.Context, deployment *model.Deployment) error
	Delete(ctx context.Context, deployment *model.Deployment) error 
	Update(ctx context.Context, deployment *model.Deployment) error
}

type DeploymentService struct {
	store store.Factory
}

// 显示判断是否进行了接口
var _ DeploymentSrv = (*DeploymentService)(nil)

func newDeployment(s *service) DeploymentSrv {
	return &DeploymentService{store: s.store}
}

func (deploymentService *DeploymentService) Create(ctx context.Context, deployment *model.Deployment) error {
	if err := deploymentService.store.Deployment().Create(ctx, deployment); err != nil {
		return err
	}
	return nil
}

func (deploymentService *DeploymentService) Delete(ctx context.Context, deployment *model.Deployment) error {
	if err := deploymentService.store.Deployment().Delete(ctx, deployment); err != nil {
		return err
	}
	return nil
}

func (DeploymentService *DeploymentService) Update(ctx context.Context, deployment *model.Deployment) error {
	if err := DeploymentService.store.Deployment().Update(ctx, deployment); err != nil {
		return err
	}
	return nil
}
