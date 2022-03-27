package crawler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func urlBithumb(cat, pg string) string {
	// make bitthumb crawling url.
	// include body, and headers for access.
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

func setHashMap(ls []string) map[string]bool {
	result := make(map[string]bool)
	for _, v := range ls {
		result[v] = true
	}
	return result
}

func CrawlBithumb() {
	var URL = urlBithumb("9", "0")
	resp, err := http.Get(URL)
	if err != nil {
		log.Println("[Crawler][Bithumb] >>> Unsuccessful request")
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("fail to read html")
		return
	} else {
		d := string(data)
		w := []string{"a", "td"}
		p, err := AssetBithumb(d, setHashMap(w))
		if err == nil {
			fmt.Println(p)
		}
	}
}
