package config_loader

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

func LoadConfig(config interface{}) {
	configPath := os.Getenv("CONFIG_FILE")
	appProfile := os.Getenv("APP_PROFILE")

	// Если профиль "local", установим локальный путь к конфигу
	if appProfile == "local" {
		configPath = "configs/" + "config-local.yaml"
	}

	// Проверим, что путь к конфигу задан
	if configPath == "" {
		log.Fatalf("CONFIG_FILE environment variable not set")
	}

	// Проверим, существует ли файл конфигурации
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("CONFIG_FILE does not exist: %s", configPath)
	}

	// Прочитаем конфигурацию в переданный интерфейс
	if err := cleanenv.ReadConfig(configPath, config); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
}
