package config

import (
	"flag"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

var readConfig = false

// getEnv -> parses flag to get environment
func getEnv() string {
	var env string
	flag.StringVar(&env, "env", "local", "Environment to fetch config from")
	flag.Parse()
	return env
}

// getAbsoluteFile -> gets absolute path of config
func getAbsoluteFile(relativeFilePath string) string {
	// Get absolute file path of the config file
	filePath, err := filepath.Abs(relativeFilePath)
	if err != nil {
		panic("Unable to read config")
	}

	return filePath
}

// getConfigFilePath -> gets config file path fot env
func getConfigFilePath() string {
	var env string = getEnv()
	var rawFilePath string
	if env == "prod" {
		rawFilePath = "config/props/prod.env"
	} else {
		rawFilePath = "config/props/local.env"
	}
	return getAbsoluteFile(rawFilePath)
}

// Load -> read props from taml
func Load() {
	if readConfig {
		return
	}
	var filePath string = getConfigFilePath()
	godotenv.Load(filePath)
	readConfig = true
}

// Get -> gets string value of env variable
func Get(key string) string {
	if !readConfig {
		Load()
	}
	return os.Getenv(key)
}

// GetInt -> get int value of env variable
func GetInt(key string) int {
	if !readConfig {
		Load()
	}

	intData, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		panic(err.Error())
	}
	return intData
}
