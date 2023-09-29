package cleanup

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

const FrigateDefaultMaxRetentionDays = 10

type FrigateConfig struct {
	Snapshots FrigateConfigSnapshots `yaml:"snapshots"`
}

type FrigateConfigSnapshots struct {
	Retain FrigateConfigRetain `yaml:"retain"`
}

type FrigateConfigRetain struct {
	Default int            `yaml:"int"`
	Objects map[string]int `yaml:"objects"`
}

func ParseFrigateConfig(configFilePath string) (FrigateConfig, error) {
	cfg := FrigateConfig{}

	content, err := os.ReadFile(configFilePath)
	if err != nil {
		return cfg, fmt.Errorf("could not read frigate config path: %w", err)
	}

	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("could not unmarshal frigate config: %w", err)
	}

	return cfg, nil
}

func getMaxRetentionFromConfigData(configData FrigateConfig) int {
	maxRetention := configData.Snapshots.Retain.Default

	for _, value := range configData.Snapshots.Retain.Objects {
		if value > maxRetention {
			maxRetention = value
		}
	}

	if maxRetention <= 0 {
		maxRetention = FrigateDefaultMaxRetentionDays
	}

	return maxRetention
}
