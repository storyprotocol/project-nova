package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/project-nova/backend/pkg/logger"
	validator "gopkg.in/validator.v2"
	yaml "gopkg.in/yaml.v2"
)

type AppConfig struct {
	DatabaseURI string `yaml:"database_uri"`
	Server      Server `yaml:"server"`
	AbiPath     string `yaml:"abi_path"`
	ContentPath string `yaml:"content_path"`
}

type Server struct {
	Port int64  `yaml:"port"`
	Env  string `yaml:"env"`
}

var (
	cfgFlag = flag.String("config", "config.yaml", "config file")
)

// InitializeConfigWithFlag loads config with list of config files and also loads secrets
func InitializeConfigWithFlag() (*AppConfig, error) {
	cfgFiles := strings.Split(*cfgFlag, ",")
	logger.Info("cfgFiles, %v", cfgFiles)
	var cfg AppConfig
	if err := LoadFiles(&cfg, cfgFiles...); err != nil {
		logger.Warn("failed to initialize configuration ", "error: ", err)
		return nil, err
	}
	return &cfg, nil
}

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
