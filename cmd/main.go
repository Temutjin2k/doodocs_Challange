package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Temutjin2k/doodocs_Challange/internal/config"
	"github.com/Temutjin2k/doodocs_Challange/internal/server"
)

func main() {
	portFlag := flag.Int("port", 8080, "Port number")
	flag.Parse()

	port := fmt.Sprintf(":%d", *portFlag)

	config.LoadEnvVariables()
	router := server.InitServer()
	log.Println("Starting server on port" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
