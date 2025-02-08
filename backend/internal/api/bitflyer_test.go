package api

import (
	"reflect"
	"testing"

	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/models"
)

func TestAPI_BitFlyerAPIHealthCheck(t *testing.T) {
	type fields struct {
		Config config.Config
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test BitFlyer API Health Check",
			fields: fields{
				Config: TestAPIConfig,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				Config: tt.fields.Config,
			}
			if err := api.BitFlyerAPIHealthCheck(); (err != nil) != tt.wantErr {
				t.Errorf("API.BitFlyerAPIHealthCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAPI_GetTicker(t *testing.T) {
	type fields struct {
		Config config.Config
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test Get Ticker",
			fields: fields{
				Config: TestAPIConfig,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				Config: tt.fields.Config,
			}
			got, err := api.GetTicker()
			if (err != nil) != tt.wantErr {
				t.Errorf("API.GetTicker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, models.Ticker{}) {
				t.Errorf("API.GetTicker() got = %v, want not zero value", got)
			}
		})
	}
}
