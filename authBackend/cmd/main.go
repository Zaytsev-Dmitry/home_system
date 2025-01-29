package main

import (
	api "authServer/api/docs"
	authServerHandler "authServer/api/handlers"
	authConfig "authServer/configs"
	"authServer/external"
	authDaoPorts "authServer/internal/dao/impl"
	authDaoInterface "authServer/internal/dao/interface"
	"database/sql"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func initOrmPort(config *authConfig.AppConfig) *authDaoPorts.OrmAuthPort {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.DataBaseName,
	)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	dbInstance, _ := db.DB()
	_ = dbInstance.Close()
	if err != nil {
		panic(err)
	}
	return authDaoPorts.CreateOrmAuthPort(db)
}

func initSqlxPort(config *authConfig.AppConfig) *authDaoPorts.SqlxAuthPort {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s?sslmode=disable",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.DataBaseName,
	)

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return authDaoPorts.CreateSqlxAuthPort(db)
}

func createDAO(config *authConfig.AppConfig) (resp authDaoInterface.AuthDao) {
	var dao authDaoInterface.AuthDao
	if config.Database.Impl == "sqlx" {
		dao = initSqlxPort(config)
	} else {
		dao = initOrmPort(config)
	}
	return dao
}

func main() {
	startMessage := "AuthServer ver 1.0"
	fmt.Printf("%s!\n", startMessage)
	//гружу конфиг
	appConfig := loadConfig("MODE")

	//мигрирую базу
	migrateDB(appConfig)

	//создаю нужный DAO на основе конфига
	dao := createDAO(appConfig)
	defer dao.CloseConnection()

	//инициализирую апи
	router, apiInterface := initAPI(appConfig, dao)
	api.RegisterHandlers(router, apiInterface)
	log.Println("Starting server on :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initAPI(config *authConfig.AppConfig, dao authDaoInterface.AuthDao) (router *gin.Engine, serverInterface api.ServerInterface) {
	client := getKeycloakClient(config)
	return gin.Default(), authServerHandler.NewAuthServerApi(client, dao)
}

func getKeycloakClient(config *authConfig.AppConfig) external.KeycloakClient {
	client := external.KeycloakClient{
		KeycloakUrl:     config.Keycloak.KeycloakUrl,
		TokenUrl:        config.Keycloak.TokenUrl,
		KeycloakHost:    config.Keycloak.KeycloakHost,
		KeycloakRealm:   config.Keycloak.KeycloakRealm,
		ClientId:        config.Keycloak.ClientId,
		ClientSecret:    config.Keycloak.ClientSecret,
		ServerGrantType: config.Keycloak.ServerGrantType,
	}
	return client
}
