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
}

type Url struct {
	BitFlyerAPI string `toml:"bitFlyerAPI"`
}

type BitFlyer struct {
	ApiKey       string
	ApiSecret    string
	BaseEndPoint string `toml:"baseEndPoint"`
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

	if err := config.mustCheck(); err != nil {
		panic(err)
	}

	return config
}

func (cfg Config) mustCheck() error {
	if cfg.Url.BitFlyerAPI == "" {
		return errors.New("port is empty")
	}
	if cfg.BitFlyer.ApiKey == "" {
		return errors.New("apiKey is empty")
	}
	if cfg.BitFlyer.ApiSecret == "" {
		return errors.New("apiSecret is empty")
	}
	if cfg.BitFlyer.BaseEndPoint == "" {
		return errors.New("baseEndPoint is empty")
	}
	return nil

}
