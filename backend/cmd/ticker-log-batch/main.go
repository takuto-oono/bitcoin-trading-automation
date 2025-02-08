package main

import (
	"flag"
	"log"
	"sync"
	"time"

	"github.com/bitcoin-trading-automation/internal/api"
	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/models"
)

// go run cmd/ticker-log-batch/main.go -toml toml/local.toml -env env/.env.local
func main() {
	tomlFilePath := flag.String("toml", "../../toml/local.toml", "tomlファイルの名前")
	envFilePath := flag.String("env", "../../env/.env.local", "envファイルのパス")
	flag.Parse()

	cfg := config.NewConfig(*tomlFilePath, *envFilePath)
	api := api.NewAPI(cfg)

	c := make(chan models.Ticker)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		getTicker(api, c)
	}()

	go func() {
		defer wg.Done()
		postTicker(api, c)
	}()

	wg.Wait()
}

func getTicker(api *api.API, c chan models.Ticker) {
	for {
		ticker, err := api.GetTicker()
		if err != nil {
			log.Printf("Failed to get ticker: %v", err)
			continue
		}

		c <- ticker

		time.Sleep(2 * time.Second)
	}
}

func postTicker(api *api.API, c chan models.Ticker) {
	for {
		ticker := <-c
		if err := api.TickerLogPostTicker(ticker); err != nil {
			log.Printf("Failed to post ticker: %v, ticker: %v", err, ticker)
		}
		log.Printf("Success to post ticker: %v", ticker)
	}
}
