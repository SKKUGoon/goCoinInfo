package main

import (
	"fmt"
	"goBinance/crawler"
	"goBinance/orderbook"
	"log"
	"time"
)

func syncUpbit() {
	// signal chan []string
	for {
		a, err := crawler.CrawlUpbit(true)
		if err != nil {
			log.Println(err)
		} else {
			for _, orderSheet := range a {
				h, f := crawler.OrderUpbit(orderSheet)
				fmt.Println(h)
				fmt.Println()
				fmt.Println(f)
				fmt.Println()
			}
		}

		time.Sleep(2 * time.Second)
		break
	}
}

func recentBithumb(post crawler.BithumbTitle, secondSlack int) bool {
	t := time.Now().Add(time.Duration(secondSlack*-1) * time.Second)
	if t.Before(post.CreatedAt) {
		return true
	} else {
		return false
	}
}

func orderBithumb(post crawler.BithumbTitle) ([]orderbook.OrderContent, []orderbook.OrderContent) {
	var highFreq []orderbook.OrderContent
	var lowFreq []orderbook.OrderContent
	for _, a := range post.Asset {
		hf, lf := crawler.OrderBithumb(a)
		highFreq = append(highFreq, hf)
		lowFreq = append(lowFreq, lf)
	}
	return highFreq, lowFreq
}

func syncBithumb(waitTime int) {
	// signal chan []string
	for {
		a, err := crawler.CrawlBithumb(true)
		if err != nil {
			fmt.Println(err)
		} else {
			for _, asset := range a {
				if recentBithumb(asset, waitTime) {
					h, l := orderBithumb(asset)
					fmt.Println(h)
					fmt.Println()
					fmt.Println(l)
					fmt.Println()
				}
			}
		}
		//time.Sleep(time.Duration(waitTime) * time.Second)
		break
	}
}

//func allCrawl(coinChannel chan []string) {
//	fmt.Println("All Crawl")
//
//}

func main() {
	syncUpbit()
	syncBithumb(60 * 60 * 24 * 4)
}
