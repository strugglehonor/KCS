package model

type Node struct {
	IP   string  `json:"ip"`
	NodeInfo  NodeInfo  `json:"node_info"`
	Instance  string  `json:"instance"`
}

type NodeInfo struct {
	CpuUsage  float32
	MemUsage  float32
}