package api

import (
	"reflect"
	"testing"

	"github.com/bitcoin-trading-automation/internal/config"
)

var TestAPIConfig config.Config

func init() {
	TestAPIConfig = config.NewConfig("../../toml/local.toml", "../../env/.env.local")
}

func Test_marshalJson(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Test marshalJson",
			args: args{
				v: map[string]interface{}{
					"key": "value",
				},
			},
			want:    []byte("{\"key\":\"value\"}"),
			wantErr: false,
		},
		{
			name: "Test marshalJson is nil",
			args: args{
				v: nil,
			},
			want:    []byte{},
			wantErr: false,
		},
		{
			name: "Test marshalJson is empty",
			args: args{
				v: "",
			},
			want:    []byte{},
			wantErr: false,
		},
		{
			name: "Test marshalJson is emptm map",
			args: args{
				v: map[string]interface{}{},
			},
			want:    []byte("{}"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := marshalJson(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("marshalJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("marshalJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractPath(t *testing.T) {
	type args struct {
		u string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test extractPath",
			args: args{
				u: "https://api.bitflyer.com/v1/getmarkets",
			},
			want:    "/v1/getmarkets",
			wantErr: false,
		},
		{
			name: "Test extractPath is empty",
			args: args{
				u: "",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Test extractPath is no path",
			args: args{
				u: "https://api.bitflyer.com",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractPath(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
