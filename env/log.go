package env

var (
	LogFile = ""
)

func init() {
	if lf := getEnv("LOG_FILE"); lf != "" {
		LogFile = lf
	}
}
