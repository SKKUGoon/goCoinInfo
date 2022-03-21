package main

import (
	"goBinance/broadcast"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Trading Session:")
	broadcast.SetupRoutes()
	log.Fatal(http.ListenAndServe(":7890", nil))
}
