package config

import (
	"os"

	"github.com/Systenix/go-cloud/generators"
	"gopkg.in/yaml.v3"
)

type ConfigData struct {
	ProjectName string                  `yaml:"project_name"`
	ProjectPath string                  `yaml:"project_path"`
	ModulePath  string                  `yaml:"module_path"`
	GoVersion   string                  `yaml:"go_version"`
	Middleware  []generators.Middleware `yaml:"middleware"`
	Services    []generators.Service    `yaml:"services"`
	Models      []generators.Model      `yaml:"models"`
	Events      []generators.Event      `yaml:"events"`
	Docker      generators.Docker       `yaml:"docker"`
	ThirdParty  generators.ThirdParty   `yaml:"third_party"`
}

func ParseConfig(filePath string) (*ConfigData, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config ConfigData
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
