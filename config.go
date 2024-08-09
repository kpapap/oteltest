// This file includes the configuration for the module
package oteltest

import (
	"fmt"
	"time"
)

// In order to be instantiated and participate in pipelines, the collector needs to identify your connector and properly load its settings from within its configuration file.
// In order to be able to give your connector access to its settings, create a Config struct.
// The struct must have an exported field for each of the connector’s settings.
// The parameter fields added will be accessible from the config.yaml file.
// Their name in the configuration file is set through a struct tag.
// Create struct and add parameters.
// You can optionally add a validator function to check if the given default values are valid for an instance of your connector.
type Config struct {
    // The name of the configMap.
    ConfigMapName string `mapstructure:"config_map_name"`
		// The interval of the collector.
    Interval string `mapstructure:"interval"`
}

// Validate checks if the receiver configuration is valid
func (cfg *Config) ValidateInterval() error {
	if cfg.Interval == "" {
		return fmt.Errorf("interval cannot be empty")
	}
	interval, err := time.ParseDuration(cfg.Interval)
	if err != nil {
		return fmt.Errorf("invalid interval: %w", err)
	}
	if interval.Minutes() < 1 {
		return fmt.Errorf("interval has to be set to at least 1 minute (1m)")
	}
	return nil
}

// ValidateConfigMapName checks if the configmap name is valid.
func (cfg *Config) ValidateConfigMapName() error {
	if cfg.ConfigMapName == "" {
		return fmt.Errorf("configmap name cannot be empty")
	}
	return nil
}
