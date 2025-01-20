package main

import (
	"fmt"
	"log"
	"net/http"
	noteApi "noteBackendApp/api"
)

func main() {
	startMessage := "Home system ver 1.0"
	fmt.Printf("Start application, %s!\n", startMessage)
	log.Fatal(http.ListenAndServe(":8080", noteApi.Init()))
}
