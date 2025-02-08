package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestMainFunction(t *testing.T) {
	type args struct {
		curl []string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "health check",
			args: args{
				curl: []string{"curl", "http://localhost:8002/healthcheck"},
			},
		},
		{
			name: "post slack message",
			args: args{
				curl: []string{"curl", "-X", "POST", "localhost:8002/message", "-H", "Content-Type: application/json", "-d", "{\"channel\": \"#bitcoin-trading-log\", \"text\": \"*integration test*\"}"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// curlコマンドを実行
			cmd := exec.Command("curl", tt.args.curl...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("Failed to execute curl command: %v", err)
			}

			// サーバーの応答を確認
			if !contains(string(output), "\"status\":\"ok\"") {
				t.Fatalf("Server is not functioning properly: %s", output)
			}
		})
	}

}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
