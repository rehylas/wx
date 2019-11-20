package worker

const TRADE_TYPE_BUY = "buy"
const TRADE_TYPE_SELL = "sell"

const TRADE_STATE_0 = 0
const TRADE_STATE_1 = 1
const TRADE_STATE_2 = 2
const TRADE_STATE_3 = 3
const TRADE_STATE_4 = 4

type Trade struct {
	Symbol     string  `json:"symbol"`
	Type       string  `json:"type"`
	State      int     `json:"state"`
	Inprice    float64 `json:"inprice"`
	Outprice   float64 `json:"outprice"`
	Vol        int     `json:"vol"`
	Indt       string  `json:"indt"`
	Outdt      string  `json:"outdt"`
	OrderIdCtp string  `json:"orderidctp"`
	OrderIdSys string  `json:"orderidsys"`
}
