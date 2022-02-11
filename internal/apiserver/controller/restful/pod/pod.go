package pod

import (
	srvv1 "github.com/strugglehonor/KCS/internal/apiserver/service/v1"
	"github.com/strugglehonor/KCS/internal/apiserver/store"
)

type PodController struct {
	srv srvv1.Service
}

func newPodController(store store.Factory) PodController {
	return PodController{srv: srvv1.NewService(store)}
}