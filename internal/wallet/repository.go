package wallet

import (
	"fmt"
	"github.com/dacoloris/javacodetest/pkg/db"
	"github.com/google/uuid"
)

type Repository struct {
	Database *db.Db
}

type IRepository interface {
	Get(id uuid.UUID) (*Wallet, error)
	Update(wallet *Wallet) error
}

func NewRepository(database *db.Db) IRepository {
	return &Repository{
		Database: database,
	}
}

func (repo *Repository) Get(id uuid.UUID) (*Wallet, error) {
	var wallet Wallet
	if err := repo.Database.First(&wallet, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (repo *Repository) Update(wallet *Wallet) error {
	fmt.Println("=====UPDATE=====")
	fmt.Println("WALLET ID = ", wallet.ID)
	return repo.Database.DB.Save(wallet).Error
}
