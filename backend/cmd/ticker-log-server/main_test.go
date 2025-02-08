package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
	"testing"
	"time"

	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/mysql"
)

func TestMainFunction(t *testing.T) {
	cfg := config.NewConfig("../../toml/local.toml", "../../env/.env.local")

	m, err := mysql.NewMYSQL(cfg)
	if err != nil {
		panic(err)
	}

	ticker := mysql.NewTicker(1, "BTC_JPY", "RUNNING", time.Now().Unix(), 1000000, 1000000, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 1000000, 1000000, 1000000)

	if err := ticker.Insert(m.DB); err != nil {
		panic(err)
	}

	type args struct {
		curl []string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "health check",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "http://localhost:8004/healthcheck"},
			},
		},
		{
			name: "get ticker logs",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "http://localhost:8004/ticker-logs"},
			},
		},
		{
			name: "get ticker log by ID",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "http://localhost:8004/ticker-logs/" + strconv.Itoa(ticker.TickID)},
			},
		},
		{
			name: "post ticker log",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "-X", "POST", "-H", "Content-Type: application/json", "-d", fmt.Sprintf(`{"id": %v, "product_code": "BTC_JPY", "state": "RUNNING", "timestamp": 1738687615, "best_bid": 1000000, "best_ask": 1000000, "best_bid_size": 0.1, "best_ask_size": 0.1, "total_bid_depth": 0.1, "total_ask_depth": 0.1, "market_bid_size": 0.1, "market_ask_size": 0.1, "ltp": 1000000, "volume": 1000000, "volume_by_product": 1000000}`, rand.Int()), "http://localhost:8004/ticker-logs"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// curlコマンドを実行
			cmd := exec.Command(tt.args.curl[0], tt.args.curl[1:]...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("Failed to execute curl command: %v", err)
			}

			// ステータスコードを確認
			statusCode := string(output[len(output)-3:]) // 最後の3文字がステータスコード
			if statusCode != "200" && statusCode != "201" {
				t.Fatalf("Server returned non-200/201 status code: %s", statusCode)
			}
		})
	}
}
