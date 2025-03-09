package config

type AppConfig struct {
	Server struct {
		BotToken       string   `yaml:"botToken"`
		CommandsToInit []string `yaml:"commandsToInit"`
		AuthServerUrl  string   `yaml:"authServerUrl"`
		NoteBackendUrl string   `yaml:"noteBackendUrl"`
	} `yaml:"server"`
	Database struct {
		Host         string `yaml:"host"`
		Username     string `yaml:"user"`
		Password     string `yaml:"password"`
		DataBaseName string `yaml:"dataBaseName"`
		Dialect      string `yaml:"dialect"`
		Impl         string `yaml:"impl"`
	} `yaml:"database"`
}
