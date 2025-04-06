package main

import (
	"authBackend/api/http"
	openapi "authBackend/api/http"
	"authBackend/internal/app/ports/out/dao"
	"authBackend/internal/infrastructure/transport/http/handler"
	"authBackend/pkg/config_loader"
	"authBackend/pkg/utilities"
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
	dao, db := dao.Create(appConfig)
	defer db.Close()

	//инициализирую апи
	router, apiInterface := gin.Default(), handler.NewAuthServerApi(appConfig, dao)
	//устанавливаю роут под swagger ui
	openapi.Load(router)
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

func getLogger() *zap.Logger {
	logger := utilities.GetLogger()
	if logger == nil {
		panic("logger is nil")
	}
	return logger
}
