package handlers

import "time"

type Order struct {
	ClientOId string `json:"clientOid"`
	Side      string `json:"side"`
	Symbol    string `json:"symbol"`
	Leverage  string `json:"leverage"`
	Type      string `json:"type"`
}

type Object struct {
	Strategy *Strategy `json:"strategy"`
}

type Strategy struct {
	MainTimeFrame          *TimeFrame   `json:"main_time_frame"`
	SubTimeFrames          []*TimeFrame `json:"sub_time_frames"`
	SubTimeFramesOperation string       `json:"sub_time_frames_operation"`
	StopLoss               int          `json:"stop_loss"`
	TakeProfit             int          `json:"take_profit"`
	Leverage               int          `json:"leverage"`
}

type TimeFrame struct {
	Storage                 *Storage `json:"storage"`
	TimeFrame               string   `json:"time_frame"`
	EnableEndOfTimeFrame    bool     `json:"enable_end_of_time_frame"`
	SignalRepeatsToConsider int      `json:"signal_repeats_to_consider"`
}

type Storage struct {
	Signals []*Signals `json:"signals"`
}

type Signals struct {
	Action      string `json:"action"`
	TimeFrame   string `json:"time_frame"`
	Volume      string `json:"volume"`
	MarketPrice string `json:"market_price"`
	Open        string `json:"open"`
	Close       string `json:"close"`
	High        string `json:"high"`
	Low         string `json:"low"`

	PushedTime   time.Time `json:"pushed_time"`
	ReceivedTime time.Time `json:"received_time"`
}
