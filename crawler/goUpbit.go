package crawler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

type upbitTitle struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	ViewCount int       `json:"view_count"`
}

type upbitAPI struct {
	Success bool `json:"success"`
	Data    struct {
		TotalCount int          `json:"total_count"`
		TotalPages int          `json:"total_pages"`
		List       []upbitTitle `json:"list"`
	} `json:"data"`
}

func upbitNewAsset(content *upbitAPI, testMode bool) ([]string, error) {
	// upbitAPI struct =>
	// => goes through one by one
	// => if recent notice -> find asset
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
	cnt := new(upbitAPI)

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
