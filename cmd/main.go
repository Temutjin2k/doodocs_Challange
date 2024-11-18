package main

import (
	"log"
	"net/http"

	"github.com/Temutjin2k/doodocs_Challange/internal/config"
	"github.com/Temutjin2k/doodocs_Challange/internal/server"
)

func main() {
	config.LoadEnvVariables()
	router := server.InitServer()
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
