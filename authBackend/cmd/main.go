package main

import (
	"authServer/api/handlers"
	"authServer/api/http"
	swaggerGenerator "authServer/api/http"
	"authServer/internal/dao"
	"authServer/pkg/config_loader"
	"authServer/pkg/utilities"
	_ "embed"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"strconv"
)

func main() {
	//гружу конфиг
	appConfig := config_loader.LoadConfig()

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

func getDao(appConfig *config_loader.AppConfig) *dao.AuthDao {
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
