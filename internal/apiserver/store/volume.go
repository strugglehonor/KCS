package store

import (
	"context"

	"github.com/strugglehonor/KCS/internal/model"
)

type VolumeStore interface {
	Create(ctx context.Context, volume *model.Volume) error
}