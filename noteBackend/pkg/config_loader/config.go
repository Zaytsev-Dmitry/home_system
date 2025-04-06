package config_loader

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type AppConfig struct {
	Server struct {
		Port string `yaml:"port"`
		Name string `yaml:"name"`
	} `yaml:"server"`
	Database struct {
		Host         string `yaml:"host"`
		Username     string `yaml:"userName"`
		Password     string `yaml:"password"`
		DataBaseName string `yaml:"dataBaseName"`
		Dialect      string `yaml:"dialect"`
		Port         string `yaml:"port"`
	} `yaml:"database"`
}

func LoadConfig() *AppConfig {
	profile := os.Getenv("APP_PROFILE")

	if profile == "" {
		log.Fatalf("APP_PROFILE environment variable not set")
	}

	fmt.Println("Profile:", profile)

	var configPath string
	switch profile {
	case "prod":
		configPath = "configs/prod.yaml"
	case "dev":
		configPath = "configs/dev.yaml"
	case "local":
		configPath = "configs/local.yaml"
	case "docker_local":
		configPath = "configs/docker_local.yaml"
	default:
		log.Fatalf("Unknown profile %s", profile)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("CONFIG_PATH does not exist: %s", configPath)
	}

	var config AppConfig
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("cannot read configs: %s", err)
	}
	return &config
}
