package repositories

import "github.com/overlineink/nellcorp-challenge-jorge-costa/domain/entities"

type TransactionRepository interface {
	Register(transaction *entities.Transaction) error
	Save(transaction *entities.Transaction) error
	FindById(id string) (*entities.Transaction, error)
}
