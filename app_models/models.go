package app_models

import "time"

type ExchangeKLineResponseModel struct {
	Code    int             `json:"code"`
	Data    [][]interface{} `json:"data"`
	Message string          `json:"message"`
}

type ExchangeKLineModel struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	TimeFrame float64
	Opening   string
	Closing   string
	Highest   string
	Lowest    string
	Volume    string
	Amount    string
}