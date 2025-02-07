package usecase

import (
	"net/http"
	"testing"
)

func TestBitflyerUseCase_GetBoard(t *testing.T) {
	type args struct {
		productCode string
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "TestBitflyerUseCase_GetBoard",
			args: args{
				productCode: ProductCodeBTCJPY,
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "TestBitflyerUseCase_GetBoard_No_Product_Code",
			args: args{
				productCode: "",
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "TestBitflyerUseCase_GetBoard_Invalid_Product_Code",
			args: args{
				productCode: "INVALID",
			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(TestUseCaseConfig)
			_, statusCode, err := bu.GetBoard(tt.args.productCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("BitflyerUseCase.GetBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if statusCode != tt.wantStatusCode {
				t.Errorf("BitflyerUseCase.GetBoard() = %v, want %v", statusCode, tt.wantStatusCode)
			}
		})
	}
}

func TestBitflyerUseCase_GetTicker(t *testing.T) {
	type args struct {
		productCode string
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "TestBitflyerUseCase_GetTicker",
			args: args{
				productCode: ProductCodeBTCJPY,
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "TestBitflyerUseCase_GetTicker_No_Product_Code",
			args: args{
				productCode: "",
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "TestBitflyerUseCase_GetTicker_Invalid_Product_Code",
			args: args{
				productCode: "INVALID",
			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(TestUseCaseConfig)
			_, statusCode, err := bu.GetTicker(tt.args.productCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("BitflyerUseCase.GetTicker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if statusCode != tt.wantStatusCode {
				t.Errorf("BitflyerUseCase.GetTicker() = %v, want %v", statusCode, tt.wantStatusCode)
			}
		})
	}
}

func TestBitflyerUseCase_GetExecutions(t *testing.T) {
	type args struct {
		productCode string
		count       string
		before      string
		after       string
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantErr        bool
		lenGot         int
	}{
		{
			name: "TestBitflyerUseCase_GetExecutions",
			args: args{
				productCode: ProductCodeBTCJPY,
				count:       "10",
				before:      "",
				after:       "",
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
			lenGot:         10,
		},
		{
			name: "TestBitflyerUseCase_GetExecutions_No_Product_Code",
			args: args{
				productCode: "",
				count:       "10",
				before:      "",
				after:       "",
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
			lenGot:         10,
		},
		{
			name: "TestBitflyerUseCase_GetExecutions_Invalid_Product_Code",
			args: args{
				productCode: "INVALID",
				count:       "10",
				before:      "",
				after:       "",
			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
			lenGot:         0,
		},
		{
			name: "TestBitflyerUseCase_GetExecutions_No_Count",
			args: args{
				productCode: ProductCodeBTCJPY,
				count:       "",
				before:      "",
				after:       "",
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
			lenGot:         100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(TestUseCaseConfig)
			got, statusCode, err := bu.GetExecutions(tt.args.productCode, tt.args.count, tt.args.before, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("BitflyerUseCase.GetExecutions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if statusCode != tt.wantStatusCode {
				t.Errorf("BitflyerUseCase.GetExecutions() = %v, want %v", statusCode, tt.wantStatusCode)
			}
			if len(got) != tt.lenGot {
				t.Errorf("BitflyerUseCase.GetExecutions() = %v, want %v", len(got), tt.lenGot)
			}
		})
	}
}

func TestBitflyerUseCase_GetBoardState(t *testing.T) {
	type args struct {
		productCode string
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "TestBitflyerUseCase_GetBoardState",
			args: args{
				productCode: ProductCodeBTCJPY,
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "TestBitflyerUseCase_GetBoardState_No_Product_Code",
			args: args{
				productCode: "",
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "TestBitflyerUseCase_GetBoardState_Invalid_Product_Code",
			args: args{
				productCode: "INVALID",
			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(TestUseCaseConfig)
			_, statusCode, err := bu.GetBoardState(tt.args.productCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("BitflyerUseCase.GetBoardState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if statusCode != tt.wantStatusCode {
				t.Errorf("BitflyerUseCase.GetBoardState() = %v, want %v", statusCode, tt.wantStatusCode)
			}
		})
	}
}

func TestBitflyerUseCase_GetHealth(t *testing.T) {
	type args struct {
		productCode string
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "TestBitflyerUseCase_GetHealth",
			args: args{
				productCode: ProductCodeBTCJPY,
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "TestBitflyerUseCase_GetHealth_No_Product_Code",
			args: args{
				productCode: "",
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "TestBitflyerUseCase_GetHealth_Invalid_Product_Code",
			args: args{
				productCode: "INVALID",
			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "TestBitflyerUseCase_GetHealth_Other_Product_Code",
			args: args{
				productCode: "ETH_JPY",
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(TestUseCaseConfig)
			_, statusCode, err := bu.GetHealth(tt.args.productCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("BitflyerUseCase.GetHealth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if statusCode != tt.wantStatusCode {
				t.Errorf("BitflyerUseCase.GetHealth() = %v, want %v", statusCode, tt.wantStatusCode)
			}
		})
	}
}

func TestBitflyerUseCase_GetBalance(t *testing.T) {
	tests := []struct {
		name           string
		wantStatusCode int
		wantErr        bool
	}{
		{
			name:           "TestBitflyerUseCase_GetBalance",
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(TestUseCaseConfig)
			got, statusCode, err := bu.GetBalance()
			if (err != nil) != tt.wantErr {
				t.Errorf("BitflyerUseCase.GetBalance() error = %v, wantErr %v", err, tt.wantErr)
			}

			if statusCode != tt.wantStatusCode {
				t.Errorf("BitflyerUseCase.GetBalance() = %v, want %v", statusCode, tt.wantStatusCode)
			}

			if len(got) == 0 {
				t.Errorf("BitflyerUseCase.GetBalance() = %v, want not empty", got)
			}
		})
	}
}

func TestBitflyerUseCase_GetCollateral(t *testing.T) {
	tests := []struct {
		name           string
		wantStatusCode int
		wantErr        bool
	}{
		{
			name:    "TestBitflyerUseCase_GetCollateral",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(TestUseCaseConfig)
			_, statusCode, err := bu.GetCollateral()
			if (err != nil) != tt.wantErr {
				t.Errorf("BitflyerUseCase.GetCollateral() error = %v, wantErr %v", err, tt.wantErr)
			}
			if statusCode != http.StatusOK {
				t.Errorf("BitflyerUseCase.GetCollateral() = %v, want %v", statusCode, http.StatusOK)
			}
		})
	}
}

func TestBitflyerUseCase_PostSendChildOrder(t *testing.T) {
	type args struct {
		productCode    string
		ChildOrderType string
		side           string
		price          int
		size           float64
		MinuteToExpire int
		TimeInForce    string
		isDry          bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestBitflyerUseCase_PostSendChildOrder",
			args: args{
				productCode:    ProductCodeBTCJPY,
				ChildOrderType: "LIMIT",
				side:           "BUY",
				price:          1000000,
				size:           0.01,
				MinuteToExpire: 1,
				TimeInForce:    "GTC",
				isDry:          true,
			},
			wantErr: false,
		},
		{
			name: "TestBitflyerUseCase_PostSendChildOrder_No_Product_Code",
			args: args{
				productCode:    "",
				ChildOrderType: "LIMIT",
				side:           "BUY",
				price:          1000000,
				size:           0.01,
				MinuteToExpire: 1,
				TimeInForce:    "GTC",
				isDry:          true,
			},
			wantErr: false,
		},
		{
			name: "TestBitflyerUseCase_PostSendChildOrder_Invalid_Product_Code",
			args: args{
				productCode:    "INVALID",
				ChildOrderType: "LIMIT",
				side:           "BUY",
				price:          1000000,
				size:           0.01,
				MinuteToExpire: 1,
				TimeInForce:    "GTC",
				isDry:          true,
			},
			wantErr: true,
		},
		{
			name: "TestBitflyerUseCase_PostSendChildOrder_No_ChildOrderType",
			args: args{
				productCode:    ProductCodeBTCJPY,
				ChildOrderType: "",
				side:           "BUY",
				price:          1000000,
				size:           0.01,
				MinuteToExpire: 1,
				TimeInForce:    "GTC",
				isDry:          true,
			},
			wantErr: true,
		},
		{
			name: "TestBitflyerUseCase_PostSendChildOrder_No_Side",
			args: args{
				productCode:    ProductCodeBTCJPY,
				ChildOrderType: "LIMIT",
				side:           "",
				price:          1000000,
				size:           0.01,
				MinuteToExpire: 1,
				TimeInForce:    "GTC",
				isDry:          true,
			},
			wantErr: true,
		},
		{
			name: "TestBitflyerUseCase_PostSendChildOrder_No_TimeInForce",
			args: args{
				productCode:    ProductCodeBTCJPY,
				ChildOrderType: "LIMIT",
				side:           "BUY",
				price:          1000000,
				size:           0.01,
				MinuteToExpire: 1,
				TimeInForce:    "",
				isDry:          true,
			},
			wantErr: true,
		},
		{
			name: "TestBitflyerUseCase_PostSendChildOrder_Invalid_TimeInForce",
			args: args{
				productCode:    ProductCodeBTCJPY,
				ChildOrderType: "LIMIT",
				side:           "BUY",
				price:          1000000,
				size:           0.01,
				MinuteToExpire: 1,
				TimeInForce:    "INVALID",
				isDry:          true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(TestUseCaseConfig)
			_, statusCode, err := bu.PostSendChildOrder(tt.args.productCode, tt.args.ChildOrderType, tt.args.side, tt.args.price, tt.args.size, tt.args.MinuteToExpire, tt.args.TimeInForce, tt.args.isDry)
			if (err != nil) != tt.wantErr {
				t.Errorf("BitflyerUseCase.PostSendChildOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if statusCode == 0 {
				t.Errorf("BitflyerUseCase.PostSendChildOrder() = %v, want not 0", statusCode)
			}
		})
	}
}

func TestBitflyerUseCase_PostCancelChildOrder(t *testing.T) {
	type args struct {
		productCode  string
		ChildOrderID string
		isDry        bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestBitflyerUseCase_PostCancelChildOrder",
			args: args{
				productCode:  ProductCodeBTCJPY,
				ChildOrderID: "30", // テキトウな値
				isDry:        true,
			},
			wantErr: false,
		},
		{
			name: "TestBitflyerUseCase_PostCancelChildOrder_No_Product_Code",
			args: args{
				productCode:  "",
				ChildOrderID: "30", // テキトウな値
				isDry:        true,
			},
			wantErr: false,
		},
		{
			name: "TestBitflyerUseCase_PostCancelChildOrder_Invalid_Product_Code",
			args: args{
				productCode:  "INVALID",
				ChildOrderID: "30", // テキトウな値
				isDry:        true,
			},
			wantErr: true,
		},
		{
			name: "TestBitflyerUseCase_PostCancelChildOrder_No_ChildOrderID",
			args: args{
				productCode:  ProductCodeBTCJPY,
				ChildOrderID: "",
				isDry:        true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(TestUseCaseConfig)
			statusCode, err := bu.PostCancelChildOrder(tt.args.productCode, tt.args.ChildOrderID, tt.args.isDry)
			if (err != nil) != tt.wantErr {
				t.Errorf("BitflyerUseCase.PostCancelChildOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
			if statusCode == 0 {
				t.Errorf("BitflyerUseCase.PostCancelChildOrder() = %v, want not 0", statusCode)
			}
		})
	}
}

func TestBitflyerUseCase_GetChildOrders(t *testing.T) {
	tests := []struct {
		name           string
		wantStatusCode int
		wantErr        bool
	}{
		{
			name:           "TestBitflyerUseCase_GetChildOrders",
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(TestUseCaseConfig)
			got, statusCode, err := bu.GetChildOrders()
			if err != nil {
				if !tt.wantErr {
					t.Errorf("BitflyerUseCase.GetChildOrders() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if statusCode != tt.wantStatusCode {
				t.Errorf("BitflyerUseCase.GetChildOrders() = %v, want %v", statusCode, tt.wantStatusCode)
			}
			if len(got) == 0 {
				t.Errorf("BitflyerUseCase.GetChildOrders() = %v, want not empty", got)
			}
		})
	}
}

func Test_validateProductCode(t *testing.T) {
	type args struct {
		productCode string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_validateProductCode",
			args: args{
				productCode: ProductCodeBTCJPY,
			},
			want: true,
		},
		{
			name: "Test_validateProductCode_Invalid",
			args: args{
				productCode: "INVALID",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateProductCode(tt.args.productCode); got != tt.want {
				t.Errorf("validateProductCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
