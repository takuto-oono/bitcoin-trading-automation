package slackapi

import (
	"errors"

	"github.com/bitcoin-trading-automation/internal/config"
)

type SlackAPI struct {
	Config config.Config
}

func NewSlackAPI(cfg *config.Config) (*SlackAPI, error) {
	if cfg.Slack.AccessToken == "" {
		return nil, errors.New("slack access token is empty")
	}

	return &SlackAPI{Config: *cfg}, nil
}
