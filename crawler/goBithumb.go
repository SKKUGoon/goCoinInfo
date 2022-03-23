package crawler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func urlBithumb(cat, pg string) string {
	const URL = "https://cafe.bithumb.com/view/boards/43"

	resp, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Println("[Crawler][Bithumb] >>> URL creation unsuccessful")
	}

	resp.Header = http.Header{
		"User-Agent": []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64)"},
	}
	qry := resp.URL.Query()
	qry.Add("noticeCategory", cat)
	qry.Add("pageNumber", pg)

	resp.URL.RawQuery = qry.Encode()
	return resp.URL.String()
}

func CrawlBithumb() {
	var URL = urlBithumb("9", "0")
	resp, err := http.Get(URL)
	if err != nil {
		log.Println("[Crawler][Bithumb] >>> Unsuccessful")
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		dataString := string(data)
		fmt.Println(dataString)
	}
}
