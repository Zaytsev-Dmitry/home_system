package configs

import (
	"gopkg.in/yaml.v3"
	"os"
)

func LoadConfig() *AppConfig {
	var appConfig AppConfig
	configEnv := os.Getenv("APP_CONFIG")
	if configEnv == "" {
		//TODO логер warn
		var defaultConfigName = "config/" + "default" + ".yaml"
		f, _ := os.Open(defaultConfigName)
		yaml.NewDecoder(f).Decode(&appConfig)
		defer f.Close()
	} else {
		yaml.Unmarshal([]byte(configEnv), &appConfig)
	}

	return &appConfig
}
