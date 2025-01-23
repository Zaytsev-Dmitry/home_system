package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

func initAPI() (router *gin.Engine, serverInterface api.ServerInterface) {
	noteRepository := noteDao.NewInMemoryNoteRepository()
	return gin.Default(), noteHandler.NewNoteBackendApi(noteRepository)
}
