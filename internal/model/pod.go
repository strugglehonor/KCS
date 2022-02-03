package model

type Pod struct {
	ID  string  `gorm:"id, omitempty, unique" json:"id"`
	Ip  string  `gorm:"ip, omitempty" json:"ip"`
	Name string `gorm:"name, omitempty, unique" json:"name"`
	NameSpace string `gorm:"namespace, omitempty" josn:"namespace"`
	Node Node `gorm:"node, omitempty" json:"node"`
	Status string `gorm:"status, omitempty" json:"status"`
	Cluster *Cluster `gorm:"cluster, omitempty" json:"cluster"`
	Volumes []*Volume `gorm:"volume, omitempty" json:"volume"`
	Creator string `gorm:"creator, omitempty" json:"creator"`
	RestartPolicy string `gorm:"restart_policy, omitempty" json:"restart_policy"`
	ImagePullPolicy string `gorm:"image_pull_policy, omitempty" json:"image_pull_policy"`
	DNSPolicy string `gorm:"dns_policy, omitempty" json:"dns_policy"`
	MemLimit  int32 `gorm:"mem_limit, omitempty" json:"mem_limit"`
	CpuLimit  int32 `gorm:"cpu_limit, omitempty" json:"cpu_limit"`
	GpuLimit  int32 `gorm:"gpu_limit, omitempty" json:"gpu_limit"`
	LifeLimit int32 `gorm:"life_limit, omitempty" json:"life_limit"`
	Env  map[interface{}]interface{} `gorm:"env, omitempty" json:"env"`
	HostIPC bool `gorm:"host_ipc, omitempty" json:"host_ipc"`
	HostNetwork bool `gorm:"host_network, omitempty" json:"host_network"`
	HostPID bool `gorm:"host_pid, omitempty" json:"host_pid"`
}