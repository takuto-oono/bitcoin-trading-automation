package main

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/router"
)

func main() {
	tomlFilePath := flag.String("toml", "../../toml/local.toml", "tomlファイルの名前")
	envFilePath := flag.String("env", "../../env/.env.local", "envファイルのパス")
	flag.Parse()

	cfg := config.NewConfig(*tomlFilePath, *envFilePath)

	r, err := router.NewRedisServerRouter(cfg)
	if err != nil {
		panic(err)
	}

	u, err := url.Parse(cfg.Url.RedisServer)
	if err != nil {
		panic(err)
	}

	if err := r.Run(fmt.Sprintf(":%s", u.Port())); err != nil {
		panic(err)
	}
}
