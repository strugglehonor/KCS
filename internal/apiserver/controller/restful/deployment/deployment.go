package deployment

import (
	srvv1 "github.com/strugglehonor/KCS/internal/apiserver/service/v1"
	"github.com/strugglehonor/KCS/internal/apiserver/store"
)

type DeploymentController struct {
	srv srvv1.Service
}

func NewDeploymentController(store store.Factory) *DeploymentController {
	return &DeploymentController{srv: srvv1.NewService(store)}
}
