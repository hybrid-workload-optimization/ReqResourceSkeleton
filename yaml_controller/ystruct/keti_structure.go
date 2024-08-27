package ystruct

type Workflow struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

type Metadata struct {
	GenerateName string            `yaml:"generateName"`
	Annotations  map[string]string `yaml:"annotations"`
	Labels       map[string]string `yaml:"labels"`
}

type Spec struct {
	Entrypoint         string     `yaml:"entrypoint"`
	Templates          []Template `yaml:"templates"`
	Arguments          Arguments  `yaml:"arguments"`
	ServiceAccountName string     `yaml:"serviceAccountName"`
}

type Template struct {
	Name         string     `yaml:"name"`
	Container    *Container `yaml:"container,omitempty"`
	Metadata     *Metadata  `yaml:"metadata,omitempty"`
	NodeSelector NodeSelect `yaml:"nodeSelector"`
	DAG          *DAG       `yaml:"dag,omitempty"`
}

type ContainerResources struct {
	Limits   map[string]string `yaml:"limits,omitempty"`
	Requests map[string]string `yaml:"requests,omitempty"`
}

type DAG struct {
	Tasks []Task `yaml:"tasks"`
}

type Task struct {
	Name         string   `yaml:"name"`
	Template     string   `yaml:"template"`
	Dependencies []string `yaml:"dependencies,omitempty"`
}

type Arguments struct {
	Parameters []interface{} `yaml:"parameters"`
}

type NodeSelect struct {
	Node string `yaml:"kubernetes.io/hostname"`
}
