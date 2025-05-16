package main

import (
	"expensia/api/rest"
	"expensia/configs"
	"expensia/internal/app/ports/out/dao"
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

	//инициализирую апи
	router, apiInterface := gin.Default(), handler.NewExpensiaApi(dao)

	//устанавливаю роут под swagger ui
	rest.Load(router)

	//регаю хэндлеры
	rest.RegisterHandlers(router, apiInterface)

	//старт сервера
	if err := router.Run(":" + strconv.Itoa(appConfig.Server.Port)); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
