package wallet

import (
	"github.com/google/uuid"
)

type Wallet struct {
	ID      uuid.UUID `gorm:"type:uuid;primarykey"`
	Balance float64
}
