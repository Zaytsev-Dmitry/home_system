package config_loader

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type AppConfig struct {
	Server struct {
		Port int    `yaml:"port"`
		Name string `yaml:"name"`
	} `yaml:"server"`
	Keycloak struct {
		KeycloakUrl     string `yaml:"keycloakUrl",env:"KEYCLOAK_URL"`
		KeycloakHost    string `yaml:"keycloakHost"`
		KeycloakRealm   string `yaml:"keycloakRealm"`
		TokenUrl        string `yaml:"tokenUrl"`
		ClientId        string `yaml:"clientId"`
		ClientSecret    string `yaml:"clientSecret"`
		ServerGrantType string `yaml:"serverGrantType"`
	}
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
	configPath := os.Getenv("CONFIG_PATH")
	profile := os.Getenv("APP_PROFILE")

	if configPath == "" && profile == "local" {
		configPath = "configs/" + "config-local.yaml"
	}

	if configPath == "" {
		log.Fatalf("CONFIG_PATH environment variable not set")
	}

	if profile == "" {
		log.Fatalf("APP_PROFILE environment variable not set")
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
