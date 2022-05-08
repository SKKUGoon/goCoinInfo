package orderbook

import "time"

type OrderContent struct {
	A   string      `json:"a"`   // Asset
	N   string      `json:"n"`   // Strategy Name
	I   int         `json:"i"`   // Strategy ID
	T   time.Time   `json:"t"`   // Strategy Signal Generated Time
	ET  time.Time   `json:"et"`  // Strategy Execution Time
	TY  string      `json:"ty"`  // Asset Type
	SB  string      `json:"sb"`  // Strategy Originate Brokerage
	SBC int         `json:"sbc"` // Strategy Originate Brokerage Code
	OB  string      `json:"ob"`  // Order Brokerage
	OBC int         `json:"obc"` // Order Brokerage Code
	OD  OrderDetail `json:"od"`  // OrderDetail
}

type OrderDetail struct {
	P  string        `json:"p"`  // Pricing Type: MARKET vs LIMIT
	Q  float32       `json:"q"`  // Quantity of the Asset
	D  time.Duration `json:"d"`  // Duration of strategy
	LV int16         `json:"lv"` // Leverage
	LC float32       `json:"lc"` // Loss Cut
	DD float32       `json:"dd"` // Draw Down
}

type Brokerage struct {
	Name string
	Id   int
}
