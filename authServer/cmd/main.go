package main

import (
	api "authServer/api/docs"
	authServerHandler "authServer/api/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	startMessage := "AuthServer ver 1.0"
	fmt.Printf("%s!\n", startMessage)
	router, apiInterface := initAPI()
	api.RegisterHandlers(router, apiInterface)
	log.Println("Starting server on :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initAPI() (router *gin.Engine, serverInterface api.ServerInterface) {
	return gin.Default(), authServerHandler.NewAuthServerApi()
}
