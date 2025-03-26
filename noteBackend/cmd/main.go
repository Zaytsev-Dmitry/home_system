package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	noteHandler "noteBackendApp/api/handlers"
	generatedApi "noteBackendApp/api/spec"
	noteConfig "noteBackendApp/config"
	noteDaoPorts "noteBackendApp/internal/dao/impl"
	"os"
)

func main() {
	config := noteConfig.LoadConfig()
	startMessage := "Note backend ver 1.0"
	fmt.Printf("%s!\n", startMessage)
	router, apiInterface := gin.Default(), noteHandler.NewNoteBackendApi(initSqlxPort(config))

	//устанавливаю роут под swagger ui
	generatedApi.Load(router)

	generatedApi.RegisterHandlers(router, apiInterface)
	log.Println(fmt.Sprintf("Starting server on : %s", config.Server.Port))
	if err := router.Run(":" + config.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initSqlxPort(config *noteConfig.AppConfig) *noteDaoPorts.SqlxAuthPort {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DataBaseName,
	)

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return noteDaoPorts.CreateSqlxAuthPort(db)
}
