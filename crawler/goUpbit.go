package crawler

import (
	"encoding/json"
	"errors"
	"goBinance/orderbook"
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

func OrderUpbit(asset string) (orderbook.OrderContent, orderbook.OrderContent) {
	/*
		/ fill in the orderContent
	*/
	orderHfreq := orderbook.OrderContent{
		A:  asset,
		N:  UpbitOrderHF,
		I:  UpbitOrderHFId,
		T:  time.Now(),
		ET: time.Now(),
		TY: UpbitAssetType,
		B:  "binance",
		BC: 01,
		OD: orderbook.OrderDetail{
			P: "market",
			D: 10 * time.Second,
		},
	}

	orderLfreq := orderbook.OrderContent{
		A:  asset,
		N:  UpbitOrderLF,
		I:  UpbitOrderLFId,
		T:  time.Now(),
		TY: UpbitAssetType,
		B:  "binance",
		BC: 01,
		OD: orderbook.OrderDetail{
			P: "limit",
			D: 60 * time.Minute,
		},
	}

	return orderHfreq, orderLfreq
}
