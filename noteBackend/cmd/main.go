package main

import (
	"fmt"
	noteSpec "github.com/Zaytsev-Dmitry/home_system_open_api/noteServerBackend"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	noteHandler "noteBackendApp/api/handlers"
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
	noteSpec.RegisterHandlers(router, apiInterface)
	log.Println(fmt.Sprintf("Starting server on : %s", config.Server.Port))
	if err := router.Run(":" + config.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func commonMiddleware(context *gin.Context) {
	context.Header("Content-Type", "application/json")
}

func initAPI(dao noteInterface.NoteDao) (router *gin.Engine, serverInterface noteSpec.ServerInterface) {
	return gin.Default(), noteHandler.NewNoteBackendApi(dao)
}

func initPostgresPort(config *noteConfig.AppConfig) *noteDaoPorts.PostgresNotePort {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.DataBaseName,
	)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return noteDaoPorts.CreatePostgresNotePort(db)
}

func createDAO(config *noteConfig.AppConfig) noteInterface.NoteDao {
	port := initPostgresPort(config)
	var dao noteInterface.NoteDao
	dao = port
	return dao
}
