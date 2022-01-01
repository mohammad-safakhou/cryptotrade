package repository

import (
	"context"
	"cryptotrade/app_models"
	"cryptotrade/models"
	"database/sql"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type CandlesRepository interface {
	SaveCandle(ctx context.Context, candle app_models.ExchangeKLineModel) (int64, error)
	GetLastNCandles(ctx context.Context, n int) ([]app_models.ExchangeKLineModel, error)
}

type candlesRepository struct {
	db *sql.DB
}

func NewCandlesRepository(db *sql.DB) CandlesRepository {
	return &candlesRepository{db: db}
}

func (cr *candlesRepository) SaveCandle(ctx context.Context, candle app_models.ExchangeKLineModel) (int64, error) {
	candleContent := models.CandlesContent{
		TimeFrame: null.NewFloat64(candle.TimeFrame, true),
		Opening:   null.NewString(candle.Opening, true),
		Closing:   null.NewString(candle.Closing, true),
		Highest:   null.NewString(candle.Highest, true),
		Lowest:    null.NewString(candle.Lowest, true),
		Volume:    null.NewString(candle.Volume, true),
		Amount:    null.NewString(candle.Amount, true),
	}
	err := candleContent.Insert(ctx, cr.db, boil.Infer())
	if err != nil {
		return 0, nil
	}
	return candleContent.ID, nil
}

func (cr *candlesRepository) GetLastNCandles(ctx context.Context, n int) ([]app_models.ExchangeKLineModel, error) {
	candles, err := models.CandlesContents(qm.OrderBy("time_frame desc"), qm.Limit(n)).All(ctx, cr.db)
	if err != nil {
		return nil, err
	}
	var kLines []app_models.ExchangeKLineModel
	for _, item := range candles {
		var kLine = app_models.ExchangeKLineModel{
			Id:        item.ID,
			CreatedAt: item.CreatedAt.Time,
			UpdatedAt: item.UpdatedAt.Time,
			DeletedAt: item.DeletedAt.Time,
			TimeFrame: item.TimeFrame.Float64,
			Opening:   item.Opening.String,
			Closing:   item.Closing.String,
			Highest:   item.Highest.String,
			Lowest:    item.Lowest.String,
			Volume:    item.Volume.String,
			Amount:    item.Amount.String,
		}
		kLines = append(kLines, kLine)
	}
	return kLines, nil
}
