package crawler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

func upbitNewAsset(content *UpbitAPI, testMode bool) ([]string, error) {
	/*
		/ UpbitAPI struct
		/ goes through one by one
		/ if recent notice
		/ find asset
	*/
	for _, notice := range content.Data.List {
		t := time.Now()
		// created recently - 10 seconds
		t = t.Add(-10 * time.Second)
		if testMode || notice.CreatedAt.After(t) {
			als, err := AssetUpbit(notice)
			if err == nil {
				return als, nil
			}
		}
	}
	return nil, errors.New("no signal")
}

func CrawlUpbit(testMode bool) ([]string, error) {
	var target string
	cnt := new(UpbitAPI)

	if testMode == true {
		target = UpbitURLTEST
	} else {
		target = UpbitURL
	}

	resp, err := http.Get(target)
	if err != nil {
		log.Println(UpbitReqErr)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(cnt)
	if err != nil {
		log.Println(UpbitJsonErr)
		return nil, err
	} else {
		result, err := upbitNewAsset(cnt, testMode)
		if err != nil {
			log.Println(err)
			return nil, err
		} else {
			log.Println(UpbitAssetFound)
			return result, nil
		}
	}
}
