package config

import (
	"errors"
	"fmt"
	"os"
	"path"

	validator "gopkg.in/validator.v2"
	yaml "gopkg.in/yaml.v2"
)

func IsContainSecrets(fileNames ...string) bool {
	for _, file := range fileNames {
		_, name := path.Split(file)
		if name == "secrets.yaml" {
			return true
		}
	}
	return false
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
