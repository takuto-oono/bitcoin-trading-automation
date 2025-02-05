package api

import (
	"fmt"
	"net/url"
)

type BitFlyerAPI string
type RedisServer string
type SlackNotification string
type TickerLogServer string

func (b BitFlyerAPI) HealthCheck() (string, error) {
	u, err := url.Parse(fmt.Sprintf("%s/healthcheck", b))
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func (b BitFlyerAPI) GetTicker() (string, error) {
	u, err := url.Parse(fmt.Sprintf("%s/ticker", b))
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func (r RedisServer) HealthCheck() (string, error) {
	u, err := url.Parse(fmt.Sprintf("%s/healthcheck", r))
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func (s SlackNotification) HealthCheck() (string, error) {
	u, err := url.Parse(fmt.Sprintf("%s/healthcheck", s))
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func (s SlackNotification) PostMessage() (string, error) {
	u, err := url.Parse(fmt.Sprintf("%s/message", s))
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func (t TickerLogServer) PostTickerLog() (string, error) {
	u, err := url.Parse(fmt.Sprintf("%s/ticker-logs", t))
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func (t TickerLogServer) HealthCheck() (string, error) {
	u, err := url.Parse(fmt.Sprintf("%s/healthcheck", t))
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
