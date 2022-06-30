package handlers

import (
	kucoin "github.com/Kucoin/kucoin-futures-go-sdk"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"strconv"
	"time"
)

const (
	ClientOIdPrefix = "crypto_trader_"
)

var SharedObject *Object
var SharedKuCoinService *kucoin.ApiService

func init() {
	SharedKuCoinService = kucoin.NewApiService(
		kucoin.ApiBaseURIOption("https://api-sandbox-futures.kucoin.com"),
		kucoin.ApiKeyOption("62aba38329c69200011e7f5d"),
		kucoin.ApiSecretOption("e1fbb51d-b338-4427-8ebf-7bdba7da8c6f"),
		kucoin.ApiPassPhraseOption("TestWow1234"),
		kucoin.ApiKeyVersionOption("2"),
	)

}

type Order struct {
	ClientOId string `json:"clientOid"`
	Side      string `json:"side"`
	Symbol    string `json:"symbol"`
	Leverage  string `json:"leverage"`
	Type      string `json:"type"`
}

type Position struct {
	Side string `json:"side"`
}

type Object struct {
	Strategy *Strategy    `json:"strategy"`
	Action   chan *Action `json:"-"`
	Exit     bool         `json:"-"`
}

type Action struct {
	Side string `json:"side"`
}

type Strategy struct {
	MainTimeFrame             *TimeFrame   `json:"main_time_frame"`
	SubTimeFrames             []*TimeFrame `json:"sub_time_frames"`
	SubTimeFramesOperation    string       `json:"sub_time_frames_operation"`
	MainSubTimeFrameOperation string       `json:"main_sub_time_frame_operation"`
	StopLoss                  int          `json:"stop_loss"`
	TakeProfit                int          `json:"take_profit"`
	Symbol                    string       `json:"symbol"`
	Leverage                  int          `json:"leverage"`
}

type TimeFrame struct {
	Storage                 *Storage `json:"-"`
	TimeFrame               string   `json:"time_frame"`
	EnableEndOfTimeFrame    bool     `json:"enable_end_of_time_frame"`
	SignalRepeatsToConsider int      `json:"signal_repeats_to_consider"`
	TimeDistribution        int      `json:"time_distribution"`
}

type Storage struct {
	Signals []*Signals `json:"signals"`
}

type Signals struct {
	Side        string `json:"side"`
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

type WeightedSignals struct {
	Signals *Signals `json:"signals"`
	Weight  int64    `json:"weight"`
}

func (o *Object) StopStrategy() {
	o.CloseAllPositions()
	o.Exit = true
}

func (o *Object) ResumeStrategy() {
	o.Exit = false
}

func (o *Object) ActionHandler() {
	for action := range o.Action {
		o.CloseAllPositions()
		if o.Exit {
			continue
		}
		o.OpenPosition(action.Side)
	}
}

func (o *Object) SendAction() {
	o.ClosePositionsExceptLast()
	position := o.GetOpenPosition()
	var subTimeFrameResults = null.NewBool(false, false)

	if len(o.Strategy.SubTimeFrames) != 0 {
		var listOfSubResults []bool
		for _, value := range o.Strategy.SubTimeFrames {
			listOfSubResults = append(listOfSubResults, TimeFrameHandler(value))
		}

		switch o.Strategy.SubTimeFramesOperation {
		case "AND":
			subTimeFrameResults = null.NewBool(AndOfArray(listOfSubResults), true)
		case "OR":
			subTimeFrameResults = null.NewBool(OrOfArray(listOfSubResults), true)
		default:
			subTimeFrameResults = null.NewBool(AndOfArray(listOfSubResults), true)
		}
	}

	mainResult := TimeFrameHandler(o.Strategy.MainTimeFrame)

	var action string
	if !subTimeFrameResults.IsZero() {
		switch o.Strategy.MainSubTimeFrameOperation {
		case "AND":
			if mainResult && subTimeFrameResults.Bool {
				action = "buy"
			} else {
				action = "sell"
			}
		case "OR":
			if mainResult || subTimeFrameResults.Bool {
				action = "buy"
			} else {
				action = "sell"
			}
		}
	} else {
		if mainResult {
			action = "buy"
		} else {
			action = "sell"
		}
	}

	if position.Side != action {
		o.Action <- &Action{Side: action}
	}
}

func (o *Object) ReceiveSignal(signal *Signals) {
	if signal.Side == o.Strategy.MainTimeFrame.TimeFrame {
		o.AddToMain(signal)
	} else {
		o.AddToSub(signal)
	}
	o.SendAction()
}

func (o *Object) AddToMain(signal *Signals) {
	SharedObject.Strategy.MainTimeFrame.Storage.Signals = append(SharedObject.Strategy.MainTimeFrame.Storage.Signals, signal)
}

func (o *Object) AddToSub(signal *Signals) {
	for _, value := range SharedObject.Strategy.SubTimeFrames {
		if value.TimeFrame == signal.TimeFrame {
			value.Storage.Signals = append(value.Storage.Signals, signal)
		}
	}
}

func (o *Object) CloseAllPositions() {
	panic("implement me")
}

func (o *Object) ClosePositionsExceptLast() {
	panic("implement me")
}

func (o *Object) OpenPosition(side string) {
	response, err := SharedKuCoinService.CreateOrder(map[string]string{
		"clientOid": ClientOIdPrefix + uuid.New().String(),
		"side":      side,
		"symbol":    o.Strategy.Symbol,
		"leverage":  strconv.Itoa(o.Strategy.Leverage),
	})
}

func (o *Object) GetOpenPosition() Position {
	panic("implement me")
}

func GetLastEndOfTimeFrameSignal(signals []*Signals) *Signals {
	for i := len(signals) - 1; i >= 0; i-- {
		if signals[i].PushedTime.Unix()%GetSecondsOfTimeFrame(signals[i].TimeFrame) < 5 {
			return signals[i]
		}
	}
	return signals[len(signals)-1]
}

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

func TimeFrameHandler(timeFrame *TimeFrame) bool {
	if timeFrame.EnableEndOfTimeFrame {
		lastSignal := GetLastEndOfTimeFrameSignal(timeFrame.Storage.Signals)
		return GetSideResult(lastSignal.Side)
	} else {
		timeDistribution := int64(timeFrame.TimeDistribution) * GetSecondsOfTimeFrame(timeFrame.TimeFrame) / 100
		var affectedSignals []WeightedSignals
		for i := len(timeFrame.Storage.Signals) - 1; i >= 0; i-- {
			timeDistanceTillNow := int64(time.Now().Sub(timeFrame.Storage.Signals[i].PushedTime).Seconds())

			var weight int64
			if len(affectedSignals) == 0 {
				weight = timeDistanceTillNow
			} else {
				weight = timeDistanceTillNow - affectedSignals[len(affectedSignals)-1].Weight
			}

			if timeDistanceTillNow >= timeDistribution {
				weight = timeDistribution - affectedSignals[len(affectedSignals)-1].Weight
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
