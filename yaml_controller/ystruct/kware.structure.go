package ystruct

type ReqResource struct {
	Version string  `yaml:"version"`
	Request Request `yaml:"request"`
}

type Request struct {
	Name       string      `yaml:"name"`
	ID         string      `yaml:"id"`
	Date       string      `yaml:"date"`
	Containers []Container `yaml:"containers"`
	Attribute  Attribute   `yaml:"attribute"`
}

type Attribute struct {
	WorkloadType           string  `yaml:"workloadType"`
	IsCronJob              bool    `yaml:"isCronJob"`
	DevOpsType             string  `yaml:"devOpsType"`
	CudaVersion            float64 `yaml:"cudaVersion"`
	GPUDriverVersion       float64 `yaml:"gpuDriverVersion"`
	MaxReplicas            int     `yaml:"maxReplicas"`
	IsNetworking           bool    `yaml:"isNetworking"`
	TotalSize              int     `yaml:"totalSize"`
	PredictedExecutionTime int     `yaml:"predictedExecutionTime"`
	UserID                 string  `yaml:"userId"`
	Yaml                   string  `yaml:"yaml"`
}

type RespResource struct {
	Response Response `yaml:"response"`
}

type Response struct {
	ID               string      `yaml:"id"`
	Date             string      `yaml:"date"`
	Container        []Container `yaml:"container,omitempty"`
	PriorityClass    string      `yaml:"priorityClass"`
	Priority         string      `yaml:"string"`
	PreemptionPolicy string      `yaml:"preemptionPolicy"`
}

type Result struct {
	Cluster          string `yaml:"cluster"`
	Node             string `yaml:"node"`
	PriorityClass    string `yaml:"priorityClass"`
	Priority         string `yaml:"string"`
	PreemptionPolicy string `yaml:"preemptionPolicy"`
}
