package models

type Ticker struct {
	ID              int     `json:"id"`
	ProductCode     string  `json:"product_code"`
	State           string  `json:"state"`
	Timestamp       int64   `json:"timestamp"`
	BestBid         float64 `json:"best_bid"`
	BestAsk         float64 `json:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     float64 `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   float64 `json:"total_ask_depth"`
	MarketBidSize   float64 `json:"market_bid_size"`
	MarketAskSize   float64 `json:"market_ask_size"`
	Ltp             float64 `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

func NewTicker(id int, productCode, state string, timestamp int64, bestBid, bestAsk, bestBidSize, bestAskSize, totalBidDepth, totalAskDepth, marketBidSize, marketAskSize, ltp, volume, volumeByProduct float64) Ticker {
	return Ticker{
		ID:              id,
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
