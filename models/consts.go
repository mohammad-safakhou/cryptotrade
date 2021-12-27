package models

import "time"

var TimeFrames = map[TimeFrame]time.Duration{
	TimeFrameUnspecified: 0 * time.Second,
	TimeFrame1min:        60 * time.Second,
	TimeFrame3min:        180 * time.Second,
	TimeFrame5min:        300 * time.Second,
	TimeFrame15min:       900 * time.Second,
	TimeFrame30min:       1800 * time.Second,
	TimeFrame1hour:       3600 * time.Second,
	TimeFrame2hour:       7200 * time.Second,
	TimeFrame4hour:       14400 * time.Second,
	TimeFrame6hour:       21600 * time.Second,
	TimeFrame12hour:      43200 * time.Second,
	//TimeFrame1day:        86400,
	//TimeFrame3day:        259200,
	//TimeFrame1week:       604800,
}

type TimeFrame string

const (
	TimeFrameUnspecified = ""
	TimeFrame1min        = "1min"
	TimeFrame3min        = "3min"
	TimeFrame5min        = "5min"
	TimeFrame15min       = "15min"
	TimeFrame30min       = "30min"
	TimeFrame1hour       = "1hour"
	TimeFrame2hour       = "2hour"
	TimeFrame4hour       = "4hour"
	TimeFrame6hour       = "6hour"
	TimeFrame12hour      = "12hour"
	//TimeFrame1day        = "1day"
	//TimeFrame3day        = "3day"
	//TimeFrame1week       = "1week"
)

type TimeFrameDailySecs []time.Duration
