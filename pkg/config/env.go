package config

import "os"

const (
	ServerENV = "SERVER_ENV"
)

type Environment string

// Environments listed all environments
var Environments = struct {
	Development Environment
	Staging     Environment
	Production  Environment
	Local       Environment
	LocalDocker Environment
	CI          Environment
	Unknown     Environment
}{
	Development: "dev",
	Staging:     "staging",
	Production:  "prod",
	Local:       "local",
	LocalDocker: "local-docker",
	CI:          "ci",
	Unknown:     "unknown",
}

func GetEnv() Environment {
	env := Environment(os.Getenv(ServerENV))
	switch env {
	case Environments.Local, Environments.LocalDocker, Environments.Development, Environments.Staging, Environments.CI, Environments.Production:
		return env
	}
	return Environments.Unknown
}
