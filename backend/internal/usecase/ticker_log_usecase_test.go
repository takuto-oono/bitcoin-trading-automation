package usecase

import (
	"math/rand"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/bitcoin-trading-automation/internal/bitflyer-api/api/models"
	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/mysql"
)

func TestTickerLog_GetTickerLogs(t *testing.T) {
	m, err := mysql.NewMYSQL(TestUseCaseConfig)
	if err != nil {
		panic(err)
	}

	ticker := mysql.NewTicker(rand.Int(), "BTC_JPY", "RUNNING", time.Now().Unix(), 1000000, 1000000, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 1000000, 1000000, 1000000)

	if err := ticker.Insert(m.DB); err != nil {
		panic(err)
	}

	type fields struct {
		Config config.Config
		MYSQL  mysql.MYSQL
	}
	tests := []struct {
		name       string
		fields     fields
		statusCode int
		wantErr    bool
	}{
		{
			name: "Test GetTickerLogs",
			fields: fields{
				Config: TestUseCaseConfig,
				MYSQL: func() mysql.MYSQL {
					mysql, err := mysql.NewMYSQL(TestUseCaseConfig)
					if err != nil {
						panic(err)
					}
					return *mysql
				}(),
			},
			statusCode: http.StatusOK,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TickerLog{
				Config: tt.fields.Config,
				MYSQL:  tt.fields.MYSQL,
			}
			_, statusCode, err := tr.GetTickerLogs()
			if (err != nil) != tt.wantErr {
				t.Errorf("TickerLog.GetTickerLogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if statusCode != tt.statusCode {
				t.Errorf("TickerLog.GetTickerLogs() = %v, want %v", statusCode, tt.statusCode)
			}
		})
	}
}

func TestTickerLog_GetTickerLogByTickID(t *testing.T) {
	m, err := mysql.NewMYSQL(TestUseCaseConfig)
	if err != nil {
		panic(err)
	}

	ticker := mysql.NewTicker(rand.Int(), "BTC_JPY", "RUNNING", time.Now().Unix(), 1000000, 1000000, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 1000000, 1000000, 1000000)

	if err := ticker.Insert(m.DB); err != nil {
		panic(err)
	}

	type fields struct {
		Config config.Config
		MYSQL  mysql.MYSQL
	}
	type args struct {
		tickerID int
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		want           *mysql.Ticker
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "Test GetTickerLogByTickID",
			fields: fields{
				Config: TestUseCaseConfig,
				MYSQL: func() mysql.MYSQL {
					mysql, err := mysql.NewMYSQL(TestUseCaseConfig)
					if err != nil {
						panic(err)
					}
					return *mysql
				}(),
			},
			args: args{
				tickerID: ticker.TickID,
			},
			want:           ticker,
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "Test GetTickerLogByTickID",
			fields: fields{
				Config: TestUseCaseConfig,
				MYSQL: func() mysql.MYSQL {
					mysql, err := mysql.NewMYSQL(TestUseCaseConfig)
					if err != nil {
						panic(err)
					}
					return *mysql
				}(),
			},
			args: args{
				tickerID: 0,
			},
			want:           nil,
			wantStatusCode: http.StatusNotFound,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TickerLog{
				Config: tt.fields.Config,
				MYSQL:  tt.fields.MYSQL,
			}
			got, statusCode, err := tr.GetTickerLogByTickID(tt.args.tickerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TickerLog.GetTickerLogByTickID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if statusCode != tt.wantStatusCode {
				t.Errorf("TickerLog.GetTickerLogByTickID() = %v, want %v", statusCode, tt.wantStatusCode)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TickerLog.GetTickerLogByTickID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTickerLog_PostTickerLog(t *testing.T) {
	type fields struct {
		Config config.Config
		MYSQL  mysql.MYSQL
	}
	type args struct {
		ticker models.Ticket
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "Test PostTickerLog",
			fields: fields{
				Config: TestUseCaseConfig,
				MYSQL: func() mysql.MYSQL {
					mysql, err := mysql.NewMYSQL(TestUseCaseConfig)
					if err != nil {
						panic(err)
					}
					return *mysql
				}(),
			},
			args: args{
				ticker: models.Ticket{
					TickID:          rand.Int(),
					ProductCode:     "BTC_JPY",
					State:           "RUNNING",
					Timestamp:       "2006-01-02T15:04:05.000",
					BestBid:         1000000,
					BestAsk:         1000000,
					BestBidSize:     0.1,
					BestAskSize:     0.1,
					TotalBidDepth:   0.1,
					TotalAskDepth:   0.1,
					MarketBidSize:   0.1,
					MarketAskSize:   0.1,
					Ltp:             1000000,
					Volume:          1000000,
					VolumeByProduct: 1000000,
				},
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "Test PostTickerLog invalid timestamp",
			fields: fields{
				Config: TestUseCaseConfig,
				MYSQL: func() mysql.MYSQL {
					mysql, err := mysql.NewMYSQL(TestUseCaseConfig)
					if err != nil {
						panic(err)
					}
					return *mysql
				}(),
			},
			args: args{
				ticker: models.Ticket{
					TickID:          rand.Int(),
					ProductCode:     "BTC_JPY",
					State:           "RUNNING",
					Timestamp:       "invalid",
					BestBid:         1000000,
					BestAsk:         1000000,
					BestBidSize:     0.1,
					BestAskSize:     0.1,
					TotalBidDepth:   0.1,
					TotalAskDepth:   0.1,
					MarketBidSize:   0.1,
					MarketAskSize:   0.1,
					Ltp:             1000000,
					Volume:          1000000,
					VolumeByProduct: 1000000,
				},
			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TickerLog{
				Config: tt.fields.Config,
				MYSQL:  tt.fields.MYSQL,
			}
			statusCode, err := tr.PostTickerLog(tt.args.ticker)
			if (err != nil) != tt.wantErr {
				t.Errorf("TickerLog.PostTickerLog() error = %v, wantErr %v", err, tt.wantErr)
			}
			if statusCode != tt.wantStatusCode {
				t.Errorf("TickerLog.PostTickerLog() = %v, want %v", statusCode, tt.wantStatusCode)
			}
		})
	}
}

func Test_parseTimestamp(t *testing.T) {
	type args struct {
		timestamp string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Test parseTimestamp",
			args: args{
				timestamp: "2006-01-02T15:04:05.000",
			},
			want:    1136214245,
			wantErr: false,
		},
		{
			name: "Test parseTimestamp invalid",
			args: args{
				timestamp: "invalid",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTimestamp(tt.args.timestamp)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTimestamp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
