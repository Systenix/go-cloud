package config

import (
	"os"

	"github.com/Systenix/go-cloud/generators"
	"gopkg.in/yaml.v3"
)

type ConfigData struct {
	ProjectName string               `yaml:"project_name"`
	ProjectPath string               `yaml:"project_path"`
	ModulePath  string               `yaml:"module_path"`
	Services    []generators.Service `yaml:"services"`
	Models      []generators.Model   `yaml:"models"`
	Events      []generators.Event   `yaml:"events"`
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
