package env

var (
	LogFile = ""
)

func init() {
	if _, lf := GetEnv("LOG_FILE"); lf != "" {
		LogFile = lf
	}
}
