package store

import (
	"context"

	"github.com/strugglehonor/KCS/internal/model"
)

type PodStore interface {
	Create(ctx context.Context, pod *model.Pod) error
	Delete(ctx context.Context, pod *model.Pod) error
	GetPodByName(ctx context.Context, podName string) (*model.Pod, error)
	InsertPod(ctx context.Context, pod *model.Pod) error
	Update(ctx context.Context, pod *model.Pod) error
	GetAllPod(ctx context.Context, limit, offset *int64) ([]*model.Pod, error) 
}