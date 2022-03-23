package main

import (
	//"goBinance/broadcast"
	"goBinance/crawler"
	"log"
	//"net/http"
)

func main() {
	log.Println("Starting Trading Session:")
	//broadcast.SetupRoutes()
	crawler.CrawlUpbit()
	//log.Fatal(http.ListenAndServe(":7890", nil))
}
