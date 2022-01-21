package store

import (
	"context"

	"github.com/strugglehonor/KCS/internal/model"
)

type PodStore interface {
	Create(ctx context.Context, pod *model.Pod) error
}