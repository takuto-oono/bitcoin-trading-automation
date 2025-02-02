package mysql

import (
	"testing"

	"github.com/bitcoin-trading-automation/internal/config"
)

var TestMYSQLConfig config.Config

func init() {
	TestMYSQLConfig = config.NewConfig("../../toml/local.toml", "../../env/.env.local")
	_, err := NewMYSQL(TestMYSQLConfig)
	if err != nil {
		panic(err)
	}
}

func TestNewMYSQL(t *testing.T) {
	type args struct {
		cfg config.Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test NewMYSQL",
			args: args{
				cfg: TestMYSQLConfig,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewMYSQL(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMYSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_connectMYSQL(t *testing.T) {
	type args struct {
		cfg config.Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test connectMYSQL",
			args: args{
				cfg: TestMYSQLConfig,
			},
			wantErr: false,
		},
		{
			name: "Test connectMYSQL",
			args: args{
				cfg: config.Config{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := connectMYSQL(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("connectMYSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
