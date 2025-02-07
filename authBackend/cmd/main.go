package main

import (
	"authServer/api/handlers"
	"authServer/configs"
	"authServer/internal/dao"
	"authServer/pkg/utilities"
	"fmt"
	"github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"strconv"
)

func main() {
	//делаю логер
	logger := utilities.GetLogger()
	defer logger.Sync()

	//гружу конфиг
	appConfig := configs.LoadConfig("MODE")

	//создаю DAO
	dao := dao.CreateDao(*appConfig)
	defer dao.Close()

	//инициализирую апи
	router, apiInterface := gin.Default(), handlers.NewAuthServerApi(appConfig, *dao)

	//регаю хэндлеры
	authSpec.RegisterHandlers(router, apiInterface)

	logger.Info(fmt.Sprintf("Start application: %s on port %x", appConfig.Server.Name, appConfig.Server.Port))

	//старт сервера
	if err := router.Run(":" + strconv.Itoa(appConfig.Server.Port)); err != nil {
		logger.Error("Failed to start server: " + err.Error())
	}
}
