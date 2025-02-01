package slackapi

import (
	"testing"
	"time"
)

const TestSlackChannel = "bitcoin-trading-log"

func TestSlackAPI_PostMessage(t *testing.T) {
	type args struct {
		channel string
		text    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				channel: TestSlackChannel,
				text:    "slack api server unit test " + time.Now().String(),
			},
			wantErr: false,
		},
		{
			name: "Empty Channel", // エラーにはならないがslackにメッセージは送信されない
			args: args{
				channel: "",
				text:    "slack api server unit test " + time.Now().String(),
			},
			wantErr: false,
		},
		{
			name: "Empty Text", // エラーにはならないがslackにメッセージは送信されない
			args: args{
				channel: TestSlackChannel,
				text:    "",
			},
			wantErr: false,
		},
		{
			name: "worng channel", // エラーにはならないがslackにメッセージは送信されない
			args: args{
				channel: "wrong-channel",
				text:    "slack api server unit test " + time.Now().String(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SlackAPI{
				Config: SlackAPITestCfg,
			}
			if err := s.PostMessage(tt.args.channel, tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("SlackAPI.PostMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
