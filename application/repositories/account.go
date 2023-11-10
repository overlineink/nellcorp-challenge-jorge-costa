package repositories

import "github.com/overlineink/nellcorp-challenge-jorge-costa/domain/entities"

type AccountRepository interface {
	FindAccountById(id string) (*entities.Account, error)
	RegisterTransaction(transactionId string) error
}
