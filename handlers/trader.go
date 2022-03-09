package handlers

import (
	"context"
	"cryptotrade/pkg"
	"log"
)

type Trader interface {
	Handler(ctx context.Context) error
	Buy(ctx context.Context) error
	Sell(ctx context.Context) error
	Exit(ctx context.Context) error
}

type traderHandler struct {
	MessageBrokerCandleHandler pkg.MessageBrokerHandler
}

func NewTraderHandler(messageBrokerCandleHandler pkg.MessageBrokerHandler) Trader {
	return &traderHandler{MessageBrokerCandleHandler: messageBrokerCandleHandler}
}

func (t *traderHandler) Handler(ctx context.Context) error {
	dataChan := make(chan []byte)
	t.MessageBrokerCandleHandler.Consumer(ctx, dataChan)
	for data := range dataChan {
		switch string(data) {
		case "BUY":
			err := t.Buy(ctx)
			if err != nil {
				log.Printf("error on calling buy method in trader: %s", err.Error())
			}
		case "SELL":
			err := t.Sell(ctx)
			if err != nil {
				log.Printf("error on calling sell method in trader: %s", err.Error())
			}
		default:
			continue
		}
	}
	return nil
}

func (t *traderHandler) Buy(ctx context.Context) error {
	panic("implement me")
}

func (t *traderHandler) Sell(ctx context.Context) error {
	panic("implement me")
}

func (t *traderHandler) Exit(ctx context.Context) error {
	panic("implement me")
}
