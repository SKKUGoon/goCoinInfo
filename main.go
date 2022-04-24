package main

import (
	"fmt"
	"goBinance/crawler"
	"time"
)

func syncUpbit() {
	// signal chan []string
	for {
		a, err := crawler.CrawlUpbit(true)
		if err == nil {
			for _, orderSheet := range a {
				fmt.Println(crawler.OrderUpbit(orderSheet))
			}
		}

		time.Sleep(2 * time.Second)
	}
}

func allCrawl(coinChannel chan []string) {
	fmt.Println("All Crawl")

}

func main() {
	//syncUpbit()
	_, err := crawler.CrawlBithumb(true)
	if err != nil {
		fmt.Println(err)
	}
}
