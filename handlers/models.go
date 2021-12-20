package handlers

type ExchangeKLineResponseModel struct {
	Code    int             `json:"code"`
	Data    [][]interface{} `json:"data"`
	Message string          `json:"message"`
}

type ExchangeKLineModel struct {
	TimeFrame int64
	Opening   float64
	Closing   float64
	Highest   float64
	Lowest    float64
	Volume    float64
	Amount    float64
}
