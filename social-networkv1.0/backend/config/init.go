// config/config.go

package config

import (
	"os"
)

// getEnv tries to get an environment variable; if not found, it returns a default value.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
