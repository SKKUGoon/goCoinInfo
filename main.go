package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Trading Session")
	SetupRoutes()
	log.Fatal(http.ListenAndServe(":7890", nil))
}
