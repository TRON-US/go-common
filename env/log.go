package env

import (
	"os"
)

var (
	LogFile = ""
)

func init() {
	if lf := getEnv("LOG_FILE"); lf != "" {
		LogFile = lf
	}
}
