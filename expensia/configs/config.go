package configs

type AppConfig struct {
	Server struct {
		Port int    `yaml:"port"`
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
