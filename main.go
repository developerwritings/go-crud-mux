package main

import (
	"go-crud/routes"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("To Do App")
	loadConfig := routes.LoadConfig()
	log.Println("Configuring the values of server values")
	server := &http.Server{
		Addr:         ":8081",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      loadConfig,
	}
	log.Println("Routing Initiated")
	log.Fatal(server.ListenAndServe())
}
