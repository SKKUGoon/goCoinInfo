package crawler

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

func parseTable(text string, wanted []string) ([]string, error) {
	content := []string{}

	tkn := html.NewTokenizer(strings.NewReader(text))
	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			return nil, errors.New("end of parse")
		case tt == html.StartTagToken:
			t := tkn.Token()
			for _, name := range wanted {
				if t.Data == name {
					text := string(tkn.Text())
					t := strings.TrimSpace(text)
					content = append(content, t)
				}
			}
		}
		fmt.Println(content)
	}

}
func CrawlBithumb() {
	var URL = urlBithumb("9", "0")
	resp, err := http.Get(URL)
	if err != nil {
		log.Println("[Crawler][Bithumb] >>> Unsuccessful")
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("fail to read html")
		return
	} else {
		d := string(data)
		w := []string{"td", "table"}
		p, err := parseTable(d, w)
		if err == nil {
			fmt.Println(p)
		}
	}

}
