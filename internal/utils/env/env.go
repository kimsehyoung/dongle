package env

import (
	"os"
	"strconv"
)

func GetEnvString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetEnvBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		if value, err := strconv.ParseBool(value); err == nil {
			return value
		}
	}
	return defaultValue
}

func GetEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if value, err := strconv.Atoi(value); err == nil {
			return value
		}
	}
	return defaultValue
}
