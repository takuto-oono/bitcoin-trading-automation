package main

import (
	"flag"

	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/router"
)

func main() {
	tomlFilePath := flag.String("toml", "../../toml/local.toml", "tomlファイルの名前")
	envFilePath := flag.String("env", "../../env/.env.local", "envファイルのパス")
	flag.Parse()

	cfg := config.NewConfig(*tomlFilePath, *envFilePath)

	r, err := router.NewTickerLogServerRouter(cfg)
	if err != nil {
		panic(err)
	}

	if err := router.RunTickerLogServer(r, cfg); err != nil {
		panic(err)
	}
}
