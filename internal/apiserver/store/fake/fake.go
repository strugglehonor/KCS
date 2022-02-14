package fake

import (
	"fmt"
	"sync"

	"github.com/strugglehonor/KCS/internal/model"
	"github.com/strugglehonor/KCS/pkg/random"
)

var (
	Pending = "Pending"
	Running = "Running"
	Submit = "Submit"
	Success = "Success"
	Failed = "Failed"

	K8sVersion = "1.23.1"

	Beijing = "Beijing"
	Shenzhen = "Shenzhen"
)

type datastore struct {
	sync.RWMutex
	pods []model.Pod
	deployments []model.Deployment
	volumes []model.Volume
	cluster []model.Cluster
}

func FakePod(count int) []*model.Pod {
	pods := make([]*model.Pod, 0)
	for i:=0; i<count; i++ {
		pods = append(pods, &model.Pod{
			Ip: random.RandomString(16),
			NameSpace: random.RandomString(5),
			Status: Pending,
			Cluster: &model.Cluster{},
			Volumes: FakeVolume(10),
			Creator: fmt.Sprintf("creator_%d", i),
			Name: fmt.Sprintf("name_%d", i),
		})
	}
	return pods
}

func FakeVolume(count int) []*model.Volume {
	volumes := make([]*model.Volume, 0)
	for i:=0; i<count; i++ {
		volumes = append(volumes, &model.Volume{
			
		})
	}
	return volumes
}

func FakeCluster(count int) []*model.Cluster {
	clusters := make([]*model.Cluster, 0)
	for i:=0; i<count; i++ {
		clusters = append(clusters, &model.Cluster{
			KubeConfig: fmt.Sprintf("kubeconfig_%d", i),
			Region: Beijing,
			Nodes: []model.Node{},
			Status: Pending,
			Version: K8sVersion,
		})
	}
	return clusters
}

func FakeDeployment(count int) []*model.Deployment {
	deployments := make([]*model.Deployment, 0)
	for i:=0; i<count; i++ {
		deployments = append(deployments, &model.Deployment{
			Pod: FakePod(10),
		})
	}
	return deployments
}