package domestic

import (
	"errors"
	"golang.org/x/net/html"
	"log"
	"strings"
	"time"
)

// AssetUpbit PUBLIC
func AssetUpbit(title UpbitTitle) ([]string, error) {
	// find asset ticker inside upbit api signals
	t := title.Title
	if strings.Contains(t, MARKETADDUPBIT) {
		data, err := ifAssetKor(t)
		if err != nil {
			return nil, errors.New("no asset found")
		} else {
			return data, nil
		}
	} else {
		return nil, errors.New("not asset statement")
	}
}

// AssetBithumb PUBLIC
func AssetBithumb(text string, wanted map[string]bool) ([]BithumbTitle, error) {
	/*
		/ bithumb does not provide news api
		/ parse the raw html files.
		/ crawl the TextToken that comes after <a> StartToken: asset name
		/ crawl the TextToken that comes after <td> StartToken: date
	*/
	var bithumbContainer []BithumbTitle
	var assetParse = false
	var tmpContainer = BithumbTitle{}

	tkn := html.NewTokenizer(strings.NewReader(text))
	for {
		tt := tkn.Next()

		switch {
		// End of Parsing
		case tt == html.ErrorToken:
			return bithumbContainer, nil

		// ex) <a> </a> <= <a> is a StartTagToken
		case tt == html.StartTagToken:
			tn := tkn.Token()

			switch {
			// process datetime
			case wanted[tn.Data] == true && tn.Data == "td":
				tt, tn = tkn.Next(), tkn.Token()

				td, err := extractDate(tt, tn)
				if err == nil {
					if assetParse == true {
						// end parsing
						assetParse = false
						tmpContainer.CreatedAt = td

						// add parsed temp container {} to main container {}
						bithumbContainer = append(bithumbContainer, tmpContainer)
					}
				}

			// process asset content
			case wanted[tn.Data] == true && tn.Data == "a":
				tt, tn = tkn.Next(), tkn.Token()

				asset, err := extractAsset(tt, tn)
				if err == nil {
					// start parsing
					assetParse = true
					tmpContainer.Asset = asset
				}

			default:
				continue
			}
		}
	}
}

func clean(s []byte) (string, error) {
	/*
		/ ticker only contains alphabet character
		/ returns (alphabet value, error)
	*/
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
	/*
		/ find out whether title contains asset
	*/
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

func extractDate(tp html.TokenType, t html.Token) (time.Time, error) {
	/*
		/ Bithumb gives out only time information
		/ prstLayout -> iff the information is uploaded that day
		/ histLayout -> historical posts.
		/ location -> always kst base
	*/
	tNow := time.Now()
	prstLayout := " 15:04"     // HH:MM
	histLayout := "2006.01.02" // YYYY.MM.DD
	tloc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		log.Panic("incorrect timezone location")
	}

	// parse dates
	if tp == html.TextToken {
		switch {
		case strings.Contains(t.Data, ":"):
			td, err := time.ParseInLocation(prstLayout, t.Data, tloc)
			// Add dates. -1 since it the default dates starts at 0000-01-01
			td = td.AddDate(tNow.Year(), int(tNow.Month())-1, tNow.Day()-1)

			if err != nil {
				return time.Time{}, errors.New("cannot parse today's date")
			} else {
				return td, nil
			}

		default:
			td, err := time.ParseInLocation(histLayout, t.Data, tloc)

			if err != nil {
				return time.Time{}, errors.New("cannot parse historical date")
			} else {
				return td, nil
			}
		}
	} else {
		return time.Time{}, errors.New("not html text token")
	}
}

func extractAsset(tp html.TokenType, t html.Token) ([]string, error) {
	if tp == html.TextToken {
		// If MARKET ADD event
		contain1 := strings.Contains(t.Data, MARKETADDBITHUMB2)
		contain2 := strings.Contains(t.Data, MARKETADDBITHUMB3)

		if contain1 || contain2 {
			asset, err := ifAssetKor(t.Data)
			if err != nil {
				return nil, errors.New("market add but parsing error")
			} else {
				return asset, nil
			}
		} else {
			return nil, errors.New("not asset statement")
		}
	} else {
		return nil, errors.New("not text token")
	}
}
