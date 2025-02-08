package api

import (
	"testing"

	"github.com/bitcoin-trading-automation/internal/config"
)

func TestAPI_RedisServerHealthCheck(t *testing.T) {
	type fields struct {
		Config config.Config
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test Redis Server Health Check",
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
			if err := api.RedisServerHealthCheck(); (err != nil) != tt.wantErr {
				t.Errorf("API.RedisServerHealthCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
