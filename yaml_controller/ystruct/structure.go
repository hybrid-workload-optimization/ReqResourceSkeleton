package ystruct

type Example struct {
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}
