package config

import (
	"os"
	"path/filepath"

	"github.com/apex/log"
	"github.com/joho/godotenv"
)

const (
	DB_USERNAME    = "DB_USERNAME"
	DB_PASSWORD    = "DB_PASSWORD"
	DB_DATABASE    = "DB_DATABASE"
	DB_SERVER_HOST = "DB_SERVER_HOST"
	CLIENT_URL     = "CLIENT_URL"
	SERVER_PORT    = "SERVER_PORT"
	JWT_KEY        = "JWT_KEY"
	RUN_MIGRATION  = "RUN_MIGRATION"
	ENVIRONEMT     = "ENV"
)

type ConfigType map[string]string

var ConfigSettings = ConfigType{
	DB_USERNAME:    "",
	DB_PASSWORD:    "",
	DB_DATABASE:    "",
	CLIENT_URL:     "",
	SERVER_PORT:    "",
	JWT_KEY:        "",
	RUN_MIGRATION:  "",
	DB_SERVER_HOST: "",
}

var requiredData = map[string]bool{
	DB_USERNAME:     true,
	DB_PASSWORD: true,
	DB_DATABASE:       true,
	DB_SERVER_HOST: true,
	CLIENT_URL:        true,
	SERVER_PORT:       true,
	RUN_MIGRATION:     true,
}

func InitilizeConfig() {
	environment, exists := os.LookupEnv(ENVIRONEMT)
	var localEnvFilePath string
	if exists && environment == "test" {
		localEnvFilePath, _ = filepath.Abs("./.env.test")
	} else {
		localEnvFilePath, _ = filepath.Abs("./.env")
	}

	if err := godotenv.Load(localEnvFilePath); err != nil {
		log.WithField("reason", err.Error()).Fatal("No .env file found")
	}

	for key := range ConfigSettings {
		envVal, exists := os.LookupEnv(key)
		if !exists {
			if requiredData[key] {
				log.Fatal(key + " not found in env")
			}
			continue
		}
		if _, ok := ConfigSettings[key]; ok {
			ConfigSettings[key] = envVal
		}
	}
	log.Info("All config & secrets set")
}
