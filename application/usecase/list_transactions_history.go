package usecase

import (
	"github.com/overlineink/nellcorp-challenge-jorge-costa/application/repositories"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/domain/entities"
)

type ListTransactionsHistory struct {
	accountRepository repositories.AccountRepository
}

func (u *ListTransactionsHistory) Execute(accountId string) ([]*entities.Transaction, error) {
	account, err := u.accountRepository.FindAccountById(accountId)
	if err != nil {
		return nil, err
	}

	return account.Transactions, nil
}
