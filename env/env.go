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

// getEnv checks the prefix settings and returns the correct application-
// specific environment variable value
func getEnv(name string) string {
	if pre := os.Getenv(EnvNamePrefixEnv); pre != "" {
		return os.Getenv(pre + name)
	}
	return os.Getenv(EnvNamePrefix + name)
}

func init() {
	// Default (and anything else invalid) to dev
	switch env := getEnv("ENV"); env {
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
