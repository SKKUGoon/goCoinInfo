package crawler

import (
	"errors"
	"golang.org/x/net/html"
	"log"
	"strings"
	"time"
)

const (
	MARKETADDUPBIT    = "마켓 디지털 자산 추가"
	MARKETADDBITHUMB0 = "마켓 추가"
	MARKETADDBITHUMB1 = "마켓 상장"
	MARKETKRW         = "KRW"
	MARKETBTC         = "BTC"
)

func clean(s []byte) (string, error) {
	// ticker only contains alphabet character
	// returns (alphabet value, error)
	j := 0
	for _, b := range s {
		sml := 'a' <= b && b <= 'z' // small letters
		big := 'A' <= b && b <= 'Z' // capital letters
		if sml || big {
			s[j] = b
			j++
		}
	}
	if j == 0 {
		return "", errors.New("character is not alphabet")
	} else {
		return string(s[:j]), nil
	}
}

func ifAssetKor(text string) ([]string, error) {
	var titleAsset []string
	for _, c := range strings.Fields(text) {
		a, err := clean([]byte(c))
		if err == nil {
			if a != MARKETKRW && a != MARKETBTC {
				titleAsset = append(titleAsset, a)
			}
		}
	}
	if len(titleAsset) <= 0 {
		return titleAsset, errors.New("no asset found")
	} else {
		return titleAsset, nil
	}
}

func AssetUpbit(title upbitTitle) ([]string, error) {
	// find asset ticker inside upbit api signals
	t := title.Title
	if strings.Contains(t, MARKETADDUPBIT) {
		data, err := ifAssetKor(t)
		if err != nil {
			log.Println(data)
			return nil, errors.New("no asset found")
		} else {
			return data, nil
		}
	} else {
		return nil, errors.New("not asset statement")
	}
}

func extractDate(tp html.TokenType, t html.Token) (interface{}, error) {
	if tp == html.TextToken {
		layout := "2006.01.02"
		td, err := time.Parse(layout, t.Data)
		if err != nil {
			return nil, errors.New("not a date")
		} else {
			return td, nil
		}
	} else {
		return nil, errors.New("not html text token")
	}
}

func extractAsset(tp html.TokenType, t html.Token) {
	if tp == html.TextToken {
		asset, err := ifAssetKor(t.Data)
		if err != nil {
			log.Println("err", err)
		} else {
			log.Println("asset", asset)
		}
	}
}

func AssetBithumb(text string, wanted map[string]bool) ([]string, error) {
	/*
		/ bithumb does not provide news api
		/ parse the raw html files.
		/ crawl the TextToken that comes after <a> StartToken: asset name
		/ crawl the TextToken that comes after <td> StartToken: date
	*/
	tkn := html.NewTokenizer(strings.NewReader(text))
	for {
		tt := tkn.Next()
		switch {
		// End of Parsing
		case tt == html.ErrorToken:
			return nil, errors.New("end of parse")

		// ex) <a> </a> <= <a> is a StartTagToken
		case tt == html.StartTagToken:
			tn := tkn.Token()

			switch {
			// process datetime
			case wanted[tn.Data] == true && tn.Data == "td":
				tt = tkn.Next()
				tn = tkn.Token()
				td, err := extractDate(tt, tn)
				if err != nil {
					continue
				} else {
					log.Println(td)
				}

			// process asset content
			case wanted[tn.Data] == true && tn.Data == "a":
				tt = tkn.Next()
				tn = tkn.Token()
				extractAsset(tt, tn)
			default:
				continue
			}
		}
	}
}
