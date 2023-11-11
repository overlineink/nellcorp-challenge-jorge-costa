package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/domain/entities"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (r *TransactionRepositoryDb) Register(transaction *entities.Transaction) error {
	err := r.Db.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepositoryDb) Save(transaction *entities.Transaction) error {
	err := r.Db.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepositoryDb) FindById(id string) (*entities.Transaction, error) {
	var transaction entities.Transaction
	r.Db.Preload("Account").Preload("Payee").Preload("CancelTransaction").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}
	return &transaction, nil
}
