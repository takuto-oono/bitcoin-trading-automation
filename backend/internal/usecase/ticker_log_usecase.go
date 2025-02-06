package usecase

import (
	"errors"
	"net/http"
	"time"

	"github.com/bitcoin-trading-automation/internal/bitflyer-api/api/models"
	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/mysql"
)

type TickerLog struct {
	Config config.Config
	MYSQL  mysql.MYSQL
}

type ITickerLog interface {
	GetTickerLogs() ([]mysql.Ticker, int, error)
	GetTickerLogByTickID(tickerID int) (*mysql.Ticker, int, error)
	PostTickerLog(ticker models.Ticker) (int, error)
}

func NewTickerLog(cfg config.Config) (ITickerLog, error) {
	sql, err := mysql.NewMYSQL(cfg)
	if err != nil {
		return nil, err
	}

	return &TickerLog{
		Config: cfg,
		MYSQL:  *sql,
	}, nil
}

func (t *TickerLog) GetTickerLogs() ([]mysql.Ticker, int, error) {
	tickerLogs, err := mysql.GetTickers(t.MYSQL.DB)
	if err != nil {
		return []mysql.Ticker{}, http.StatusInternalServerError, err
	}
	return tickerLogs, http.StatusOK, nil
}

func (t *TickerLog) GetTickerLogByTickID(tickerID int) (*mysql.Ticker, int, error) {
	ticker, err := mysql.GetTicker(t.MYSQL.DB, tickerID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if ticker == nil {
		return nil, http.StatusNotFound, nil
	}

	return ticker, http.StatusOK, nil
}

func (t *TickerLog) PostTickerLog(ticker models.Ticker) (int, error) {
	timestamp, err := parseTimestamp(ticker.Timestamp)
	if err != nil {
		return http.StatusBadRequest, err
	}

	myTicker := mysql.NewTicker(
		ticker.TickID,
		ticker.ProductCode,
		ticker.State,
		timestamp,
		ticker.BestBid,
		ticker.BestAsk,
		ticker.BestBidSize,
		ticker.BestAskSize,
		ticker.TotalBidDepth,
		ticker.TotalAskDepth,
		ticker.MarketBidSize,
		ticker.MarketAskSize,
		ticker.Ltp,
		ticker.Volume,
		ticker.VolumeByProduct,
	)

	if err := myTicker.Insert(t.MYSQL.DB); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func parseTimestamp(timestamp string) (int64, error) {
	layouts := []string{
		"2006-01-02T15:04:05.000Z",
		"2006-01-02T15:04:05.000",
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05",
	}

	for _, layout := range layouts {
		timestamp, err := time.Parse(layout, timestamp)
		if err == nil {
			return timestamp.Unix(), nil
		}
	}

	return 0, errors.New("invalid timestamp format")
}
