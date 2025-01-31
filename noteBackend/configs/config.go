package noteConfigs

type AppConfig struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Url          string `yaml:"url"`
		Host         string `yaml:"host"`
		Username     string `yaml:"user"`
		Password     string `yaml:"password"`
		DataBaseName string `yaml:"dataBaseName"`
		Dialect      string `yaml:"dialect"`
		Impl         string `yaml:"impl"`
	} `yaml:"database"`
}
