package configs

type AppConfig struct {
	Keycloak struct {
		KeycloakUrl     string `env:"KEYCLOAK_URL, default=http://localhost:7080/realms/home-system/protocol/openid-connect"`
		KeycloakHost    string `env:"KEYCLOAK_HOST, default=http://localhost:7080"`
		KeycloakRealm   string `env:"KEYCLOAK_REALM, default=home-system"`
		TokenUrl        string `env:"TOKEN_URL, default=/token"`
		ClientId        string `env:"CLIENT_ID, default=d3f1202c-d7d6-4694-9a6c-6e1575321baf"`
		ClientSecret    string `env:"CLIENT_SECRET, default=W2aXh7r5NrGPPROdgS1x1Ic5cMXqCZHx"`
		ServerGrantType string `env:"GRANT_TYPE, default=client_credentials"`
	}
	Database struct {
		Host         string `env:"DB_HOST, default=localhost"`
		Username     string `env:"DB_USERNAME, default=auth_user"`
		Password     string `env:"DB_PASSWORD, default=8c6b18ac-84ff-4436-8916-6e11aaa41e92"`
		DataBaseName string `env:"DB_NAME, default=auth"`
		Dialect      string `env:"DB_DIALECT, default=postgres"`
		Impl         string `env:"DB_IMPL, default=sqlx"`
	} `yaml:"database"`
	Server struct {
		Name string `env:"SERVER_NAME, default=auth backend server"`
		Port int    `env:"SERVER_PORT, default=8081"`
	} `yaml:"server"`
}
