package usecase

import "github.com/bitcoin-trading-automation/internal/config"

var TestUseCaseConfig config.Config

func init() {
	TestUseCaseConfig = config.NewConfig("../../toml/local.toml", "../../env/.env.local")
}
