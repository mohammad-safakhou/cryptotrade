package handlers

import (
	"context"
	"cryptotrade/models"
	"database/sql"
	"encoding/json"
	kucoin "github.com/Kucoin/kucoin-futures-go-sdk"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
	"strconv"
	"sync"
	"time"
)

const (
	ClientOIdPrefix = "trader_"
)

var SharedObject *Object
var SharedKuCoinService *kucoin.ApiService
var SharedPostgresDB *sql.DB

func init() {
	SharedKuCoinService = kucoin.NewApiService(
		kucoin.ApiBaseURIOption("https://api-futures.kucoin.com"),
		kucoin.ApiKeyOption("62ab98e10d472900012cd167"),
		kucoin.ApiSecretOption("a068883d-7400-4d3d-b997-e3e473abec17"),
		kucoin.ApiPassPhraseOption("TestWow1234"),
		kucoin.ApiKeyVersionOption("2"),
	)
	spew.Dump("kucoin connected...", SharedKuCoinService)
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
	mu       sync.Mutex
	Strategy *Strategy    `json:"strategy"`
	Action   chan *Action `json:"-"`
	Exit     bool         `json:"-"`
}

type Action struct {
	Side string `json:"side"`
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
	Storage                 *Storage `json:"-"`
	TimeFrame               string   `json:"time_frame"`
	EnableEndOfTimeFrame    bool     `json:"enable_end_of_timeframe"`
	SignalRepeatsToConsider int      `json:"signal_repeats_to_consider"`
	TimeDistribution        int      `json:"time_distribution"`
}

type Storage struct {
	Signals       []*Signals `json:"signals"`
	StableSignals []*Signals `json:"stable_signals"`
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
	IsStable    bool   `json:"is_stable"`

	PushedTime   int64     `json:"pushed_time"`
	ReceivedTime time.Time `json:"received_time"`
}

type WeightedSignals struct {
	Signals *Signals `json:"signals"`
	Weight  int64    `json:"weight"`
}

func (o *Object) StopStrategy() {
	o.mu.Lock()
	defer o.mu.Unlock()
	{
		log.Printf("stopping strategy...\n")
		o.ClosePosition()
		o.Exit = true
	}
}

func (o *Object) ResumeStrategy() {
	o.mu.Lock()
	defer o.mu.Unlock()
	{
		log.Printf("resuming strategy...\n")
		o.Exit = false
	}
}

func (o *Object) ActionHandler() {
	for action := range o.Action {
		o.mu.Lock()
		{
			o.ClosePosition()
			if o.Exit {
				log.Printf("strategy is stopped...\n")
				o.mu.Unlock()
				continue
			}
			o.OpenPosition(action.Side)
			log.Printf("action on %s completed...\n", action.Side)
			log.Println("new signal ended ---------------------------------------------------------------------------------------------------")
		}
		o.mu.Unlock()
	}
}

func (o *Object) SendAction() {
	position := o.GetOpenPosition()
	var subTimeFrameResults = null.NewBool(false, true)

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
	if len(o.Strategy.SubTimeFrames) != 0 {
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

	log.Printf("new action = %s\n", action)
	if position.IsOpen {
		if position.Side != action {
			o.Action <- &Action{Side: action}
		} else {
			log.Printf("this action is open right now, skipping new action...")
		}
	} else {
		o.Action <- &Action{Side: action}
	}
}

func (o *Object) ReceiveSignal(signal *Signals) {
	o.mu.Lock()
	defer o.mu.Unlock()
	{
		log.Printf("receiving signal on timeframe %s - side %s\n", signal.TimeFrame, signal.Side)
		if signal.TimeFrame == o.Strategy.MainTimeFrame.TimeFrame {
			if signal.Side == "prev" {
				if len(o.Strategy.MainTimeFrame.Storage.StableSignals) == 0 {
					log.Println("signal is not valid yet, no prev signals available")
					return
				}
				signal.Side = o.Strategy.MainTimeFrame.Storage.StableSignals[len(o.Strategy.MainTimeFrame.Storage.StableSignals)-1].Side
			}
			if !o.AddToMain(signal) {
				return
			}
		} else {
			if signal.Side == "prev" {
				for _, value := range o.Strategy.SubTimeFrames {
					if value.TimeFrame == signal.TimeFrame {
						if len(value.Storage.StableSignals) == 0 {
							log.Println("signal is not valid yet, no prev signals available")
							return
						}
						signal.Side = value.Storage.StableSignals[len(value.Storage.StableSignals)-1].Side
					}
				}
			}
			if !o.AddToSub(signal) {
				log.Println("signal with time frame given is not defined in strategy... closing")
				log.Println("new signal ended ---------------------------------------------------------------------------------------------------")
				return
			}
		}
		o.SendAction()
	}
}

func (o *Object) AddToMain(signal *Signals) bool {
	if signal.IsStable {
		o.Strategy.MainTimeFrame.Storage.StableSignals = append(o.Strategy.MainTimeFrame.Storage.StableSignals, signal)
	} else {
		if len(o.Strategy.MainTimeFrame.Storage.StableSignals) == 0 {
			log.Println("strategy cant start without stable signal... closing")
			return false
		}
		o.Strategy.MainTimeFrame.Storage.Signals = append(o.Strategy.MainTimeFrame.Storage.Signals, signal)
	}
	return true
}

func (o *Object) AddToSub(signal *Signals) (isOk bool) {
	isTimeF := false
	for _, value := range o.Strategy.SubTimeFrames {
		if value.TimeFrame == signal.TimeFrame {
			isTimeF = true
			if signal.IsStable {
				value.Storage.StableSignals = append(value.Storage.StableSignals, signal)
			} else {
				if len(value.Storage.StableSignals) == 0 {
					log.Println("strategy cant start without stable signal... closing")
					return false
				}
				value.Storage.Signals = append(value.Storage.Signals, signal)
			}
		}
	}
	return isTimeF
}

func (o *Object) ClosePosition() {
	position := o.GetOpenPosition()
	if !position.IsOpen {
		return
	}
	side := "buy"
	if position.Side == "buy" {
		side = "sell"
	}
	if position.CurrentQty < 0 {
		position.CurrentQty = position.CurrentQty * -1
	}
	request := map[string]string{
		"clientOid": uuid.New().String(),
		"side":      side,
		"symbol":    o.Strategy.Symbol,
		"leverage":  strconv.Itoa(o.Strategy.Leverage),
		"type":      "market",
		"size":      strconv.Itoa(position.CurrentQty),
	}
	spew.Dump("closing position with: ", request)
	o.CreateOrder(request, 100, true)
}

func (o *Object) OpenPosition(side string) {
	log.Printf("opening position %s", side)
	account := o.GetAccountOverView()
	market := o.MarketPrice()
	size := int(float64(o.Strategy.SizePercent) / 100 * (account.AvailableBalance / market.Value * float64(o.Strategy.Leverage)) * 1000)
	if size == 0 {
		size = 1
	}
	request := map[string]string{
		"clientOid": uuid.New().String(),
		"side":      side,
		"symbol":    o.Strategy.Symbol,
		"leverage":  strconv.Itoa(o.Strategy.Leverage),
		"type":      "market",
		"size":      strconv.Itoa(size),
	}
	spew.Dump("opening position with:", request)
	o.CreateOrder(request, 100, false)
}

func (o *Object) GetOpenPosition() (position Position) {
	resp, err := SharedKuCoinService.Position(o.Strategy.Symbol)
	if err != nil {
		spew.Dump("problem in calling get position, ", err)
		return Position{}
	}
	json.Unmarshal(resp.RawData, &position)

	if position.CurrentQty < 0 {
		position.Side = "sell"
	} else {
		position.Side = "buy"
	}
	spew.Dump("getting open position: ", position)
	return position
}

func (o *Object) GetAccountOverView() (account Account) {
	response, err := SharedKuCoinService.AccountOverview(map[string]string{"currency": o.Strategy.Currency})
	if err != nil {
		spew.Dump("problem in calling get account, ", err)
		return Account{}
	}
	json.Unmarshal(response.RawData, &account)
	return account
}

func (o *Object) MarketPrice() (market Market) {
	response, err := SharedKuCoinService.MarkPrice(o.Strategy.Symbol)
	if err != nil {
		spew.Dump("problem in calling get account, ", err)
	}
	json.Unmarshal(response.RawData, &market)
	return market
}

func (o *Object) CreateOrder(request map[string]string, retry int, isClose bool) {
	for i := 0; i < retry; i++ {
		response, err := SharedKuCoinService.CreateOrder(request)
		if err != nil {
			spew.Dump("problem in placing order: ", err)
			continue
		}
		var orderResponse struct {
			OrderId string `json:"orderId"`
		}
		err = json.Unmarshal(response.RawData, &orderResponse)
		if err != nil {
			spew.Dump("problem in placing order: ", response.Code, response.Message, response.RawData)
			continue
		}
		openPosition := o.GetOpenPosition()
		if isClose {
			if openPosition.IsOpen {
				continue
			} else {
				log.Printf("order created with order id : %s", orderResponse.OrderId)
				position := models.Position{
					MarketPrice:  null.NewFloat64(openPosition.MarkPrice, true),
					Side:         null.NewString(request["side"], true),
					Leverage:     null.NewString(request["leverage"], true),
					PositionSize: null.NewString(request["size"], true),
					PositionType: null.NewString(request["type"], true),
					Symbol:       null.NewString(request["symbol"], true),
					IsClose:      null.NewBool(true, true),
				}
				err = position.Insert(context.TODO(), SharedPostgresDB, boil.Infer())
				if err != nil {
					log.Println("problem in inserting into position err = ", err.Error())
				}
				break
			}
		} else {
			if openPosition.IsOpen {
				log.Printf("order created with order id : %s", orderResponse.OrderId)
				position := models.Position{
					MarketPrice:  null.NewFloat64(openPosition.MarkPrice, true),
					Side:         null.NewString(request["side"], true),
					Leverage:     null.NewString(request["leverage"], true),
					PositionSize: null.NewString(request["size"], true),
					PositionType: null.NewString(request["type"], true),
					Symbol:       null.NewString(request["symbol"], true),
					IsClose:      null.NewBool(false, true),
				}
				err = position.Insert(context.TODO(), SharedPostgresDB, boil.Infer())
				if err != nil {
					log.Println("problem in inserting into position err = ", err.Error())
				}
				break
			}
		}
		continue
	}
}
