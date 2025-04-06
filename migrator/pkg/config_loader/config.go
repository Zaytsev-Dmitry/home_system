package config_loader

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type DatabaseConfig struct {
	Name          string `yaml:"name" env:"DB_NAME" env-required:"true"`
	User          string `yaml:"user" env:"DB_USER" env-required:"true"`
	Password      string `yaml:"password" env:"DB_PASSWORD" env-required:"true"`
	DirectoryName string `yaml:"directoryName"`
}

type Config struct {
	App struct {
		DbUrl string `yaml:"dbUrl" env:"DB_URL" env-required:"true"`
	} `yaml:"app"`
	Databases []DatabaseConfig `yaml:"databases"`
}

func LoadConfig() *Config {
	configPath := os.Getenv("CONFIG_FILE")

	if configPath == "" {
		log.Fatalf("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("CONFIG_PATH does not exist: %s", configPath)
	}

	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("cannot read configs: %s", err)
	}
	return &config
}
