package broadcast

type DataRecv struct {
	Msg string `json:"msg"`
	Dep string `json:"dep"`
}

type DataResp struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

type MessageRecv struct {
	SignalType string   `json:"signal_type"`
	Data       DataRecv `json:"data"`
}

type MessageResp struct {
	SignalType string   `json:"signal_type"`
	Data       DataResp `json:"data"`
}
