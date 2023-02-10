package config

import (
	"flag"
	"strings"

	"github.com/project-nova/backend/pkg/config"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/secrets"
)

type AppConfig struct {
	AppID       string `yaml:"app_id"`
	Region      string `yaml:"region"`
	DatabaseURI string `yaml:"database_uri"`
	ProviderURL string `yaml:"provider_url"`
	Server      Server `yaml:"server"`
	AbiPath     string `yaml:"abi_path"`
	ContentPath string `yaml:"content_path"`
}

type Server struct {
	Port int64  `yaml:"port"`
	Env  string `yaml:"env"`
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
	cfgFiles := strings.Split(*cfgFlag, ",")
	logger.Info(cfgFiles)
	var cfg AppConfig
	if err := config.LoadFiles(&cfg, cfgFiles...); err != nil {
		logger.Fatalf("Failed to load config file: %v", err)
	}

	if !config.IsContainSecrets(cfgFiles...) {
		logger.Infof("Loading secrets %s from secret manager", cfg.AppID)
		if err := secrets.FetchSecrets(cfg.Region, cfg.AppID, &cfg); err != nil {
			return nil, err
		}
	}

	configInstance = &cfg
	return &cfg, nil
}
