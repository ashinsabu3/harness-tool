package config_parser

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type HarnessMeta struct {
	ApiKey    string `yaml:"apiKey"`
	AccountId string `yaml:"accountId"`
	OrgId     string `yaml:"orgId"`
	ProjectId string `yaml:"projectId"`
}

type InputSetItem struct {
	Name  string `yaml:"name"`
	Type  string `yaml:"type"`
	Value string `yaml:"value"`
}

type InputSet map[string][]InputSetItem

type Pipeline struct {
	Id        string   `yaml:"id"`
	InputSets InputSet `yaml:"inputsets"`
}

type HarnessConfig struct {
	HarnessMeta      HarnessMeta `yaml:"harness-meta"`
	HarnessPipelines []Pipeline  `yaml:"harness-pipelines"`
}

func ParseConfig(filePath string) (*HarnessConfig, error) {
	// Read the YAML file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Parse the YAML file into the struct
	var config HarnessConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml: %w", err)
	}

	return &config, nil
}
