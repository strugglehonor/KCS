package v1

import (
	"context"

	"github.com/strugglehonor/KCS/internal/model"
	"github.com/strugglehonor/KCS/internal/apiserver/store"
)

type PodSrv interface {
	Create(ctx context.Context, pod *model.Pod) error
	Delete(ctx context.Context)
}

type PodService struct {
	store store.Factory
}

var _ PodSrv = (*PodService)(nil)

func newPod(s *service) PodSrv {
	return &PodService{store: s.store}
}

func (podService *PodService) Create(ctx context.Context, pod *model.Pod) error {
	if err := podService.store.Pod().Create(ctx, pod); err != nil {
		return err
	}

	return nil
}

func (podService *PodService) Delete(ctx context.Context) {

}