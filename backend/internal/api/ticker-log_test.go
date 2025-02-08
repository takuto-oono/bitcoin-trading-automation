package api

import (
	"math/rand"
	"testing"

	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/models"
)

func TestAPI_TickerLogHealthCheck(t *testing.T) {
	type fields struct {
		Config config.Config
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test TickerLogHealthCheck",
			fields: fields{
				Config: TestAPIConfig,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				Config: tt.fields.Config,
			}
			if err := api.TickerLogHealthCheck(); (err != nil) != tt.wantErr {
				t.Errorf("API.TickerLogHealthCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAPI_TickerLogPostTicker(t *testing.T) {
	type fields struct {
		Config config.Config
	}
	type args struct {
		ticker models.Ticker
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test TickerLogPostTicker",
			fields: fields{
				Config: TestAPIConfig,
			},
			args: args{
				ticker: models.Ticker{
					ID:              rand.Int(),
					ProductCode:     "BTC_USD",
					State:           "RUNNING",
					Timestamp:       1,
					BestBid:         1.0,
					BestAsk:         1.0,
					BestBidSize:     1.0,
					BestAskSize:     1.0,
					TotalBidDepth:   1.0,
					TotalAskDepth:   1.0,
					MarketBidSize:   1.0,
					MarketAskSize:   1.0,
					Ltp:             1.0,
					Volume:          1.0,
					VolumeByProduct: 1.0,
				},
			},
			wantErr: false,
		},
		{
			name: "Test TickerLogPostTicker Invalid Body",
			fields: fields{
				Config: TestAPIConfig,
			},
			args: args{
				ticker: models.Ticker{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				Config: tt.fields.Config,
			}
			if err := api.TickerLogPostTicker(tt.args.ticker); (err != nil) != tt.wantErr {
				t.Errorf("API.TickerLogPostTicker() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
