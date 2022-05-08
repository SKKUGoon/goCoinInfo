package main

import (
	"fmt"
	"goBinance/domestic"
	"goBinance/orderbook"
	"log"
	"math/rand"
	"os"
	"time"
)

func syncUpbit(lowFreqSig, highFreqSig chan orderbook.OrderContent) {
	// lowFreqSig chan []string
	for {
		a, err := domestic.CrawlUpbit(false)
		if err != nil {
			log.Println(err)
		} else {
			for _, asset := range a {
				orderStrat1 := orderbook.Strategy1(asset, 2, 1)
				orderStrat2 := orderbook.Strategy2(asset, 2, 1)
				orderStrat3 := orderbook.Strategy3(asset, 2, 1)

				// Insert it in a channel - High Frequency
				highFreqSig <- orderStrat1
				// Insert it in a channel - Low Frequency
				lowFreqSig <- orderStrat2
				lowFreqSig <- orderStrat3
			}
		}
		time.Sleep(2 * time.Second)
	}
}

func syncBithumb(lowFreqSig, highFreqSig chan orderbook.OrderContent, waitTime int) {
	rand.Seed(time.Now().UnixNano())
	for {
		a, err := domestic.CrawlBithumb(false)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, asset := range a {
			if domestic.RecentBithumb(asset, waitTime) {
				h, l := domestic.OrderBithumb(asset)

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
		// Sleep random time to avoid detection
		domestic.RandomSleep()
	}
}

func crawlerEx(e chan string) {
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
	go crawlerEx(exit)

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
