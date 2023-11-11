package repositories

import "github.com/overlineink/nellcorp-challenge-jorge-costa/domain/entities"

type AccountRepository interface {
	FindAccountById(id string) (*entities.Account, error)
	Save(account *entities.Account) error
	Register(account *entities.Account) (*entities.Account, error)
}
