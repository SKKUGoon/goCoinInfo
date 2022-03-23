package main

import (
	"fmt"
	"goBinance/crawler"
)

func main() {
	a, err := crawler.CrawlUpbit()
	if err == nil {
		fmt.Println(a)
	}

	crawler.CrawlBithumb()
}
