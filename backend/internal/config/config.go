package config

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Config struct {
	Url      Url      `toml:"url"`
	BitFlyer BitFlyer `toml:"bitFlyer"`
	Slack    Slack
	Redis    Redis `toml:"redis"`
	MYSQL    MYSQL `toml:"mysql"`
}

type Url struct {
	BitFlyerAPI       string `toml:"bitFlyerAPI"`
	SlackNotification string `toml:"slackNotification"`
	RedisServer       string `toml:"redisServer"`
	TickerLogServer   string `toml:"tickerLogServer"`
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

type MYSQL struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	DbName   string `toml:"db"`
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

	if gin.Mode() == gin.TestMode {
		config.MYSQL.DbName = "test_" + config.MYSQL.DbName
	}

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
	if cfg.Url.TickerLogServer == "" {
		return errors.New("tickerLogServer is empty")
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
	if cfg.MYSQL.User == "" {
		return errors.New("user is empty")
	}
	if cfg.MYSQL.Password == "" {
		return errors.New("password is empty")
	}
	if cfg.MYSQL.Host == "" {
		return errors.New("host is empty")
	}
	if cfg.MYSQL.Port == "" {
		return errors.New("port is empty")
	}
	if cfg.MYSQL.DbName == "" {
		return errors.New("dbname is empty")
	}
	return nil

}
