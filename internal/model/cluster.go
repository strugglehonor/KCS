package model

type Cluster struct {
	KubeConfig   string  `json:"kube_config"`
	nodes  []Node
}