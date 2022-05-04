package domestic

import (
	"goBinance/orderbook"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func urlBithumb(category, pageNo string, testMode bool) string {
	/*
		/ make Bitthumb crawling url.
		/ include body, and headers for access.
	*/
	var target string
	if testMode == true {
		target = BithumbURLTEST
	} else {
		target = BithumbURL
	}

	// Create Request URL
	resp, err := http.NewRequest("GET", target, nil)
	if err != nil {
		log.Println(BithumbURLErr)
	}

	// Add Header and Body(Param)
	resp.Header = http.Header{
		"User-Agent": []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64)"},
	}
	qry := resp.URL.Query()
	qry.Add("noticeCategory", category)
	qry.Add("pageNumber", pageNo)

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

func CrawlBithumb(testMode bool) ([]BithumbTitle, error) {
	var URL = urlBithumb("9", "0", testMode)
	resp, err := http.Get(URL)
	if err != nil {
		log.Println(BithumbReqErr)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(BithumbParseErr)
		return nil, err
	} else {
		d := string(data)
		w := []string{"a", "td"}
		p, err := AssetBithumb(d, setHashMap(w))

		if err != nil {
			return nil, err
		} else {
			return p, nil
		}
	}
}

func OrderBithumb(asset string) (orderbook.OrderContent, orderbook.OrderContent) {
	/*
		/ fill in the orderContent
	*/
	orderHfreq := orderbook.OrderContent{
		A:  asset,
		N:  BithumbOrderHF,
		I:  BithumbOrderHFId,
		T:  time.Now(),
		ET: time.Now(),
		TY: BithumbAssetType,
		B:  "binance",
		BC: 01,
		OD: orderbook.OrderDetail{
			P: "market",
			D: 10 * time.Second,
		},
	}

	orderLfreq := orderbook.OrderContent{
		A:  asset,
		N:  BithumbOrderLF,
		I:  BithumbOrderLFId,
		T:  time.Now(),
		TY: BithumbAssetType,
		B:  "binance",
		BC: 01,
		OD: orderbook.OrderDetail{
			P: "limit",
			D: 60 * time.Minute,
		},
	}

	return orderHfreq, orderLfreq
}
