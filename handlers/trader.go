package handlers

import (
	"context"
	kucoin "github.com/Kucoin/kucoin-futures-go-sdk"
)

type Trader interface {
	Handler(ctx context.Context) error
	PlaceOrder(ctx context.Context, order Object) (*kucoin.ApiResponse, error)
	CloseOrder(ctx context.Context) error
	Exit(ctx context.Context) error
}

type traderHandler struct {
	KucoinService *kucoin.ApiService
}

func NewTraderHandler(KucoinService *kucoin.ApiService) Trader {
	return &traderHandler{KucoinService: KucoinService}
}

func (t *traderHandler) Handler(ctx context.Context) error {
	//dataChan := make(chan []byte)
	//t.MessageBrokerCandleHandler.Consumer(ctx, dataChan)
	//for data := range dataChan {
	//	switch string(data) {
	//	case "BUY":
	//		err := t.Buy(ctx)
	//		if err != nil {
	//			log.Printf("error on calling buy method in trader: %s", err.Error())
	//		}
	//	case "SELL":
	//		err := t.Sell(ctx)
	//		if err != nil {
	//			log.Printf("error on calling sell method in trader: %s", err.Error())
	//		}
	//	default:
	//		continue
	//	}
	//}
	return nil
}

type PlaceOrderResponse struct {
	Code string `json:"code"`
	Data struct {
		OrderId string `json:"order_id"`
	} `json:"data"`
}

func (t *traderHandler) PlaceOrder(ctx context.Context, order Object) (*kucoin.ApiResponse, error) {
	resp, err := t.KucoinService.CreateOrder(map[string]string{
		"clientOid": order.ClientOId,
		"side":      order.Side,
		"symbol":    order.Symbol,
		"leverage":  order.Leverage,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *traderHandler) CloseOrder(ctx context.Context) error {
	panic("implement me")
}

func (t *traderHandler) Exit(ctx context.Context) error {
	panic("implement me")
}
