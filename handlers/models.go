package handlers

var KucoinBaseUrl string = ""
type Order struct {
	ClientOId string `json:"clientOid"`
	Side      string `json:"side"`
	Symbol    string `json:"symbol"`
	Leverage  string `json:"leverage"`
}
