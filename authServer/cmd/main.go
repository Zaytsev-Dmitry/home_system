package main

import (
	api "authServer/api/docs"
	authServerHandler "authServer/api/handlers"
	authConfig "authServer/configs"
	"authServer/external"
	authDaoPorts "authServer/internal/dao/impl"
	authDaoInterface "authServer/internal/dao/interface"
	"context"
	"database/sql"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Connection struct {
	gormDb  *gorm.DB
	pgxConn *pgx.Conn
}

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

func initPgxPort(config *authConfig.AppConfig) *authDaoPorts.PgxAuthPort {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.DataBaseName,
	)
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return authDaoPorts.CreatePgxAuthPort(conn)
}

func createDAO(config *authConfig.AppConfig) (interf authDaoInterface.AuthDao, conn Connection) {
	var dao authDaoInterface.AuthDao
	var connDto Connection
	if config.Database.Impl == "pgx" {
		port := initPgxPort(config)
		connDto.pgxConn = port.Conn
		dao = port
	} else {
		port := initOrmPort(config)
		connDto.gormDb = port.Db
		dao = port
	}
	return dao, connDto
}

func main() {
	startMessage := "AuthServer ver 1.0"
	fmt.Printf("%s!\n", startMessage)
	//гружу конфиг
	appConfig := loadConfig("MODE")

	//мигрирую базу
	migrateDB(appConfig)

	//создаю нужный DAO на основе конфига
	dao, connDTO := createDAO(appConfig)
	if connDTO.pgxConn != nil {
		defer connDTO.pgxConn.Close(context.Background())
	} else {
		dbInstance, _ := connDTO.gormDb.DB()
		defer dbInstance.Close()
	}

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
