// Package model contains structures related to position service
package model

// Position contains all needed information about position
type Position struct {
	ID                 string  `json:"id"`
	AccountID          string  `json:"account_id"`
	OrderID            string  `json:"order_id"`
	OpenPrice          float64 `json:"open_price"`
	ClosePrice         float64 `json:"close_price"`
	TakeProfit         float64 `json:"take_profit"`
	StopLoss           float64 `json:"stop_loss"`
	Symbol             string  `json:"symbol"`
	GuaranteedStopLoss bool    `json:"guaranteed_stop_loss"`
	State              string  `json:"state"`
	Quantity           float64 `json:"quantity"`
	Leverage           bool    `json:"leverage"`
	Side               string  `json:"side"`
}

// PostgresChannelMessage is used for pg notify trigger
type PostgresChannelMessage struct {
	Action   string `json:"action"`
	Position Position
}

// GeneratedPrice describe stock
type GeneratedPrice struct {
	Ask    float64 `json:"ask"`
	Bid    float64 `json:"bid"`
	Symbol string  `json:"symbol"`
}
