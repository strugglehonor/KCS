package v1

import (
	"context"

	"github.com/strugglehonor/KCS/internal/apiserver/store"
	"github.com/strugglehonor/KCS/internal/model"
	"github.com/strugglehonor/KCS/pkg/redis"
)

var (
	PodChangedNotification = "k8s pod changed notification" 
)

type PodSrv interface {
	Create(ctx context.Context, pod *model.Pod) error
	Delete(ctx context.Context, pod *model.Pod) error
	GetPodByName(ctx context.Context, podName string) (*model.Pod, error) 
	InsertPod(ctx context.Context, ID, name, namespace, status, creator, restartPolicy,
		imagePullPolicy, dnsPolicy, ip string, memLimit, cpulimit, gpulimit, lifelimit int32, env map[interface{}]interface{},
		hostIPC, hostNetwork, hostPID bool, node model.Node, cluster *model.Cluster) error
	Update(ctx context.Context, pod *model.Pod) error
}

type PodService struct {
	store store.Factory
	redisCli redis.RedisCli
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

func (podService *PodService) Delete(ctx context.Context, pod *model.Pod) error {
	if err := podService.store.Pod().Delete(ctx, pod); err != nil {
		return err
	}
	return nil
}

func (podService *PodService) GetPodByName(ctx context.Context, podName string) (*model.Pod, error) {
	pod, err := podService.store.Pod().GetPodByName(ctx, podName)
	if err != nil {
		return nil, err
	} 

	return pod, err
}

func (podService *PodService) InsertPod(ctx context.Context, ID, name, namespace, status, creator, restartPolicy,
    imagePullPolicy, dnsPolicy, ip string, memLimit, cpulimit, gpulimit, lifelimit int32, env map[interface{}]interface{},
    hostIPC, hostNetwork, hostPID bool, node model.Node, cluster *model.Cluster) error {
	
	pod := model.Pod{
		Ip: ip,
		Name: name,
		NameSpace: namespace,
		Node: node,
		Status: status,
		Cluster: cluster,
		Creator: creator,
		RestartPolicy: restartPolicy,
		ImagePullPolicy: imagePullPolicy,
		MemLimit: memLimit,
		CpuLimit: cpulimit,
		GpuLimit: gpulimit,
		LifeLimit: lifelimit,
		Env: env,
		HostIPC: hostIPC,
		HostNetwork: hostNetwork,
		HostPID: hostPID,
	}

	if err := podService.store.Pod().Create(ctx, &pod); err != nil {
		return err
	}
	return nil
}

func (podService *PodService) Update(ctx context.Context, pod *model.Pod) error {
	if err := podService.store.Pod().Update(ctx, pod); err != nil {
		return err
	} 
	return nil
}