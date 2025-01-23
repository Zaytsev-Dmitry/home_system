package main

import (
	api "authServer/api/docs"
	authServerHandler "authServer/api/handlers"
	authConfig "authServer/configs"
	"authServer/external"
	"database/sql"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func loadConfig(env string) *authConfig.AppConfig {
	var appProfile = "configs/" + "%s" + ".yaml"
	getenv := os.Getenv(env)
	switch getenv {
	case "dev":
		appProfile = fmt.Sprintf(appProfile, "dev")
	case "test":
		appProfile = fmt.Sprintf(appProfile, "test")
	}
	log.Println(fmt.Sprintf("Run application in mode : %s", getenv))
	f, err := os.Open(appProfile)
	if err != nil {
	}
	defer f.Close()

	var cfg authConfig.AppConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {

	}
	return &cfg
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func migrateDB(config *authConfig.AppConfig) {
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s?sslmode=disable",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.DataBaseName,
	)
	db, err := sql.Open(config.Database.Dialect, dataSourceName)
	if err != nil {
		panic(err)
	}
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect(config.Database.Dialect); err != nil {
		panic(err)
	}
	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
}

func main() {
	startMessage := "AuthServer ver 1.0"
	fmt.Printf("%s!\n", startMessage)
	appConfig := loadConfig("MODE")
	router, apiInterface := initAPI(appConfig)
	migrateDB(appConfig)
	api.RegisterHandlers(router, apiInterface)
	log.Println("Starting server on :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initAPI(config *authConfig.AppConfig) (router *gin.Engine, serverInterface api.ServerInterface) {
	client := getKeycloakClient(config)
	return gin.Default(), authServerHandler.NewAuthServerApi(client)
}

func getKeycloakClient(config *authConfig.AppConfig) external.KeycloakClient {
	client := external.KeycloakClient{
		KeycloakUrl:     config.Keycloak.KeycloakUrl,
		TokenUrl:        config.Keycloak.TokenUrl,
		ClientId:        config.Keycloak.ClientId,
		ClientSecret:    config.Keycloak.ClientSecret,
		ServerGrantType: config.Keycloak.ServerGrantType,
	}
	return client
}
