package usecase

import (
	"errors"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/application/repositories"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/domain/entities"
)

type WithdrawMoney struct {
	AccountRepository     repositories.AccountRepository
	TransactionRepository repositories.TransactionRepository
}

func (u *WithdrawMoney) Execute(accountId string, amount float64, description string) error {
	account, errFindingAccount := u.AccountRepository.FindAccountById(accountId)
	if errFindingAccount != nil {
		return errFindingAccount
	}

	transaction, errBuildingTransaction := entities.NewTransaction(entities.MoneyWithdraw, amount, nil, account, description)
	if errBuildingTransaction != nil {
		return errBuildingTransaction
	}

	err := u.TransactionRepository.Register(transaction)
	if err != nil {
		return errors.New("unable to register this transaction")
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
