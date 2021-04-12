package env

import (
	"log"
	"os"
	"strings"
)

var logFatalf = log.Fatalf

// GetAndValidateF exits the app if the environment variable is empty
func GetAndValidateF(key string) string {
	val := strings.TrimSpace(os.Getenv(key))

	if val == "" {
		logFatalf("[env] %s - not found", key)
	}
	return val
}
