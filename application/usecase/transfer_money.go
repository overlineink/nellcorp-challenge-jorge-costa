package usecase

import (
	"errors"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/application/repositories"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/domain/entities"
)

type TransferMoney struct {
	AccountRepository     repositories.AccountRepository
	TransactionRepository repositories.TransactionRepository
}

func (u *TransferMoney) Execute(
	accountId string,
	payeeId string,
	amount float64,
	description string,
) error {
	account, err := u.AccountRepository.FindAccountById(accountId)
	if err != nil {
		return err
	}
	payee, err := u.AccountRepository.FindAccountById(payeeId)
	if err != nil {
		return err
	}

	transaction, err := entities.NewTransaction(entities.MoneyTransfer, amount, payee, account, description)
	if err != nil {
		return err
	}

	err = transaction.Commit()
	if err != nil {
		return err
	}

	err = u.TransactionRepository.Save(transaction)
	if err != nil {
		return errors.New("error while saving transaction")
	}

	return nil
}
