package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	noteHandler "noteBackendApp/api/handlers"
	openapi "noteBackendApp/api/http"
	"noteBackendApp/internal/dao/sqlx"
	"noteBackendApp/pkg/config_loader"
)

func main() {
	config := config_loader.LoadConfig()
	startMessage := "Note backend ver 1.0"
	fmt.Printf("%s!\n", startMessage)
	router, apiInterface := gin.Default(), noteHandler.NewNoteBackendApi(sqlx.CreateSqlxPort(config))

	//устанавливаю роут под swagger ui
	openapi.Load(router)

	openapi.RegisterHandlers(router, apiInterface)
	log.Println(fmt.Sprintf("Starting server on : %s", config.Server.Port))
	if err := router.Run(":" + config.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
