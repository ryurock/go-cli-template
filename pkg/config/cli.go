package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type CliConfig struct {
	Config      CliConfigRoot `yaml:"config"`
	Description string        `yaml:"description"`
}

type CliConfigRoot struct {
	Name        string
	Description CliConfigRootDescription
}

type CliConfigRootDescription struct {
	Short string
	Long  string
}

func NewCliConfig() *CliConfig {
	var config CliConfig
	b, err := os.ReadFile("cmd/cli/config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(b, &config)
	if err != nil {
		panic(err)
	}

	return &config
}
