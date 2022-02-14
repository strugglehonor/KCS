package fake

import (
	"context"
	"errors"

	"github.com/strugglehonor/KCS/internal/model"
)

type pods struct {
	ds *datastore
}

func newPods(ds *datastore) *pods {
	return &pods{ds}
}

func (p pods) Create(ctx context.Context, pod *model.Pod) error {
	p.ds.Lock()
	defer p.ds.Unlock()

	for _, pol := range p.ds.pods {
		if pol.Name == pod.Name {
			return errors.New("record already exits")
		}
	}

	if len(p.ds.pods) > 0 {
		pod.ID = p.ds.pods[len(p.ds.pods)-1].ID + 1
	}
	p.ds.pods = append(p.ds.pods, *pod)

	return nil
}