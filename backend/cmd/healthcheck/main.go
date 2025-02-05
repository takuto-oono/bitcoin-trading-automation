package main

import (
	"flag"
	"log"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/bitcoin-trading-automation/internal/api"
	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/models"
)

const (
	StatusOK = "OK"
	StatusNG = "NG"

	BitFlyerAPI       = "bitflyer-api"
	RedisServer       = "redis-server"
	SlackNotification = "slack-notification"
	TickerLog         = "ticker-log"

	SlackChannel = "#bitcoin-trading-log"

	HealthCheckInterval  = 2 * time.Hour
	FirstHealthCheckTime = 15 * time.Second
)

// go run cmd/healthcheck/main.go -toml toml/local.toml -env env/.env.local
func main() {
	tomlFilePath := flag.String("toml", "../../toml/local.toml", "tomlファイルの名前")
	envFilePath := flag.String("env", "../../env/.env.local", "envファイルのパス")
	flag.Parse()

	api := api.NewAPI(config.NewConfig(*tomlFilePath, *envFilePath))

	time.Sleep(FirstHealthCheckTime)
	log.Println("Start Health Check")
	runHealthChecks(api)
	log.Println("End Health Check")

	ticker := time.NewTicker(HealthCheckInterval)

	// テスト用のticker
	// ticker := time.NewTicker(2 * time.Second)

	defer ticker.Stop()

	for range ticker.C {
		runHealthChecks(api)
	}
}

func runHealthChecks(api *api.API) {
	result := healthCheckServers(api)

	resultStr := resultToString(result)

	if err := api.SlackNotificationPostMessage(models.SlackNotificationPostMessage{Text: resultStr, Channel: SlackChannel}); err != nil {
		log.Printf("Failed to post message to Slack: %v", err)
	}
}

func healthCheckServers(api *api.API) map[string]string {
	results := make(map[string]string)
	var mu sync.Mutex
	wg := sync.WaitGroup{}

	healthChecks := []struct {
		name string
		fn   func() error
	}{
		{BitFlyerAPI, api.BitFlyerAPIHealthCheck},
		{RedisServer, api.RedisServerHealthCheck},
		{SlackNotification, api.SlackNotificationHealthCheck},
		{TickerLog, api.TickerLogHealthCheck},
	}

	wg.Add(len(healthChecks))

	for _, hc := range healthChecks {
		go func(name string, fn func() error) {
			defer wg.Done()
			s := status(fn() != nil)

			mu.Lock()
			results[name] = s
			mu.Unlock()
		}(hc.name, hc.fn)
	}

	wg.Wait()
	return results
}

func resultToString(results map[string]string) string {
	keys := make([]string, 0, len(results))
	for k := range results {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var sb strings.Builder

	sb.WriteString("Health Check Results\n")
	for _, k := range keys {
		sb.WriteString(k)
		sb.WriteString(": ")
		sb.WriteString(results[k])
		sb.WriteString("\n")
	}
	sb.WriteString("++++++++++++++++++++\n")

	return sb.String()
}

func status(hasErr bool) string {
	if hasErr {
		return StatusNG
	}
	return StatusOK
}
