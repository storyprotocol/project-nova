package config

import (
	"flag"
	"strings"

	"github.com/project-nova/backend/pkg/config"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/secrets"
)

type AppConfig struct {
	AppID            string `yaml:"app_id"`
	Region           string `yaml:"region"`
	DatabaseURI      string `yaml:"database_uri"`
	ProviderURL      string `yaml:"provider_url"`
	Server           Server `yaml:"server"`
	AbiPath          string `yaml:"abi_path"`
	ContentPath      string `yaml:"content_path"`
	AdminAuthMessage string `yaml:"admin_auth_message"`
	AuthKeyId        string `yaml:"auth_key_id"`
}

type StreamerConfig struct {
	AppID             string   `yaml:"app_id"`
	Region            string   `yaml:"region"`
	ProviderWebsocket string   `yaml:"provider_websocket"`
	ApiGatewayUrl     string   `yaml:"api_gateway_url"`
	MonitorAddresses  []string `yaml:"monitor_addresses"`
	AdminAuthMessage  string   `yaml:"admin_auth_message"`
	AuthKeyId         string   `yaml:"auth_key_id"`
}

type Server struct {
	Port int64  `yaml:"port"`
	Env  string `yaml:"env"`
}

var (
	cfgFlag                = flag.String("config", "config.yaml", "config file")
	configInstance         *AppConfig
	streamerConfigInstance *StreamerConfig
)

// GetConfig loads the config and return cached instance once loaded
func GetConfig() (*AppConfig, error) {
	if configInstance != nil {
		return configInstance, nil
	}
	var cfg AppConfig

	cfgFiles := strings.Split(*cfgFlag, ",")
	logger.Info(cfgFiles)

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

func GetStreamerConfig() (*StreamerConfig, error) {
	if streamerConfigInstance != nil {
		return streamerConfigInstance, nil
	}
	var cfg StreamerConfig

	cfgFiles := strings.Split(*cfgFlag, ",")
	logger.Info(cfgFiles)

	if err := config.LoadFiles(&cfg, cfgFiles...); err != nil {
		logger.Fatalf("Failed to load config file: %v", err)
	}

	if !config.IsContainSecrets(cfgFiles...) {
		logger.Infof("Loading secrets %s from secret manager", cfg.AppID)
		if err := secrets.FetchSecrets(cfg.Region, cfg.AppID, &cfg); err != nil {
			return nil, err
		}
	}

	streamerConfigInstance = &cfg
	return &cfg, nil
}
