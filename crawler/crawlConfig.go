package crawler

// Upbit Configuration
const (
	UpbitURL        = "https://api-manager.upbit.com/api/v1/notices?page=1"
	UpbitURLTEST    = "https://api-manager.upbit.com/api/v1/notices?page=3"
	UpbitReqErr     = "[Crawler][Upbit] >>> Unsuccessful"
	UpbitJsonErr    = "[Crawler][Upbit] >>> JSON Decode Unsuccessful"
	UpbitAssetFound = "[Crawler][Upbit] >>> Asset found"

	UpbitOrderHF   = "upbit market add event high freq"
	UpbitOrderHFId = 001
	UpbitOrderLF   = "upbit market add event low freq"
	UpbitOrderLFId = 002
	UpbitAssetType = "spot"
)

// Bithumb Configuration
const (
	BithumbURL      = "https://cafe.bithumb.com/view/boards/43"
	BithumbURLTEST  = ""
	BithumbURLErr   = "[Crawler][Bithumb] >>> URL creation unsuccessful"
	BithumbReqErr   = "[Crawler][Bithumb] >>> Unsuccessful request"
	BithumbParseErr = "[Crawler][Bithumb] >>> Fail to read HTML"
)

// Keywords
const (
	MARKETADDUPBIT    = "마켓 디지털 자산 추가"
	MARKETADDBITHUMB0 = "마켓 추가"
	MARKETADDBITHUMB1 = "마켓 상장"
	MARKETKRW         = "KRW"
	MARKETBTC         = "BTC"
)
