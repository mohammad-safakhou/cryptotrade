package repository

import (
	"context"
	"cryptotrade/handlers"
	"database/sql"
	"time"
)

type CandlesRepository interface {
	SaveCandle(ctx context.Context, candle handlers.ExchangeKLineModel) (int64, error)
}

type candlesRepository struct {
	db *sql.DB
}

func NewCandlesRepository(db *sql.DB) CandlesRepository {
	return &candlesRepository{db: db}
}

func (cr *candlesRepository) SaveCandle(ctx context.Context, candle handlers.ExchangeKLineModel) (int64, error) {
	res, err := cr.db.ExecContext(ctx, "INSERT INTO candles_content ("+
		"time_frame,opening,closing,highest,lowest,volume,amount,created_at,updated_at,deleted_at) VALUES ("+
		"?,?,?,?,?,?,?,?,?,?)",
		candle.TimeFrame, candle.Opening, candle.Closing, candle.Highest, candle.Lowest,
		candle.Volume, candle.Amount, time.Now(), time.Now(), nil)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
