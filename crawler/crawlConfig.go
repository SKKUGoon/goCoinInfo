package crawler

// Upbit Configuration
const (
	UpbitURL        = "https://api-manager.upbit.com/api/v1/notices?page=1"
	UpbitURLTEST    = "https://api-manager.upbit.com/api/v1/notices?page=1"
	UpbitReqErr     = "[Crawler][Upbit] >>> Unsuccessful"
	UpbitJsonErr    = "[Crawler][Upbit] >>> JSON Decode Unsuccessful"
	UpbitAssetFound = "[Crawler][Upbit] >>> Asset found"

	UpbitOrderHF   = "upbit_market_add_event_high_freq"
	UpbitOrderHFId = 001
	UpbitOrderLF   = "upbit_market_add_event_low_freq"
	UpbitOrderLFId = 002
	UpbitAssetType = "spot"
)

// Bithumb Configuration
const (
	BithumbURL      = "https://cafe.bithumb.com/view/boards/43"
	BithumbURLTEST  = "https://cafe.bithumb.com/view/boards/43"
	BithumbURLErr   = "[Crawler][Bithumb] >>> URL creation unsuccessful"
	BithumbReqErr   = "[Crawler][Bithumb] >>> Unsuccessful request"
	BithumbParseErr = "[Crawler][Bithumb] >>> Fail to read HTML"

	BithumbOrderHF   = "bithumb_market_add_event_high_freq"
	BithumbOrderHFId = 011
	BithumbOrderLF   = "bithumb_market_add_event_low_freq"
	BithumbOrderLFId = 012
	BithumbAssetType = "spot"
)

// Keywords
const (
	MARKETADDUPBIT    = "마켓 디지털 자산 추가"
	MARKETADDBITHUMB2 = "원화 마켓 추가"
	MARKETADDBITHUMB3 = "BTC 마켓 추가"
	MARKETKRW         = "KRW"
	MARKETBTC         = "BTC"
)
