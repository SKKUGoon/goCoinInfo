package orderbook

import "time"

type OrderContent struct {
	A  string      `json:"a"`  // Asset
	N  string      `json:"n"`  // Strategy Name
	I  int         `json:"i"`  // Strategy ID
	T  time.Time   `json:"t"`  // Strategy Signal Generated Time
	ET time.Time   `json:"et"` // Strategy Execution Time
	TY string      `json:"ty"` // Asset Type
	B  string      `json:"b"`  // Brokerage
	BC int         `json:"bc"` // Brokerage Code
	OD OrderDetail `json:"od"` // OrderDetail
}

type OrderDetail struct {
	P string        `json:"p"` // Pricing Type: MARKET vs LIMIT
	Q float32       `json:"q"` // Quantity of the Asset
	D time.Duration `json:"D"` // Duration of strategy
}
