package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/router"
)

// go run cmd/bitflyer-api/main.go -toml toml/local.toml -env env/.env.local
func main() {
	tomlFilePath := flag.String("toml", "../../toml/local.toml", "tomlファイルの名前")
	envFilePath := flag.String("env", "../../env/.env.local", "envファイルのパス")
	flag.Parse()

	cfg := config.NewConfig(*tomlFilePath, *envFilePath)

	r := router.NewBitFlyerRouter(cfg)

	u, err := url.Parse(cfg.Url.BitFlyerAPI)
	if err != nil {
		panic(err)
	}

	if err := r.Run(fmt.Sprintf(":%s", u.Port())); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
