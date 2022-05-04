package main

import (
	"fmt"
	"goBinance/crawler"
	"goBinance/orderbook"
	"log"
	"math/rand"
	"os"
	"time"
)

func syncUpbit(lowFreqSig, highFreqSig chan orderbook.OrderContent) {
	// lowFreqSig chan []string
	for {
		a, err := crawler.CrawlUpbit(true)
		if err != nil {
			log.Println(err)
		} else {
			for _, orderSheet := range a {
				h, l := crawler.OrderUpbit(orderSheet)
				// Insert it in a channel - High Frequency
				highFreqSig <- h
				// Insert it in a channel - Low Frequency
				lowFreqSig <- l

			}
		}
		time.Sleep(2 * time.Second)
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

func syncBithumb(lowFreqSig, highFreqSig chan orderbook.OrderContent, waitTime int) {
	rand.Seed(time.Now().UnixNano())
	for {
		a, err := crawler.CrawlBithumb(false)
		if err != nil {
			fmt.Println(err)
		} else {
			for _, asset := range a {
				if recentBithumb(asset, waitTime) {
					h, l := orderBithumb(asset)
					// Insert it in a channel - High Frequency
					for _, orders := range h {
						highFreqSig <- orders
					}
					// insert it in a channel - Low Frequency
					for _, orders := range l {
						lowFreqSig <- orders
					}
				}
			}
		}
		//time.Sleep(time.Duration(waitTime) * time.Second)
		min, max := 60, 150
		n := rand.Intn(max-min+1) + min
		log.Printf("syncBithumb sleeping %d seconds\n", n)
		time.Sleep(time.Duration(n) * time.Second)
	}
}

func serverEx(e chan string) {
	/*
		/ Exit crawlers after 1 day (UTC standard)
		/ 	by sending signal to exit channel(e)
	*/
	tNow := time.Now().Minute()
	for {
		if time.Now().Minute() != tNow {
			e <- "Date Change. Restart Crawler"
		}
	}
}

func main() {
	hfSigChan := make(chan orderbook.OrderContent)
	lfSigChan := make(chan orderbook.OrderContent)
	exit := make(chan string)

	go syncUpbit(hfSigChan, lfSigChan)
	go syncBithumb(lfSigChan, hfSigChan, 60*60*24*5)
	go serverEx(exit)

	for {
		select {
		case trade := <-hfSigChan:
			// High Frequency Trading Signals
			fmt.Println("hrec", trade)
		case trade := <-lfSigChan:
			// Low Frequency Trading Signals
			fmt.Println("lrec", trade)
		case msg := <-exit:
			// Exit Signal
			log.Println(msg)
			os.Exit(0)
		}
	}

}
