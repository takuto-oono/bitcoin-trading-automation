package slackapi

import "github.com/bitcoin-trading-automation/internal/config"

var SlackAPITestCfg config.Config

func init() {
	SlackAPITestCfg = config.NewConfig("../../../toml/local.toml", "../../../env/.env.local")
}
