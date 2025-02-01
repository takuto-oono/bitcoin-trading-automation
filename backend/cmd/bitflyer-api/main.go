package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bitcoin-trading-automation/internal/bitflyer-api/router"
	"github.com/bitcoin-trading-automation/internal/config"
)

func main() {
	tomlFilePath := flag.String("toml", "../../toml/local.toml", "tomlファイルの名前")
	envFilePath := flag.String("env", "../../env/.env.local", "envファイルのパス")
	flag.Parse()

	cfg := config.NewConfig(*tomlFilePath, *envFilePath)

	r := router.NewRouter(cfg)

	if err := r.Run(fmt.Sprintf(":%s", cfg.BaseConfig.Port)); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
