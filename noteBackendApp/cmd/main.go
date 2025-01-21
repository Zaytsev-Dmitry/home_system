package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	api "noteBackendApp/api/docs"
	noteHandler "noteBackendApp/api/handlers"
	noteDao "noteBackendApp/internal/dao"
)

func main() {
	startMessage := "Note backend ver 1.0"
	fmt.Printf("%s!\n", startMessage)
	router, apiInterface := initAPI()
	api.RegisterHandlers(router, apiInterface)
	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initAPI() (router *gin.Engine, serverInterface api.ServerInterface) {
	noteRepository := noteDao.NewInMemoryNoteRepository()
	return gin.Default(), noteHandler.NewNoteBackendApi(noteRepository)
}
