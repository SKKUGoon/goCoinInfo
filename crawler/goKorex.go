package crawler

import (
	"errors"
	"strings"
)

const (
	MARKETADD = "마켓 디지털 자산 추가"
	MARKETKRW = "KRW"
	MARKETBTC = "BTC"
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

func IfAssetKor(title upbitTitle) ([]string, error) {
	// find asset ticker inside title
	var titleAsset []string
	if strings.Contains(title.Title, MARKETADD) {
		// separate title by whitespace
		// return if alphabet = Ticker
		for _, c := range strings.Fields(title.Title) {
			a, err := clean([]byte(c))
			if err == nil {
				if a != MARKETKRW && a != MARKETBTC {
					// "KRW", "BTC" is a market standard
					titleAsset = append(titleAsset, a)
				}
			}
		}
	}
	if len(titleAsset) <= 0 {
		return titleAsset, errors.New("no asset found")
	} else {
		return titleAsset, nil
	}
}
