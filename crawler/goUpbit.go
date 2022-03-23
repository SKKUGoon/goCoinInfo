package crawler

import (
	"encoding/json"
	"fmt"
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

func upbitNewAsset(content *upbitAPI) {
	//fmt.Println(content.Data.List)
	for _, notice := range content.Data.List {
		// 10 secs before
		t := time.Now()
		t = t.Add(-10 * time.Second)
		if notice.CreatedAt.After(t) {
			als, err := IfAssetKor(notice)
			if err == nil {
				fmt.Println(als)
			}
		}
	}
}

func CrawlUpbit() {
	URL := "https://api-manager.upbit.com/api/v1/notices?page=1"
	cnt := new(upbitAPI)

	resp, err := http.Get(URL)
	if err != nil {
		log.Println("[Crawler][Upbit] >>> Unsuccessful")
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(cnt)
	if err != nil {
		log.Println("[Crawler][Upbit] >>> JSON Decode Unsuccessful")
	} else {
		upbitNewAsset(cnt)
	}
}
