package main

import (
	authServerHandler "authServer/api/handlers"
	authConfig "authServer/configs"
	daoImpl "authServer/internal/dao"
	"fmt"
	authSpec "github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strconv"
)

func loadConfig(env string) *authConfig.AppConfig {
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

	var cfg authConfig.AppConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {

	}
	return &cfg
}

func main() {
	//гружу конфиг
	appConfig := loadConfig("MODE")

	//создаю DAO
	dao := daoImpl.CreateDao(*appConfig)

	//инициализирую апи
	router, apiInterface := gin.Default(), authServerHandler.NewAuthServerApi(appConfig, *dao)

	//регаю хэндлеры
	authSpec.RegisterHandlers(router, apiInterface)

	log.Printf("Start application: %s on port %x", appConfig.Server.Name, appConfig.Server.Port)

	//старт сервера
	if err := router.Run(":" + strconv.Itoa(appConfig.Server.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
