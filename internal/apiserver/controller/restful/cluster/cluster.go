package cluster

import (
	srvv1 "github.com/strugglehonor/KCS/internal/apiserver/service/v1"
	"github.com/strugglehonor/KCS/internal/apiserver/store"
)

type ClusterController struct {
	srv srvv1.Service
}

func NewClusterController(store store.Factory) *ClusterController {
	return &ClusterController{srv: srvv1.NewService(store)}
}

