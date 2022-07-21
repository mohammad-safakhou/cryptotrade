package cmd

import (
	"cryptotrade/handlers"
)

type Object struct {
	Strategy *Strategy    `json:"strategy"`
	Exit     bool         `json:"exit"`
}

type Strategy struct {
	MainTimeFrame             *TimeFrame   `json:"main_timeframe"`
	SubTimeFrames             []*TimeFrame `json:"sub_time_frames"`
	SubTimeFramesOperation    string       `json:"sub_time_frames_operation"`
	MainSubTimeFrameOperation string       `json:"main_sub_time_frame_operation"`
	StopLoss                  int          `json:"stop_loss"`
	TakeProfit                int          `json:"take_profit"`
	Symbol                    string       `json:"symbol"`
	Leverage                  int          `json:"leverage"`
	SizePercent               int          `json:"size_percent"`
	Currency                  string       `json:"currency"`
}

type TimeFrame struct {
	Storage                 *Storage `json:"storage"`
	TimeFrame               string   `json:"time_frame"`
	EnableEndOfTimeFrame    bool     `json:"enable_end_of_timeframe"`
	SignalRepeatsToConsider int      `json:"signal_repeats_to_consider"`
	TimeDistribution        int      `json:"time_distribution"`
}

type Storage struct {
	Signals       []*handlers.Signals `json:"signals"`
	StableSignals []*handlers.Signals `json:"stable_signals"`
}
