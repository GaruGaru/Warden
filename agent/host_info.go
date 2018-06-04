package agent

type HostInfo struct {
	Hostname        string `json:"hostname"`
	UpTime          uint64 `json:"up_time"`
	OS              string `json:"os"`
	Platform        string `json:"platform"`
	PlatformFamily  string `json:"platform_family"`
	PlatformVersion string `json:"platform_version"`
}

type CpuInfo struct {
	Vendor            string    `json:"vendor"`
	Family            string    `json:"family"`
	Model             string    `json:"model"`
	Cores             int32     `json:"cores"`
	ModelName         string    `json:"model_name"`
	Frequency         float64   `json:"frequency"`
	UsagePercent      []float64 `json:"usage_percent"`
	UsagePercentTotal float64   `json:"usage_percent_total"`
}

type MemoryInfo struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
}

type DiskInfo struct {
	Name              string  `json:"name"`
	Mount             string  `json:"mount"`
	Total             uint64  `json:"total"`
	Free              uint64  `json:"free"`
	Used              uint64  `json:"used"`
	UsedPercent       float64 `json:"usedPercent"`
	InodesTotal       uint64  `json:"inodesTotal"`
	InodesUsed        uint64  `json:"inodesUsed"`
	InodesFree        uint64  `json:"inodesFree"`
	InodesUsedPercent float64 `json:"inodesUsedPercent"`
}

type AgentInfo struct {
	Host       HostInfo  `json:"host"`
	CpusInfo   []CpuInfo `json:"cpus_info"`
	MemoryInfo `json:"memory"`
	Disks      []DiskInfo `json:"disks"`
}
