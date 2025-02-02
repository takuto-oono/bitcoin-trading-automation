package mysql

import (
	"errors"

	"gorm.io/gorm"
)

type Ticker struct {
	TickID          int     `json:"tick_id" gorm:"column:tick_id primary_key"`
	ProductCode     string  `json:"product_code" gorm:"column:product_code"`
	State           string  `json:"state" gorm:"column:state"`
	Timestamp       int64   `json:"timestamp" gorm:"column:timestamp"`
	BestBid         float64 `json:"best_bid" gorm:"column:best_bid"`
	BestAsk         float64 `json:"best_ask" gorm:"column:best_ask"`
	BestBidSize     float64 `json:"best_bid_size" gorm:"column:best_bid_size"`
	BestAskSize     float64 `json:"best_ask_size" gorm:"column:best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth" gorm:"column:total_bid_depth"`
	TotalAskDepth   float64 `json:"total_ask_depth" gorm:"column:total_ask_depth"`
	MarketBidSize   float64 `json:"market_bid_size" gorm:"column:market_bid_size"`
	MarketAskSize   float64 `json:"market_ask_size" gorm:"column:market_ask_size"`
	Ltp             float64 `json:"ltp" gorm:"column:ltp"`
	Volume          float64 `json:"volume" gorm:"column:volume"`
	VolumeByProduct float64 `json:"volume_by_product" gorm:"column:volume_by_product"`
}

func NewTicker(tickID int, productCode, state string, timestamp int64, bestBid, bestAsk, bestBidSize, bestAskSize, totalBidDepth, totalAskDepth, marketBidSize, marketAskSize, ltp, volume, volumeByProduct float64) *Ticker {
	return &Ticker{
		TickID:          tickID,
		ProductCode:     productCode,
		State:           state,
		Timestamp:       timestamp,
		BestBid:         bestBid,
		BestAsk:         bestAsk,
		BestBidSize:     bestBidSize,
		BestAskSize:     bestAskSize,
		TotalBidDepth:   totalBidDepth,
		TotalAskDepth:   totalAskDepth,
		MarketBidSize:   marketBidSize,
		MarketAskSize:   marketAskSize,
		Ltp:             ltp,
		Volume:          volume,
		VolumeByProduct: volumeByProduct,
	}
}

func (Ticker) TableName() string {
	return "ticker"
}

func (t *Ticker) BeforeCreate(tx *gorm.DB) error {
	if t.TickID == 0 {
		return errors.New("tick_id is required")
	}
	if t.ProductCode == "" {
		return errors.New("product_code is required")
	}
	if t.State == "" {
		return errors.New("state is required")
	}
	if t.Timestamp == 0 {
		return errors.New("timestamp is required")
	}
	return nil
}

func (t *Ticker) Insert(tx *gorm.DB) error {
	return tx.Create(t).Error
}

func GetTicker(tx *gorm.DB, tickID int) (*Ticker, error) {
	var ticker Ticker
	if err := tx.First(&ticker, tickID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &ticker, nil
}

func GetTickers(tx *gorm.DB) ([]Ticker, error) {
	var tickers []Ticker
	if err := tx.Find(&tickers).Error; err != nil {
		return nil, err
	}
	return tickers, nil
}
