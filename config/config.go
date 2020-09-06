package config

import (
	"os"
)

type Environment string

var config map[string]interface{}

const(
	Development = "Development"
	Staging = "Staging"
	Production = "Production"
)

func Port() string {
	port := os.Getenv("__BN_PORT__")
	if port == "" {
		return "8080"
	}
	return port
}

func Env() string {
	env := os.Getenv("__BN_ENVIRONMENT__")
	if env == "" {
		return Development
	}
	return env
}