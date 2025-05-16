package configs

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
