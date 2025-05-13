package main

import (
	"expensia/api/http"
	openapi "expensia/api/http"
	"expensia/internal/app/ports/out/dao"
	"expensia/internal/infrastructure/transport/http/handler"
	"expensia/pkg/config_loader"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"strconv"
)

func main() {
	//гружу конфиг
	appConfig := config_loader.LoadConfig()

	//создаю dao
	dao, db := dao.Create(appConfig)
	defer db.Close()

	//инициализирую апи
	router, apiInterface := gin.Default(), handler.NewExpensiaApi(dao)

	//устанавливаю роут под swagger ui
	openapi.Load(router)

	//регаю хэндлеры
	http.RegisterHandlers(router, apiInterface)

	//старт сервера
	if err := router.Run(":" + strconv.Itoa(appConfig.Server.Port)); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
