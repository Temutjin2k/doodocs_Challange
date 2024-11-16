package main

import (
	"log"
	"net/http"

	"github.com/Temutjin2k/doodocs_Challange/internal/server"
)

func main() {
	router := server.InitServer()
	log.Fatal(http.ListenAndServe(":8080", router))
}
