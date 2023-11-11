package usecase

import (
	"github.com/overlineink/nellcorp-challenge-jorge-costa/application/repositories"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/domain/entities"
)

type RegisterAccount struct {
	AccountRepository repositories.AccountRepository
}

func (u *RegisterAccount) Execute(fullName string, balance float64) (*entities.Account, error) {
	account, err := entities.NewAccount(fullName, balance)
	if err != nil {
		return nil, err
	}

	_, err = u.AccountRepository.Register(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
