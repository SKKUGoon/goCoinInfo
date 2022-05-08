package orderbook

const (
	Strat1 = "market_add_event_high_freq_short"
	Strat2 = "market_add_event_low_freq_long"
	Strat3 = "market_add_event_low_freq_short"

	Strat1Id = 001
	Strat2Id = 002
	Strat3Id = 003

	AssetStrat1 = "futures"
	AssetStrat2 = "futures"
	AssetStrat3 = "futures"
)

var BrokerageInfo = map[int]Brokerage{
	1: Brokerage{
		Name: "binance", Id: 001,
	},
	2: Brokerage{
		Name: "upbit", Id: 002,
	},
	3: Brokerage{
		Name: "bithumb", Id: 003,
	},
}
