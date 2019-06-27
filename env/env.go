package env

import (
	"os"
)

const (
	EnvNamePrefix    = "TGC_"       // Tron-Go-Common
	EnvNamePrefixEnv = "TGC_PREFIX" // Env var to override "TGC_"

	EnvDev     = "dev"
	EnvStaging = "staging"
	EnvProd    = "prod"
)

var (
	serverEnv = EnvDev
)

// GetEnv checks the prefix settings and returns the correct application-
// specific environment variable value
func GetEnv(name string) (string, string) {
	if pre := os.Getenv(EnvNamePrefixEnv); pre != "" {
		name = pre + name
	} else {
		name = EnvNamePrefix + name
	}
	return name, os.Getenv(name)
}

func init() {
	// Default (and anything else invalid) to dev
	switch _, env := GetEnv("ENV"); env {
	case string(EnvStaging):
		serverEnv = EnvStaging
	case string(EnvProd):
		serverEnv = EnvProd
	}
}

func IsDev() bool {
	return serverEnv == EnvDev
}

func IsStaging() bool {
	return serverEnv == EnvStaging
}

func IsProd() bool {
	return serverEnv == EnvProd
}

func GetCurrentEnv() string {
	return serverEnv
}

func SetCurrentEnv(env string) {
	serverEnv = env
}
