package main

import (
	"authServer/api/handlers"
	generatedApi "authServer/api/spec"
	"authServer/configs"
	"authServer/internal/dao"
	"authServer/pkg/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"strconv"
)

func main() {
	//гружу конфиг
	appConfig := getConfig()

	//делаю логер
	logger := getLogger()
	defer logger.Sync()

	//создаю DAO
	dao := getDao(appConfig)
	defer dao.Close()

	//инициализирую апи
	router, apiInterface := gin.Default(), handlers.NewAuthServerApi(appConfig, dao)

	//регаю хэндлеры
	generatedApi.RegisterHandlers(router, apiInterface)

	logger.Info("Start application",
		zap.String("name", appConfig.Server.Name),
		zap.Int("port", appConfig.Server.Port),
	)
	//старт сервера
	if err := router.Run(":" + strconv.Itoa(appConfig.Server.Port)); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}

func getDao(appConfig *configs.AppConfig) *dao.AuthDao {
	dao := dao.New(appConfig)
	if dao == nil {
		panic("dao is nil")
	}
	return dao
}

func getLogger() *zap.Logger {
	logger := utilities.GetLogger()
	if logger == nil {
		panic("logger is nil")
	}
	return logger
}

func getConfig() *configs.AppConfig {
	appConfig := configs.LoadConfig()

	fmt.Println(fmt.Sprintf("Keycloak URL: %s", appConfig.Keycloak.KeycloakUrl))
	fmt.Println(fmt.Sprintf("Keycloak HOST: %s", appConfig.Keycloak.KeycloakHost))

	if appConfig == nil {
		panic("appConfig is nil")
	}
	return appConfig
}
