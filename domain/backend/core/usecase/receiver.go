package usecase

import (
	"context"
	"crypto/sha256"
	"cryptotrade/domain/backend/core/ports"
	"cryptotrade/utils"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type receiverHandler struct {
}

func NewReceiverHandler() ports.RawDataRepository {
	return &receiverHandler{}
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
	tt := time.Now().UnixNano()/int64(time.Millisecond)
	t := strconv.FormatInt(tt, 10)
	signData := fmt.Sprintf("access_id={%s}&timestamp={%s}&secret_key={%s}", accessId, t, secretKey)
	s := sha256.New()
	s.Write([]byte(signData))
	hash := hex.EncodeToString(s.Sum(nil))

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
	//err = ws.WriteMessage(websocket.BinaryMessage, getMessage(1, "state.subscribe", []string{}))
	//if err != nil {
	//	log.Println(err)
	//}
	err = ws.WriteMessage(websocket.BinaryMessage, getMessage(1, "server.sign", []interface{}{accessId, hash, tt}))
	if err != nil {
		log.Println(err)
	}
	//err = ws.WriteMessage(websocket.BinaryMessage, getMessage(1, "state.query", []interface{}{"BTCUSD", 86400}))
	//if err != nil {
	//	log.Println(err)
	//}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
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
