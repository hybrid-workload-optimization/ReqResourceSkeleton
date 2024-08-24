package ystruct

type Container struct {
	Name      string    `yaml:"name"`
	Resources Resources `yaml:"resources"`
}

type Resources struct {
	Requests ResourceDetails `yaml:"requests"`
	Limits   ResourceDetails `yaml:"limits"`
}

type ResourceDetails struct {
	CPU              string `yaml:"cpu"`
	Memory           string `yaml:"memory"`
	GPU              string `yaml:"gpu"`
	EphemeralStorage string `yaml:"ephemeral-storage"`
}
