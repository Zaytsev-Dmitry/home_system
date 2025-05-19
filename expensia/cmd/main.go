package main

import (
	"expensia/api/openapi"
	"expensia/configs"
	"expensia/internal/app/ports/out/dao"
	"expensia/internal/app/prepare"
	"expensia/internal/infrastructure/transport/http/handler"
	"github.com/Zaytsev-Dmitry/configkit"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"strconv"
)

func main() {
	//гружу конфиг
	appConfig, _ := configkit.LoadConfig[configs.AppConfig]("configs/config-local.yaml")

	//создаю dao
	dao, db := dao.Create(appConfig)
	defer db.Close()

	//инициализирую prepareRegistry
	registry := prepare.NewPrepareRegistry()
	//инициализирую апи
	router, apiInterface := gin.Default(), handler.NewExpensiaApi(dao, registry)

	//устанавливаю роут под swagger ui
	openapi.Load(router)

	//регаю хэндлеры
	openapi.RegisterHandlers(router, apiInterface)

	//старт сервера
	if err := router.Run(":" + strconv.Itoa(appConfig.Server.Port)); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
