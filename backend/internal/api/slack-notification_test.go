package api

import (
	"testing"

	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/models"
)

func TestAPI_SlackNotificationHealthCheck(t *testing.T) {
	type fields struct {
		Config config.Config
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test SlackNotificationHealthCheck",
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
			if err := api.SlackNotificationHealthCheck(); (err != nil) != tt.wantErr {
				t.Errorf("API.SlackNotificationHealthCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAPI_SlackNotificationPostMessage(t *testing.T) {
	type fields struct {
		Config config.Config
	}
	type args struct {
		reqBody models.SlackNotificationPostMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test SlackNotificationPostMessage",
			fields: fields{
				Config: TestAPIConfig,
			},
			args: args{
				reqBody: models.SlackNotificationPostMessage{
					Channel: "test",
					Text:    "test",
				},
			},
			wantErr: false,
		},
		{
			name: "Test SlackNotificationPostMessage Invalid body",
			fields: fields{
				Config: TestAPIConfig,
			},
			args: args{
				reqBody: models.SlackNotificationPostMessage{
					Channel: "",
					Text:    "",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				Config: tt.fields.Config,
			}
			if err := api.SlackNotificationPostMessage(tt.args.reqBody); (err != nil) != tt.wantErr {
				t.Errorf("API.SlackNotificationPostMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
