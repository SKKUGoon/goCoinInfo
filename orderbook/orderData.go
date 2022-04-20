package orderbook

import "time"

type OrderContent struct {
	A  string    `example:"BTC"`             // Asset
	N  string    `example:"upbit_ico_event"` // Strategy Name
	I  int       // Strategy ID
	T  time.Time // Strategy Signal Generated Time
	ET time.Time // Strategy Execution Time
	TY string    `example:"spot"`    // Asset Type
	B  string    `example:"binance"` // Brokerage
	BC int       // Brokerage Code
}
