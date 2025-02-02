package mysql

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestTicker_Insert(t *testing.T) {
	type fields struct {
		TickID          int
		ProductCode     string
		State           string
		Timestamp       int64
		BestBid         float64
		BestAsk         float64
		BestBidSize     float64
		BestAskSize     float64
		TotalBidDepth   float64
		TotalAskDepth   float64
		MarketBidSize   float64
		MarketAskSize   float64
		Ltp             float64
		Volume          float64
		VolumeByProduct float64
	}
	type args struct {
		tx *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test Insert",
			fields: fields{
				TickID:          rand.Int(),
				ProductCode:     "BTC_USD",
				State:           "RUNNING",
				Timestamp:       time.Now().Unix(),
				BestBid:         100.0,
				BestAsk:         200.0,
				BestBidSize:     0.1,
				BestAskSize:     0.2,
				TotalBidDepth:   1000.0,
				TotalAskDepth:   2000.0,
				MarketBidSize:   10.0,
				MarketAskSize:   20.0,
				Ltp:             150.0,
				Volume:          10000.0,
				VolumeByProduct: 1000.0,
			},
			args: args{
				tx: func() *gorm.DB {
					db, err := connectMYSQL(TestMYSQLConfig)
					if err != nil {
						panic(err)
					}
					return db
				}(),
			},
			wantErr: false,
		},
		{
			name: "Test Insert ticker id is nil",
			fields: fields{
				TickID:          0,
				ProductCode:     "",
				State:           "",
				Timestamp:       0,
				BestBid:         0.0,
				BestAsk:         0.0,
				BestBidSize:     0.0,
				BestAskSize:     0.0,
				TotalBidDepth:   0.0,
				TotalAskDepth:   0.0,
				MarketBidSize:   0.0,
				MarketAskSize:   0.0,
				Ltp:             0.0,
				Volume:          0.0,
				VolumeByProduct: 0.0,
			},
			args: args{
				tx: func() *gorm.DB {
					db, err := connectMYSQL(TestMYSQLConfig)
					if err != nil {
						panic(err)
					}
					return db
				}(),
			},
			wantErr: true,
		},
		{
			name: "Test Insert product code is nil",
			fields: fields{
				TickID:          rand.Int(),
				ProductCode:     "",
				State:           "",
				Timestamp:       0,
				BestBid:         0.0,
				BestAsk:         0.0,
				BestBidSize:     0.0,
				BestAskSize:     0.0,
				TotalBidDepth:   0.0,
				TotalAskDepth:   0.0,
				MarketBidSize:   0.0,
				MarketAskSize:   0.0,
				Ltp:             0.0,
				Volume:          0.0,
				VolumeByProduct: 0.0,
			},
			args: args{
				tx: func() *gorm.DB {
					db, err := connectMYSQL(TestMYSQLConfig)
					if err != nil {
						panic(err)
					}
					return db
				}(),
			},
			wantErr: true,
		},
		{
			name: "Test Insert state is nil",
			fields: fields{
				TickID:          rand.Int(),
				ProductCode:     "BTC_USD",
				State:           "",
				Timestamp:       0,
				BestBid:         0.0,
				BestAsk:         0.0,
				BestBidSize:     0.0,
				BestAskSize:     0.0,
				TotalBidDepth:   0.0,
				TotalAskDepth:   0.0,
				MarketBidSize:   0.0,
				MarketAskSize:   0.0,
				Ltp:             0.0,
				Volume:          0.0,
				VolumeByProduct: 0.0,
			},
			args: args{
				tx: func() *gorm.DB {
					db, err := connectMYSQL(TestMYSQLConfig)
					if err != nil {
						panic(err)
					}
					return db
				}(),
			},
			wantErr: true,
		},
		{
			name: "Test Insert timestamp is nil",
			fields: fields{
				TickID:          rand.Int(),
				ProductCode:     "BTC_USD",
				State:           "RUNNING",
				Timestamp:       0,
				BestBid:         0.0,
				BestAsk:         0.0,
				BestBidSize:     0.0,
				BestAskSize:     0.0,
				TotalBidDepth:   0.0,
				TotalAskDepth:   0.0,
				MarketBidSize:   0.0,
				MarketAskSize:   0.0,
				Ltp:             0.0,
				Volume:          0.0,
				VolumeByProduct: 0.0,
			},
			args: args{
				tx: func() *gorm.DB {
					db, err := connectMYSQL(TestMYSQLConfig)
					if err != nil {
						panic(err)
					}
					return db
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Ticker{
				TickID:          tt.fields.TickID,
				ProductCode:     tt.fields.ProductCode,
				State:           tt.fields.State,
				Timestamp:       tt.fields.Timestamp,
				BestBid:         tt.fields.BestBid,
				BestAsk:         tt.fields.BestAsk,
				BestBidSize:     tt.fields.BestBidSize,
				BestAskSize:     tt.fields.BestAskSize,
				TotalBidDepth:   tt.fields.TotalBidDepth,
				TotalAskDepth:   tt.fields.TotalAskDepth,
				MarketBidSize:   tt.fields.MarketBidSize,
				MarketAskSize:   tt.fields.MarketAskSize,
				Ltp:             tt.fields.Ltp,
				Volume:          tt.fields.Volume,
				VolumeByProduct: tt.fields.VolumeByProduct,
			}
			if err := tr.Insert(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("Ticker.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetTicker(t *testing.T) {
	db, err := connectMYSQL(TestMYSQLConfig)
	if err != nil {
		panic(err)
	}

	ticker := NewTicker(rand.Int(), "BTC_USD", "RUNNING", time.Now().Unix(), 100.0, 200.0, 0.1, 0.2, 1000.0, 2000.0, 10.0, 20.0, 150.0, 10000.0, 1000.0)
	if err := ticker.Insert(db); err != nil {
		panic(err)
	}

	existTicker := ticker

	type args struct {
		tx     *gorm.DB
		tickID int
	}
	tests := []struct {
		name    string
		args    args
		want    *Ticker
		wantErr bool
	}{
		{
			name: "Test GetTicker",
			args: args{
				tx: func() *gorm.DB {
					db, err := connectMYSQL(TestMYSQLConfig)
					if err != nil {
						panic(err)
					}
					return db
				}(),
				tickID: existTicker.TickID,
			},
			want:    existTicker,
			wantErr: false,
		},
		{
			name: "Test GetTicker tick id is nil",
			args: args{
				tx: func() *gorm.DB {
					db, err := connectMYSQL(TestMYSQLConfig)
					if err != nil {
						panic(err)
					}
					return db
				}(),
				tickID: 0,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTicker(tt.args.tx, tt.args.tickID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTicker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTicker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTickers(t *testing.T) {
	db, err := connectMYSQL(TestMYSQLConfig)
	if err != nil {
		panic(err)
	}

	ticker := NewTicker(rand.Int(), "BTC_USD", "RUNNING", time.Now().Unix(), 100.0, 200.0, 0.1, 0.2, 1000.0, 2000.0, 10.0, 20.0, 150.0, 10000.0, 1000.0)
	if err := ticker.Insert(db); err != nil {
		panic(err)
	}

	type args struct {
		tx *gorm.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test GetTickers",
			args: args{
				tx: func() *gorm.DB {
					db, err := connectMYSQL(TestMYSQLConfig)
					if err != nil {
						panic(err)
					}
					return db
				}(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTickers(tt.args.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTickers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf("GetTickers() = %v, want not empty", got)
			}
		})
	}
}
