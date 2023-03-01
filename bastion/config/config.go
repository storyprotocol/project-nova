package config

import (
	"fmt"
	"os"

	"github.com/project-nova/backend/pkg/config"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/secrets"
)

type TaskConfig struct {
	AppID                 string `yaml:"app_id"`
	Region                string `yaml:"region"`
	DatabaseURI           string `yaml:"database_uri"`
	ProviderURL           string `yaml:"provider_url"`
	ApiGatewayUrl         string `yaml:"api_gateway_url"`
	AdminAuthMessage      string `yaml:"admin_auth_message"`
	AuthKeyId             string `yaml:"auth_key_id"`
	S3OperationBucketName string `yaml:"s3_operation_bucket_name"`
}

const (
	configPath = "/bastion/config"
)

// GetConfig loads the config
func GetConfig() (*TaskConfig, error) {
	var cfg TaskConfig

	path, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get config directories")
	}

	path = path + configPath

	cfgFiles := []string{path + "/base.yaml"}
	env := config.GetEnv()
	switch env {
	case config.Environments.Development:
		cfgFiles = append(cfgFiles, path+"/dev.yaml")
	case config.Environments.Staging:
		cfgFiles = append(cfgFiles, path+"/staging.yaml")
	case config.Environments.Production:
		cfgFiles = append(cfgFiles, path+"/prod.yaml")
	case config.Environments.LocalDocker:
		cfgFiles = append(cfgFiles, path+"/local-docker.yaml")
		cfgFiles = append(cfgFiles, path+"/secrets.yaml")
	case config.Environments.Local:
		cfgFiles = append(cfgFiles, path+"/local.yaml")
		cfgFiles = append(cfgFiles, path+"/secrets.yaml")
	default:
		return nil, fmt.Errorf("unknown environment: %s", env)
	}

	if err := config.LoadFiles(&cfg, cfgFiles...); err != nil {
		logger.Fatalf("Failed to load config file: %v", err)
	}

	if !config.IsContainSecrets(cfgFiles...) {
		logger.Infof("Loading secrets %s from secret manager", cfg.AppID)
		if err := secrets.FetchSecrets(cfg.Region, cfg.AppID, &cfg); err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}
