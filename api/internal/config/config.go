package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/project-nova/backend/pkg/logger"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	AppID string `yaml:"app_id"`
	Port  int64  `yaml:"port"`
	Env   string `yaml:"env"`
}

var (
	cfgFlag        = flag.String("config", "config.yaml", "config file")
	configInstance *AppConfig
)

// GetConfig loads the config and return cached instance once loaded
func GetConfig() (*AppConfig, error) {
	if configInstance != nil {
		return configInstance, nil
	}
	var cfg AppConfig

	cfgFiles := strings.Split(*cfgFlag, ",")
	logger.Info(cfgFiles)

	if err := LoadFiles(&cfg, cfgFiles...); err != nil {
		logger.Fatalf("Failed to load config file: %v", err)
	}

	configInstance = &cfg
	return &cfg, nil
}

// LoadFiles combines configs in the files and validate the configs
func LoadFiles(config interface{}, fileNames ...string) error {
	if len(fileNames) == 0 {
		return errors.New("no config files to load")
	}

	for _, file := range fileNames {
		fmt.Println("Load config file: " + file)
		data, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		if err := yaml.Unmarshal(data, config); err != nil {
			return err
		}
	}

	if err := validator.Validate(config); err != nil {
		return err.(validator.ErrorMap)
	}
	return nil
}
