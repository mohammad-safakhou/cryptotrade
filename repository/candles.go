package repository

import (
	"context"
	"cryptotrade/models"
	"database/sql"
	"time"
)

type CandlesRepository interface {
	SaveCandle(ctx context.Context, candle models.ExchangeKLineModel) (int64, error)
	GetLastNCandles(ctx context.Context, n int) ([]models.ExchangeKLineModel, error)
}

type candlesRepository struct {
	db *sql.DB
}

func NewCandlesRepository(db *sql.DB) CandlesRepository {
	return &candlesRepository{db: db}
}

func (cr *candlesRepository) SaveCandle(ctx context.Context, candle models.ExchangeKLineModel) (int64, error) {
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

func (cr *candlesRepository) GetLastNCandles(ctx context.Context, n int) ([]models.ExchangeKLineModel, error) {
	rows, err := cr.db.QueryContext(ctx, "SELECT * FROM candles_content order by id desc limit ?", n)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var kLines []models.ExchangeKLineModel
	for rows.Next() {
		var kLine models.ExchangeKLineModel
		err = rows.Scan(&kLine.Id, &kLine.TimeFrame, &kLine.Opening, &kLine.Closing, &kLine.Highest, &kLine.Lowest,
			&kLine.Volume, &kLine.Amount, &kLine.CreatedAt, &kLine.UpdatedAt, &kLine.DeletedAt)
		if err != nil {
			return nil, err
		}
		kLines = append(kLines, kLine)
	}
	return kLines, nil
}
