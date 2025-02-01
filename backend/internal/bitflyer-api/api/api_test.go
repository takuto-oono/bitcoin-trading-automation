package api

import "github.com/bitcoin-trading-automation/internal/config"

var APITestConfig config.Config

func init() {
	APITestConfig = config.NewConfig("../../../toml/local.toml", "../../../env/.env.local")
}
