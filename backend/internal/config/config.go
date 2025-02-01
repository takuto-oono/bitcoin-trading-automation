package config

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type Config struct {
	Url      Url      `toml:"url"`
	BitFlyer BitFlyer `toml:"bitFlyer"`
	Slack    Slack
	Redis    Redis `toml:"redis"`
}

type Url struct {
	BitFlyerAPI       string `toml:"bitFlyerAPI"`
	SlackNotification string `toml:"slackNotification"`
	RedisServer       string `toml:"redisServer"`
}

type BitFlyer struct {
	ApiKey       string
	ApiSecret    string
	BaseEndPoint string `toml:"baseEndPoint"`
}

type Slack struct {
	AccessToken string
}

type Redis struct {
	Address    string `toml:"address"`
	IndexCount int    `toml:"indexCount"`
}

func NewConfig(tomlFilePath, envFilePath string) Config {
	var config Config

	if _, err := toml.DecodeFile(tomlFilePath, &config); err != nil {
		panic(err)
	}

	if err := godotenv.Load(envFilePath); err != nil {
		panic(err)
	}

	config.BitFlyer.ApiKey = os.Getenv("BITFLYER_API_KEY")
	config.BitFlyer.ApiSecret = os.Getenv("BITFLYER_API_SECRET")

	config.Slack.AccessToken = os.Getenv("SLACK_ACCESS_TOKEN")

	if err := config.mustCheck(); err != nil {
		panic(err)
	}

	return config
}

func (cfg Config) mustCheck() error {
	if cfg.Url.BitFlyerAPI == "" {
		return errors.New("port is empty")
	}
	if cfg.Url.SlackNotification == "" {
		return errors.New("slackNotification is empty")
	}
	if cfg.Url.RedisServer == "" {
		return errors.New("redisServer is empty")
	}
	if cfg.BitFlyer.ApiKey == "" {
		return errors.New("apiKey is empty")
	}
	if cfg.BitFlyer.ApiSecret == "" {
		return errors.New("apiSecret is empty")
	}
	if cfg.Slack.AccessToken == "" {
		return errors.New("accessToken is empty")
	}
	if cfg.BitFlyer.BaseEndPoint == "" {
		return errors.New("baseEndPoint is empty")
	}
	if cfg.Redis.Address == "" {
		return errors.New("address is empty")
	}
	if cfg.Redis.IndexCount == 0 {
		return errors.New("indexCount is empty")
	}
	return nil

}
