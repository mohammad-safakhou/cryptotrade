package handlers

import "context"

type Trader interface {
	Buy(ctx context.Context) error
	Sell(ctx context.Context) error
	Exit(ctx context.Context) error
}

type traderHandler struct {
}

func NewTraderHandler() Trader {
	return &traderHandler{}
}

func (t traderHandler) Buy(ctx context.Context) error {
	panic("implement me")
}

func (t traderHandler) Sell(ctx context.Context) error {
	panic("implement me")
}

func (t traderHandler) Exit(ctx context.Context) error {
	panic("implement me")
}
