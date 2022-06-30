package handlers

import (
	"encoding/json"
	kucoin "github.com/Kucoin/kucoin-futures-go-sdk"
	"github.com/google/martian/log"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"strconv"
	"time"
)

const (
	ClientOIdPrefix = "trader_"
)

var SharedObject *Object
var SharedKuCoinService *kucoin.ApiService

func init() {
	SharedKuCoinService = kucoin.NewApiService(
		kucoin.ApiBaseURIOption("https://api-futures.kucoin.com"),
		kucoin.ApiKeyOption("62ab98e10d472900012cd167"),
		kucoin.ApiSecretOption("a068883d-7400-4d3d-b997-e3e473abec17"),
		kucoin.ApiPassPhraseOption("TestWow1234"),
		kucoin.ApiKeyVersionOption("2"),
	)

}

type Position struct {
	ID                string  `json:"id"`
	Symbol            string  `json:"symbol"`
	AutoDeposit       bool    `json:"autoDeposit"`
	MaintMarginReq    float64 `json:"maintMarginReq"`
	RiskLimit         int     `json:"riskLimit"`
	RealLeverage      float64 `json:"realLeverage"`
	CrossMode         bool    `json:"crossMode"`
	DelevPercentage   float64 `json:"delevPercentage"`
	OpeningTimestamp  int64   `json:"openingTimestamp"`
	CurrentTimestamp  int64   `json:"currentTimestamp"`
	CurrentQty        int     `json:"currentQty"`
	CurrentCost       float64 `json:"currentCost"`
	CurrentComm       float64 `json:"currentComm"`
	UnrealisedCost    float64 `json:"unrealisedCost"`
	RealisedGrossCost int     `json:"realisedGrossCost"`
	RealisedCost      float64 `json:"realisedCost"`
	IsOpen            bool    `json:"isOpen"`
	MarkPrice         float64 `json:"markPrice"`
	MarkValue         float64 `json:"markValue"`
	PosCost           float64 `json:"posCost"`
	PosCross          int     `json:"posCross"`
	PosInit           float64 `json:"posInit"`
	PosComm           float64 `json:"posComm"`
	PosLoss           int     `json:"posLoss"`
	PosMargin         float64 `json:"posMargin"`
	PosMaint          float64 `json:"posMaint"`
	MaintMargin       float64 `json:"maintMargin"`
	RealisedGrossPnl  int     `json:"realisedGrossPnl"`
	RealisedPnl       float64 `json:"realisedPnl"`
	UnrealisedPnl     float64 `json:"unrealisedPnl"`
	UnrealisedPnlPcnt float64 `json:"unrealisedPnlPcnt"`
	UnrealisedRoePcnt float64 `json:"unrealisedRoePcnt"`
	AvgEntryPrice     float64 `json:"avgEntryPrice"`
	LiquidationPrice  float64 `json:"liquidationPrice"`
	BankruptPrice     float64 `json:"bankruptPrice"`
	SettleCurrency    string  `json:"settleCurrency"`
	MaintainMargin    float64 `json:"maintainMargin"`
	RiskLimitLevel    int     `json:"riskLimitLevel"`

	Side string `json:"side"`
}

type Account struct {
	AccountEquity    float64 `json:"accountEquity"`
	UnrealisedPNL    float64 `json:"unrealisedPNL"`
	MarginBalance    float64 `json:"marginBalance"`
	PositionMargin   float64 `json:"positionMargin"`
	OrderMargin      int     `json:"orderMargin"`
	FrozenFunds      int     `json:"frozenFunds"`
	AvailableBalance float64 `json:"availableBalance"`
	Currency         string  `json:"currency"`
}

type Market struct {
	Symbol      string  `json:"symbol"`
	Granularity int     `json:"granularity"`
	TimePoint   int64   `json:"timePoint"`
	Value       float64 `json:"value"`
	IndexPrice  float64 `json:"indexPrice"`
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
	SizePercent               int          `json:"size_percent"`
	Currency                  string       `json:"currency"`
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
	o.ClosePosition()
	o.Exit = true
}

func (o *Object) ResumeStrategy() {
	o.Exit = false
}

func (o *Object) ActionHandler() {
	for action := range o.Action {
		o.ClosePosition()
		if o.Exit {
			continue
		}
		o.OpenPosition(action.Side)
	}
}

func (o *Object) SendAction() {
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

func (o *Object) ClosePosition() {
	position := o.GetOpenPosition()
	side := "buy"
	if position.Side == "buy" {
		side = "sell"
	}
	response, err := SharedKuCoinService.CreateOrder(map[string]string{
		"clientOid": uuid.New().String(),
		"side":      side,
		"symbol":    o.Strategy.Symbol,
		"leverage":  strconv.Itoa(o.Strategy.Leverage),
		"type":      "market",
		"size":      strconv.Itoa(position.CurrentQty),
	})
	if err != nil {
		log.Errorf("problem in placing order: %v", err)
		log.Errorf("problem in placing order: %v", response)
	}
}

func (o *Object) OpenPosition(side string) {
	account := o.GetAccountOverView()
	market := o.MarketPrice()
	size := int(float64(o.Strategy.SizePercent) / 100 * (account.AvailableBalance / market.Value * float64(o.Strategy.Leverage)) * 1000)
	response, err := SharedKuCoinService.CreateOrder(map[string]string{
		"clientOid": uuid.New().String(),
		"side":      side,
		"symbol":    o.Strategy.Symbol,
		"leverage":  strconv.Itoa(o.Strategy.Leverage),
		"type":      "market",
		"size":      strconv.Itoa(size),
	})
	if err != nil {
		log.Errorf("problem in placing order: %v", err)
		log.Errorf("problem in placing order: %v", response)
	}
}

func (o *Object) GetOpenPosition() (position Position) {
	resp, err := SharedKuCoinService.Position(o.Strategy.Symbol)
	if err != nil {
		log.Errorf("problem in calling get position, %v", err)
		return Position{}
	}
	json.Unmarshal(resp.RawData, &position)

	if position.CurrentQty < 0 {
		position.Side = "sell"
	} else {
		position.Side = "buy"
	}
	return position
}

func (o *Object) GetAccountOverView() (account Account) {
	response, err := SharedKuCoinService.AccountOverview(map[string]string{"currency": o.Strategy.Currency})
	if err != nil {
		log.Errorf("problem in calling get account, %v", err)
		return Account{}
	}
	json.Unmarshal(response.RawData, &account)
	return account
}

func (o *Object) MarketPrice() (market Market) {
	response, err := SharedKuCoinService.MarkPrice(o.Strategy.Symbol)
	if err != nil {
		log.Errorf("problem in calling get account, %v", err)
	}
	json.Unmarshal(response.RawData, &market)
	return market
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