package main

import (
	"fmt"
	"goBinance/crawler"
	"log"
	"time"
)

func syncUpbit() {
	// signal chan []string
	for {
		a, err := crawler.CrawlUpbit(false)
		if err != nil {
			log.Println(err)
		} else {
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
	t, err := crawler.CrawlBithumb(true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t)
	}
}
