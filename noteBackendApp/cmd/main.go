package main

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"gopkg.in/yaml.v3"
	"log"
	api "noteBackendApp/api/docs"
	noteHandler "noteBackendApp/api/handlers"
	noteConfig "noteBackendApp/configs"
	noteDao "noteBackendApp/internal/dao"
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
	migrateDB(config)
	startMessage := "Note backend ver 1.0"
	fmt.Printf("%s!\n", startMessage)
	router, apiInterface := initAPI()
	router.Use(commonMiddleware)
	api.RegisterHandlers(router, apiInterface)
	log.Println(fmt.Sprintf("Starting server on : %s", config.Server.Port))
	if err := router.Run(":" + config.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func commonMiddleware(context *gin.Context) {
	context.Header("Content-Type", "application/json")
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func migrateDB(config *noteConfig.AppConfig) {
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

func initAPI() (router *gin.Engine, serverInterface api.ServerInterface) {
	noteRepository := noteDao.NewInMemoryNoteRepository()
	return gin.Default(), noteHandler.NewNoteBackendApi(noteRepository)
}
