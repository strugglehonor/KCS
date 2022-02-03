package store

import (
	"context"

	"github.com/strugglehonor/KCS/internal/model"
)

type DeploymentStore interface {
	Create(ctx context.Context, deployment *model.Deployment) error
	Delete(ctx context.Context, depployment *model.Deployment) error 
	InsertDeployment(ctx context.Context, depployment *model.Deployment) error 
	Update(ctx context.Context, deployment *model.Deployment) error
}