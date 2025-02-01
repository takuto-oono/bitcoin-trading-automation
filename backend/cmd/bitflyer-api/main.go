package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/bitcoin-trading-automation/internal/bitflyer-api/router"
	"github.com/bitcoin-trading-automation/internal/config"
)

func main() {
	tomlFilePath := flag.String("toml", "../../toml/local.toml", "tomlファイルの名前")
	envFilePath := flag.String("env", "../../env/.env.local", "envファイルのパス")
	flag.Parse()

	cfg := config.NewConfig(*tomlFilePath, *envFilePath)

	r := router.NewRouter(cfg)

	u, err := url.Parse(cfg.Url.BitFlyerAPI)
	if err != nil {
		panic(err)
	}

	if err := r.Run(fmt.Sprintf(":%s", u.Port())); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
