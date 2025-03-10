package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	noteHandler "noteBackendApp/api/handlers"
	generatedApi "noteBackendApp/api/spec"
	noteConfig "noteBackendApp/configs"
	noteDaoPorts "noteBackendApp/internal/dao/impl"
	noteInterface "noteBackendApp/internal/dao/interface"
	"os"
)

func LoadConfig(env string) *noteConfig.AppConfig {
	var appProfile = "configs/" + "%s" + ".yaml"
	getenv := os.Getenv(env)
	switch getenv {
	case "dev":
		appProfile = fmt.Sprintf(appProfile, "dev")
	case "test":
		appProfile = fmt.Sprintf(appProfile, "test")
	case "docker":
		appProfile = fmt.Sprintf(appProfile, "docker")
	}
	log.Println(fmt.Sprintf("Run application in mode : %s", getenv))
	f, err := os.Open(appProfile)
	if err != nil {
	}
	defer f.Close()

	var cfg noteConfig.AppConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {

	}

	return &cfg
}

func main() {
	config := LoadConfig("MODE")
	startMessage := "Note backend ver 1.0"
	fmt.Printf("%s!\n", startMessage)
	router, apiInterface := initAPI(createDAO(config))
	router.Use(commonMiddleware)
	generatedApi.RegisterHandlers(router, apiInterface)
	log.Println(fmt.Sprintf("Starting server on : %s", config.Server.Port))
	if err := router.Run(":" + config.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func commonMiddleware(context *gin.Context) {
	context.Header("Content-Type", "application/json")
}

func initAPI(dao noteInterface.NoteDao) (router *gin.Engine, serverInterface generatedApi.ServerInterface) {
	return gin.Default(), noteHandler.NewNoteBackendApi(dao)
}

func createDAO(config *noteConfig.AppConfig) noteInterface.NoteDao {
	var dao noteInterface.NoteDao
	if config.Database.Impl == "sqlx" {
		dao = initSqlxPort(config)
	} else if config.Database.Impl == "orm" {
		dao = initOrmPort(config)
	} else {
		//dao = initInMemoryPort()
	}

	return dao
}

func initOrmPort(config *noteConfig.AppConfig) *noteDaoPorts.OrmNotePort {
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
	return noteDaoPorts.CreateOrmNotePort(db)
}

func initSqlxPort(config *noteConfig.AppConfig) *noteDaoPorts.SqlxAuthPort {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DataBaseName,
	)

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return noteDaoPorts.CreateSqlxAuthPort(db)
}

func initInMemoryPort() *noteDaoPorts.InMemoryPort {
	return noteDaoPorts.NewInMemoryPort()
}
