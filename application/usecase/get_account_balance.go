package usecase

import "github.com/overlineink/nellcorp-challenge-jorge-costa/application/repositories"

type GetAccountBalance struct {
	accountRepository repositories.AccountRepository
}

func (u *GetAccountBalance) Execute(accountId string) (float64, error) {
	account, err := u.accountRepository.FindAccountById(accountId)
	if err != nil {
		return 0, err
	}
	return account.Balance, nil
}
