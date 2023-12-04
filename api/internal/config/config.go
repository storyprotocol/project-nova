package config

import (
	"flag"
	"strings"

	"github.com/project-nova/backend/pkg/config"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/secrets"
)

type AppConfig struct {
	AppID                  string          `yaml:"app_id"`
	Region                 string          `yaml:"region"`
	Port                   int64           `yaml:"port"`
	Env                    string          `yaml:"env"`
	DatabaseURI            string          `yaml:"database_uri"`
	ProviderURL            string          `yaml:"provider_url"`
	ContentPath            string          `yaml:"content_path"`
	PrimitiveTpeAbiPath    string          `yaml:"primitive_type_abi_path"`
	AdminAuthMessage       string          `yaml:"admin_auth_message"`
	AuthKeyId              string          `yaml:"auth_key_id"`
	S3ContentBucketName    string          `yaml:"s3_content_bucket_name"`
	S3FileUploadBucketName string          `yaml:"s3_file_upload_bucket_name"`
	Protocol               *ProtocolConfig `yaml:"protocol"`
	GrpcWeb3Gateway        string          `yaml:"grpc_web3_gateway"`
	StoryBlocksRegistry    string          `yaml:"story_blocks_registry"`
}

type StreamerConfig struct {
	AppID                string `yaml:"app_id"`
	Region               string `yaml:"region"`
	ProviderWebsocket    string `yaml:"provider_websocket"`
	ApiGatewayUrl        string `yaml:"api_gateway_url"`
	AdminAuthMessage     string `yaml:"admin_auth_message"`
	AuthKeyId            string `yaml:"auth_key_id"`
	OrchestratorContract string `yaml:"orchestrator_contract"`
}

type ProtocolConfig struct {
	ContentUri   string             `yaml:"content_uri"`
	FranchiseMap []*FranchiseConfig `yaml:"franchise_map"`
}

type FranchiseConfig struct {
	FranchiseInfo   *Franchise               `yaml:"franchise_info"`
	ContractInfoMap map[string]*ContractInfo `yaml:"contract_info_map"`
}

type Franchise struct {
	Address            string                 `json:"address" yaml:"address"`
	Name               string                 `json:"name" yaml:"name"`
	VaultAddress       string                 `json:"vaultAddress" yaml:"vault_address"`
	CharacterRegistry  string                 `json:"characterRegistry" yaml:"character_registry"`
	CharacterContracts []*CharacterCollection `json:"characterContract" yaml:"character_contracts"`
	StoryRegistry      string                 `json:"storyRegistry" yaml:"story_registry"`
	StoryContracts     []*StoryCollection     `json:"storyContract" yaml:"story_contracts"`
	LicenseRepository  string                 `json:"licenseRepository" yaml:"license_repository"`
	LicenseRegistry    string                 `json:"licenseRegistry" yaml:"license_registry"`
}

type CharacterCollection struct {
	Name    string `json:"name" yaml:"name"`
	Address string `json:"address" yaml:"address"`
}

type StoryCollection struct {
	Address string `json:"address" yaml:"address"`
	IsCanon bool   `json:"isCanon" yaml:"is_canon"`
}

type ContractType string

var ContractTypes = struct {
	Character ContractType
	Story     ContractType
}{
	Character: "character",
	Story:     "story",
}

type ContractInfo struct {
	Type    ContractType `yaml:"type"`
	IsCanon bool         `yaml:"is_canon"`
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
