package main

import (
	"authServer/api/handlers"
	generatedApi "authServer/api/spec"
	"authServer/configs"
	"authServer/internal/dao"
	"authServer/pkg/utilities"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"strconv"
)

func main() {
	//гружу конфиг
	appConfig := configs.LoadConfig("MODE")

	if appConfig == nil {
		panic("appConfig is nil")
	}
	//делаю логер
	logger := utilities.GetLogger()
	if logger == nil {
		panic("logger is nil")
	}
	defer logger.Sync()

	//создаю DAO
	dao := dao.CreateDao(*appConfig)
	if dao == nil {
		panic("dao is nil")
	}
	defer dao.Close()

	//инициализирую апи
	router, apiInterface := gin.Default(), handlers.NewAuthServerApi(appConfig, *dao)

	//регаю хэндлеры
	generatedApi.RegisterHandlers(router, apiInterface)

	logger.Info("Start application",
		zap.String("name", appConfig.Server.Name),
		zap.Int("port", appConfig.Server.Port),
	)
	//старт сервера
	if err := router.Run(":" + strconv.Itoa(appConfig.Server.Port)); err != nil {
		logger.Error("Failed to start server: " + err.Error())
	}
}
