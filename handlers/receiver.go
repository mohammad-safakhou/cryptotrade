package handlers

import (
	"context"
	"crypto/sha256"
	"cryptotrade/models"
	"cryptotrade/pkg"
	"cryptotrade/repository"
	"cryptotrade/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type Receiver interface {
	CallRepeater(ctx context.Context, coin string)
	Reader(ctx context.Context)
	Start(ctx context.Context)
}

type receiverHandler struct {
	WS              *websocket.Conn
	TimeFrame       models.TimeFrame
	NumberOfCandles int

	MessageBrokerCandleHandler pkg.MessageBrokerHandler
	CandlesRepository          repository.CandlesRepository
}

func NewReceiverHandler(
	WS *websocket.Conn,
	TimeFrame models.TimeFrame,
	NumberOfCandles int,

	CandlesRepository repository.CandlesRepository,
	MessageBrokerCandleHandler pkg.MessageBrokerHandler,
) Receiver {
	return &receiverHandler{
		WS:                         WS,
		TimeFrame:                  TimeFrame,
		NumberOfCandles:            NumberOfCandles,
		CandlesRepository:          CandlesRepository,
		MessageBrokerCandleHandler: MessageBrokerCandleHandler,
	}
}

func (rh *receiverHandler) CallRepeater(ctx context.Context, coin string) {
	tTimeFrame := models.TimeFrames[rh.TimeFrame]
	iTimeFrame := int(tTimeFrame.Seconds())
	for {
		now := time.Now()
		d := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		<-time.After(time.Duration(iTimeFrame-int(time.Now().Sub(d).Seconds())%iTimeFrame) * time.Second)

		go func() {
			time.Sleep(1 * time.Second)
			ctx, f := context.WithTimeout(ctx, 15*time.Second)
			defer f()
			response, _, err := utils.HttpCall(
				ctx,
				fmt.Sprintf("https://api.coinex.com/v1/market/kline?market=%s&type=%s&limit=%d", coin, rh.TimeFrame, 1),
				"",
				"GET",
				[]byte{},
			)
			if err != nil {
				fmt.Println(err)
				return
			}

			var exchangeKLineResponseData models.ExchangeKLineResponseModel
			json.Unmarshal(response, &exchangeKLineResponseData)

			exchangeKLineData := models.ExchangeKLineModel{
				TimeFrame: exchangeKLineResponseData.Data[0][0].(int64),
				Opening:   exchangeKLineResponseData.Data[0][1].(float64),
				Closing:   exchangeKLineResponseData.Data[0][2].(float64),
				Highest:   exchangeKLineResponseData.Data[0][3].(float64),
				Lowest:    exchangeKLineResponseData.Data[0][4].(float64),
				Volume:    exchangeKLineResponseData.Data[0][5].(float64),
				Amount:    exchangeKLineResponseData.Data[0][6].(float64),
			}

			// pushing data into data base
			_, err = rh.CandlesRepository.SaveCandle(ctx, exchangeKLineData)
			if err != nil {
				fmt.Println(err)
			}

			candles, err := rh.CandlesRepository.GetLastNCandles(ctx, rh.NumberOfCandles)
			if err != nil {
				fmt.Println(err)
			}

			// pushing data into kafka
			b, _ := json.Marshal(candles)
			err = rh.MessageBrokerCandleHandler.Push(ctx, b)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
}

func (rh *receiverHandler) Reader(ctx context.Context) {
	panic("implement me")
}

func (rh *receiverHandler) Start(ctx context.Context) {
	panic("implement me")
}

//Access ID
//DDF0F96626FD41BE9B2B588F39ED24F3
//Secret Key
//47564D69BC03066F4E440ED591AC7AB3F48E64188A17664B

func (rh *receiverHandler) DataReceiver(ctx context.Context) error {
	db, err := utils.PostgresConnection("localhost", "5432", "root", "root", "crytotrade", "disable")
	if err != nil {
		log.Println(err)
	}
	//var addr = flag.String("addr", "https://api.coinex.com/v1/market/list", "http service address")

	//u := url.URL{Scheme: "ws", Host: "https://api.coinex.com/v1/market/list", Path: "/echo"}
	//log.Printf("connecting to %s", u.String())

	//var d = websocket.Dialer{
	//	Subprotocols:     []string{"p1", "p2"},
	//	ReadBufferSize:   1024,
	//	WriteBufferSize:  1024,
	//	//Proxy:            http.ProxyFromEnvironment,
	//	HandshakeTimeout: time.Second * 10,
	//}
	ws, _, err := websocket.DefaultDialer.Dial("wss://perpetual.coinex.com/", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer ws.Close()

	accessId := "DDF0F96626FD41BE9B2B588F39ED24F3"
	secretKey := "47564D69BC03066F4E440ED591AC7AB3F48E64188A17664B"
	tt := time.Now().UnixNano() / int64(time.Millisecond)
	t := strconv.FormatInt(tt, 10)
	signData := fmt.Sprintf("access_id={%s}&timestamp={%s}&secret_key={%s}", accessId, t, secretKey)
	s := sha256.New()
	s.Write([]byte(signData))
	//hash := hex.EncodeToString(s.Sum(nil))

	go func(db *sql.DB) {
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			_, err = db.Exec(" INSERT INTO content (request_id, content, created_at) VALUES ($1, $2, $3)", 1, message, time.Now())
			if err != nil {
				log.Println(err)
			}
		}
	}(db)

	err = ws.WriteMessage(websocket.BinaryMessage, getMessage(1, "kline.query", []interface {
	}{
		"BTCUSDT",
		1638316800,
		1639353600,
		14400,
	}))
	//time.Sleep(2*time.Second)
	//err = ws.WriteMessage(websocket.BinaryMessage, getMessage(2, "state.query", []interface
	//{
	//}{
	//	"BTCUSDT",
	//	14400,
	//}))
	//time.Sleep(2*time.Second)
	//err = ws.WriteMessage(websocket.BinaryMessage, getMessage(3, "state.query", []interface
	//{
	//}{
	//	"BTCUSDT",
	//	14400,
	//}))
	//time.Sleep(2*time.Second)
	//err = ws.WriteMessage(websocket.BinaryMessage, getMessage(4, "state.query", []interface
	//{
	//}{
	//	"BTCUSDT",
	//	14400,
	//}))

	if err != nil {
		log.Println(err)
	}
	//err = ws.WriteMessage(websocket.BinaryMessage, getMessage(1, "server.sign", []interface{}{accessId, hash, tt}))
	//if err != nil {
	//	log.Println(err)
	//}
	//err = ws.WriteMessage(websocket.BinaryMessage, getMessage(1, "state.query", []interface{}{"BTCUSD", 86400}))
	//if err != nil {
	//	log.Println(err)
	//}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	ws.Close()
	return nil
}

func getMessage(id int, method string, params []interface{}) []byte {
	var data = struct {
		Id     int           `json:"id"`
		Method string        `json:"method"`
		Params []interface{} `json:"params"`
	}{
		Id:     id,
		Method: method,
		Params: params,
	}
	bData, err := json.Marshal(data)
	if err != nil {

	}
	return bData
}
