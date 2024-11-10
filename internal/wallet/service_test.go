package wallet

import (
	"github.com/google/uuid"
	"testing"
)

type MockWalletRepository struct{}

func (m *MockWalletRepository) Get(id uuid.UUID) (*Wallet, error) {
	return &Wallet{
		ID:      id,
		Balance: 10,
	}, nil
}
func (m *MockWalletRepository) Update(wallet *Wallet) error {
	return nil
}

func TestGetBalance(t *testing.T) {
	ws := NewService(&MockWalletRepository{})

	res, err := ws.GetBalance(uuid.New())
	if err != nil {
		t.Error(err)
	}
	if res != 10 {
		t.Error("wrong result")
	}
}

func TestService_CalcBalance(t *testing.T) {
	type args struct {
		balance float64
		opType  OperationType
		amount  float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "correct_1",
			args: args{
				balance: 10,
				opType:  Deposit,
				amount:  15,
			},
			want:    25,
			wantErr: false,
		},
		{
			name: "correct_2",
			args: args{
				balance: 10,
				opType:  Withdraw,
				amount:  5,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "negative balance",
			args: args{
				balance: 10,
				opType:  Withdraw,
				amount:  15,
			},
			want:    10,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(&MockWalletRepository{})
			got, err := s.CalcBalance(tt.args.balance, tt.args.opType, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalcBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalcBalance() got = %v, want %v", got, tt.want)
			}
		})
	}
}
