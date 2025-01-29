package config

type AppConfig struct {
	BotToken       string   `yaml:"botToken"`
	HandlersToInit []string `yaml:"handlersToInit"`
	AuthServerUrl  string   `yaml:"authServerUrl"`
}
