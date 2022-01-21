package store

import (
	"context"

	"github.com/strugglehonor/KCS/internal/model"
)

type ClusterStore interface {
	Create(ctx context.Context, cluster *model.Cluster) error
	
}