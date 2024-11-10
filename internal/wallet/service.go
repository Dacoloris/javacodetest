package wallet

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type WalletService interface {
	CalcBalance(balance float64, operationType OperationType, amount float64) (float64, error)
	GetBalance(walletID uuid.UUID) (float64, error)
}

type Service struct {
	Repo IRepository
	mu   sync.Mutex
}

func NewService(walletRepository IRepository) WalletService {
	return &Service{Repo: walletRepository}
}

func (s *Service) GetBalance(id uuid.UUID) (float64, error) {
	wallet, err := s.Repo.Get(id)
	if err != nil {
		return 0, err
	}

	if wallet == nil {
		return 0, fmt.Errorf(ErrWalletDoesntExist)
	}

	return wallet.Balance, nil
}

func (s *Service) CalcBalance(balance float64, opType OperationType, amount float64) (float64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	switch opType {
	case Deposit:
		return balance + amount, nil
	case Withdraw:
		if balance < amount {
			return balance, fmt.Errorf(ErrImpossibleOperation)
		}
		return balance - amount, nil
	default:
		return 0, fmt.Errorf(ErrOperationType)
	}
}
