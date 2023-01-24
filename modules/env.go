// Read environment variables and set defaults

package modules

import (
	"os"
	"strconv"
)

type Env struct {
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
	LogLevel string
}

func NewEnv() *Env {
	return &Env{
		DbHost:     getEnv("DB_HOST"),
		DbPort:     getEnvInt("DB_PORT", 3306),
		DbUser:		getEnv("DB_USER"),
		DbPassword:	getEnv("DB_PASSWORD"),
		DbName:	getEnv("DB_NAME"),
		LogLevel:   getEnv("LOG_LEVEL", "info")
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return defaultValue
	}
	return value
}




