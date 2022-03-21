package model

type Position struct {
	ID                 string  `json:"id"`
	AccountID          string  `json:"accountid"`
	OrderID            string  `json:"orderid"`
	OpenPrice          float64 `json:"openprice"`
	ClosePrice         float64 `json:"closeprice"`
	TakeProfit         float64 `json:"takeprofit"`
	StopLoss           float64 `json:"stoploss"`
	Symbol             string  `json:"symbol"`
	GuaranteedStopLoss bool    `json:"guaranteedstoploss"`
	State              string  `json:"state"`
	Quantity           float64 `json:"quantity"`
	Leverage           bool    `json:"leverage"`
	Side               string  `json:"side"`
}

type PostgresChannelMessage struct {
	Action   string `json:"action"`
	Position Position
}

type GeneratedPrice struct {
	Ask    float64 `json:"ask"`
	Bid    float64 `json:"bid"`
	Symbol string  `json:"symbol"`
}
