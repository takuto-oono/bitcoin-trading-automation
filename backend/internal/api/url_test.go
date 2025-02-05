package api

import "testing"

const TestURL = "http://hogehoge:8000"

func TestBitFlyerAPIHealthCheck(t *testing.T) {
	expected := TestURL + "/healthcheck"
	got, err := BitFlyerAPI(TestURL).HealthCheck()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}

func TestBitFlyerAPIGetTicker(t *testing.T) {
	expected := TestURL + "/ticker"
	got, err := BitFlyerAPI(TestURL).GetTicker()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}

func TestRedisServerHealthCheck(t *testing.T) {
	expected := TestURL + "/healthcheck"
	got, err := RedisServer(TestURL).HealthCheck()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}

func TestSlackNotificationHealthCheck(t *testing.T) {
	expected := TestURL + "/healthcheck"
	got, err := SlackNotification(TestURL).HealthCheck()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}

func TestSlackNotificationPostMessage(t *testing.T) {
	expected := TestURL + "/message"
	got, err := SlackNotification(TestURL).PostMessage()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}

func TestTickerLogServerHealthCheck(t *testing.T) {
	expected := TestURL + "/healthcheck"
	got, err := TickerLogServer(TestURL).HealthCheck()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}

func TestTickerLogServerPostTickerLog(t *testing.T) {
	expected := TestURL + "/ticker-logs"
	got, err := TickerLogServer(TestURL).PostTickerLog()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
