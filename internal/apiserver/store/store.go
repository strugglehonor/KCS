package store

var client Factory

type Factory interface {
	Pod()  PodStore
	Cluster()  ClusterStore
	Volume() VolumeStore
	Deployment() DeploymentStore
}

func Client() Factory {
	return client
}

func SetClient(factory Factory) {
	client = factory
}