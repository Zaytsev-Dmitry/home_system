package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	noteHandler "noteBackendApp/api/handlers"
	api "noteBackendApp/api/openapi"
	noteDao "noteBackendApp/internal/dao"
)

func main() {
	startMessage := "Home system ver 1.0"
	fmt.Printf("%s!\n", startMessage)
	noteRepository := noteDao.NewInMemoryNoteRepository()
	router := gin.Default()
	noteApi := noteHandler.NewNoteApi(noteRepository)
	api.RegisterHandlers(router, noteApi)

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
