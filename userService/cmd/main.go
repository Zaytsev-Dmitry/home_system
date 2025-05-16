package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"strconv"
	"userService/api/http"
	openapi "userService/api/http"
	"userService/internal/app/ports/out/dao"
	"userService/internal/app/ports/out/keycloak"
	"userService/internal/infrastructure/transport/http/handler"
	"userService/internal/infrastructure/transport/http/middleware"
	"userService/pkg/config_loader"
)

func main() {
	//гружу конфиг
	appConfig := config_loader.LoadConfig()

	//делаю логер
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	//создаю DAO
	dao, db := dao.Create(appConfig)
	defer db.Close()

	//создаю keycloak client
	keycloakClient := keycloak.NewKeycloakClient(appConfig)

	//инициализирую апи
	router, apiInterface := gin.Default(), handler.NewAuthServerApi(keycloakClient, dao)

	//устанавливаю middlewares
	router.Use(middleware.TokenIntrospectionMiddleware(keycloakClient, appConfig))
	router.Use(middleware.TraceMiddleware(logger))
	router.Use(middleware.LogParamsAndResponseMiddleware(logger))

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
