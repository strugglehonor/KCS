package store

var client Factory

type Factory interface {
	Pod()  PodStore
	Cluster()  ClusterStore
	Volume() VolumeStore
}

func Client() Factory {
	return client
}

func SetClient(factory Factory) {
	client = factory
}