package handlers

import (
	"github.com/davecgh/go-spew/spew"
	"log"
	"time"
)

func GetSecondsOfTimeFrame(timeFrame string) int64 {
	switch timeFrame {
	case "1m":
		return 60 * 1
	case "3m":
		return 60 * 3
	case "5m":
		return 60 * 5
	case "10m":
		return 60 * 10
	case "15m":
		return 60 * 15
	case "30m":
		return 60 * 30
	case "1h":
		return 60 * 60
	case "2h":
		return 60 * 60 * 2
	case "3h":
		return 60 * 60 * 3
	case "4h":
		return 60 * 60 * 4
	case "1d":
		return 60 * 60 * 24
	case "1w":
		return 60 * 60 * 24 * 7
	case "1M":
		return 60 * 60 * 24 * 30
	default:
		return 60 * 1
	}
}

func GetSideResult(side string) bool {
	switch side {
	case "buy":
		return true
	case "sell":
		return false
	default:
		return true
	}
}

func TimeFrameHandler(timeFrame *TimeFrame) (response bool) {
	log.Printf("handling timeframe = %s\n", timeFrame.TimeFrame)
	defer func(resp bool) {
		if resp {
			log.Printf("handled timeframe %s with action = buy\n", timeFrame.TimeFrame)
		} else {
			log.Printf("handled timeframe %s with action = sell\n", timeFrame.TimeFrame)
		}
	}(response)
	if timeFrame.EnableEndOfTimeFrame {
		lastSignal := timeFrame.Storage.StableSignals[len(timeFrame.Storage.StableSignals)-1]
		return GetSideResult(lastSignal.Side)
	} else {
		timeDistribution := int64(timeFrame.TimeDistribution) * GetSecondsOfTimeFrame(timeFrame.TimeFrame) / 100
		var affectedSignals []WeightedSignals
		for i := len(timeFrame.Storage.Signals) - 1; i >= 0; i-- {
			timeDistanceTillNow := time.Now().Unix() - (timeFrame.Storage.Signals[i].PushedTime / 1000)

			var weight int64
			if len(affectedSignals) == 0 {
				weight = timeDistanceTillNow
			} else {
				weight = timeDistanceTillNow - affectedSignals[len(affectedSignals)-1].Weight
			}

			if timeDistanceTillNow >= timeDistribution {
				if len(affectedSignals) == 0 {
					weight = timeDistribution
				} else {
					weight = timeDistribution - affectedSignals[len(affectedSignals)-1].Weight
				}
				affectedSignals = append(affectedSignals, WeightedSignals{
					Signals: timeFrame.Storage.Signals[i],
					Weight:  weight,
				})
				break
			}
			affectedSignals = append(affectedSignals, WeightedSignals{
				Signals: timeFrame.Storage.Signals[i],
				Weight:  weight,
			})
		}

		var buy int64
		var sell int64
		for _, value := range affectedSignals {
			if value.Signals.Side == "buy" {
				buy += value.Weight
			} else {
				sell += value.Weight
			}
		}
		spew.Dump("signal weights", affectedSignals)
		if buy >= sell {
			return true
		} else {
			return false
		}
	}
}

func AndOfArray(array []bool) bool {
	result := array[0]
	for _, v := range array[1:] {
		result = result && v
	}
	return result
}

func OrOfArray(array []bool) bool {
	result := array[0]
	for _, v := range array[1:] {
		result = result || v
	}
	return result
}
