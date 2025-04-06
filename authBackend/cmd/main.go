package main

import (
	"authServer/api/handlers"
	"authServer/api/http"
	swaggerGenerator "authServer/api/http"
	"authServer/configs"
	"authServer/internal/dao"
	"authServer/pkg/utilities"
	_ "embed"
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
	//устанавливаю роут под swagger ui
	swaggerGenerator.Load(router)
	//регаю хэндлеры
	http.RegisterHandlers(router, apiInterface)

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
	fmt.Printf("%+v\n", appConfig)
	if appConfig == nil {
		panic("appConfig is nil")
	}
	return appConfig
}
