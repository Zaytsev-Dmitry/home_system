package noteConfigs

type AppConfig struct {
	Keycloak struct {
		KeycloakUrl     string `yaml:"keycloakUrl"`
		KeycloakHost    string `yaml:"keycloakHost"`
		KeycloakRealm   string `yaml:"KeycloakRealm"`
		TokenUrl        string `yaml:"tokenUrl"`
		ClientId        string `yaml:"clientId"`
		ClientSecret    string `yaml:"clientSecret"`
		ServerGrantType string `yaml:"serverGrantType"`
	}
	Database struct {
		Host         string `yaml:"host"`
		Username     string `yaml:"user"`
		Password     string `yaml:"password"`
		DataBaseName string `yaml:"dataBaseName"`
		Dialect      string `yaml:"dialect"`
		Impl         string `yaml:"impl"`
	} `yaml:"database"`
	Server struct {
		Name string `yaml:"name"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
}
