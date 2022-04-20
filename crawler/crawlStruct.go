package crawler

import "time"

type UpbitTitle struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	ViewCount int       `json:"view_count"`
}

type UpbitAPI struct {
	Success bool `json:"success"`
	Data    struct {
		TotalCount int          `json:"total_count"`
		TotalPages int          `json:"total_pages"`
		List       []UpbitTitle `json:"list"`
	} `json:"data"`
}
