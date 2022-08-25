package main

import (
	"log"

	"ts-rest-api/internal/http"
	"ts-rest-api/internal/server"
)

func main() {
	// todo starting rest api
	handler := http.NewHandler()
	srv := new(server.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %v", err.Error())
	}
}
