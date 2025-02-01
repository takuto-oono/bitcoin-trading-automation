package models

import "github.com/bitcoin-trading-automation/internal/config"

var modelsTestConfig config.Config

const redisTestIndex = 0

func init() {
	modelsTestConfig = config.NewConfig("../../../toml/local.toml", "../../../env/.env.local")
}
