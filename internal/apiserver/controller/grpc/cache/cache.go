package cache

import (
	"context"
	"fmt"
	"sync"

	pb "github.com/strugglehonor/KCS/internal/apiserver/controller/grpc/proto/v1"
	"github.com/strugglehonor/KCS/internal/apiserver/store"
	"github.com/strugglehonor/KCS/pkg/log"
)

var (
	once sync.Once
	c *Cache
)

type Cache struct {
	store store.Factory
}

func GetCacheOr(store store.Factory) (*Cache, error) {
	if store != nil {
		once.Do(func() {
			c = &Cache{store: store}
		})
	}

	if c == nil {
		return nil, fmt.Errorf("CacheIns init failed, got nil CacheIns")
	}

	return c, nil
}

func (c *Cache) ListPods(ctx context.Context, in *pb.ListPodsRequest) (*pb.ListPodsResponse, error) {
	log.INFO.Println("RPC ListPods called")
	pods, err := c.store.Pod().GetAllPod(ctx, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}

	items := make([]*pb.PodsInfo, 0)
	for _, pod := range pods {
		items = append(items, &pb.PodsInfo{
			Ip: pod.Ip,
			Status: pod.Status,
			Name: pod.Name,
			Namespace: pod.NameSpace,
			Creator: pod.Creator,
			RestartPolicy: pod.RestartPolicy,
			ImagePullPolicy: pod.ImagePullPolicy,
			MemLimit: pod.MemLimit,
			CpuLimit: pod.CpuLimit,
			GpuLimit: pod.GpuLimit,
			LifeLimit: pod.LifeLimit,
			HostIpc: pod.HostIPC,
			HostNetwork: pod.HostNetwork,
			HostPid: pod.HostPID,
		})
	}

	totalCount := len(pods)
	return &pb.ListPodsResponse{TotalCount: int64(totalCount), Items: items}, nil
}

// deployment/ cluster to do 
//func (c *Cache) ListDeployment(ctx context.Context, in )