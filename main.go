package main

import (
	"fmt"
	"goBinance/crawler"
	"log"
	"time"
)

func ccrUpbit(signal chan []string) {

}

func allCrawl(coinChannel chan []string) {
	fmt.Println("All Crawl")

}

func main() {
	for {
		a, err := crawler.CrawlUpbit(true)
		if err == nil {
			log.Println("asset", a)
		}
		time.Sleep(time.Second * 2)
	}

	//crawler.CrawlBithumb()
}
