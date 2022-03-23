package main

import (
	"fmt"
	//"goBinance/broadcast"
	"goBinance/crawler"
	"log"
	//"net/http"
)

func main() {
	log.Println("Starting Trading Session:")
	//broadcast.SetupRoutes()
	a, err := crawler.CrawlUpbit()
	if err == nil {
		fmt.Println(a)
	}
	//log.Fatal(http.ListenAndServe(":7890", nil))
}
