package domestic

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
	return nil, errors.New(UpbitAssetNotFind)
}

func CrawlUpbit(testMode bool) ([]string, error) {
	/*
		/ HTTP Request for Upbit URL.
		/ returns the list of strings
	*/
	var target string
	cnt := new(UpbitAPI)

	// Test Mode
	if testMode == true {
		target = UpbitURLTEST
	} else {
		target = UpbitURL
	}

	// HTTP Request
	resp, err := http.Get(target)
	if err != nil {
		log.Println(UpbitReqErr)
	}

	// Process Request to Struct Container(cnt)
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(cnt)

	// Process Container Data(cnt)
	if err != nil {
		log.Println(UpbitJsonErr)
		return nil, err
	}

	result, err := upbitNewAsset(cnt, testMode)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
