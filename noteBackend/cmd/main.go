package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	openapi "noteBackendApp/api/http"
	"noteBackendApp/internal/app/ports/out/dao"
	noteHandler "noteBackendApp/internal/infrastructure/transport/http/handler"
	"noteBackendApp/pkg/config_loader"
)

func main() {
	config := config_loader.LoadConfig()
	startMessage := "Note backend ver 1.0"
	fmt.Printf("%s!\n", startMessage)

	dao, db := dao.Create(config)
	defer db.Close()

	router, apiInterface := gin.Default(), noteHandler.NewNoteBackendApi(dao)

	//устанавливаю роут под swagger ui
	openapi.Load(router)

	openapi.RegisterHandlers(router, apiInterface)
	log.Println(fmt.Sprintf("Starting server on : %s", config.Server.Port))
	if err := router.Run(":" + config.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
