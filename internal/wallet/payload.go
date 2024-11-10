package wallet

import "github.com/google/uuid"

type OperationType string

const (
	Deposit  OperationType = "DEPOSIT"
	Withdraw OperationType = "WITHDRAW"
)

type OperationRequest struct {
	WalletID      uuid.UUID     `json:"walletUuid"`
	OperationType OperationType `json:"operationType"`
	Amount        float64       `json:"amount"`
}

type GetRequest struct {
	WalletID uuid.UUID `json:"walletUuid"`
}

type GetResponse struct {
	Balance float64 `json:"balance"`
}
