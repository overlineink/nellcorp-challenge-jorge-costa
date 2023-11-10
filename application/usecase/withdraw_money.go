package usecase

import (
	"errors"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/application/repositories"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/domain/entities"
	"sync"
)

type WithdrawMoney struct {
	accountRepository     repositories.AccountRepository
	transactionRepository repositories.TransactionRepository
}

func (u *WithdrawMoney) Execute(accountId string, amount float64, description string, processTransactionChan chan<- *entities.Transaction, wg *sync.WaitGroup) error {
	defer wg.Done()

	account, errFindingAccount := u.accountRepository.FindAccountById(accountId)
	if errFindingAccount != nil {
		return errFindingAccount
	}

	transaction, errBuildingTransaction := entities.NewTransaction(entities.MoneyWithdraw, amount, nil, account, description)
	if errBuildingTransaction != nil {
		return errBuildingTransaction
	}

	err := u.transactionRepository.Register(transaction)
	if err != nil {
		return errors.New("unable to register this transaction")
	}

	account.AddTransaction(transaction)
	err = u.accountRepository.Save(account)
	if err != nil {
		return errors.New("unable to register transaction")
	}

	processTransactionChan <- transaction

	return nil
}
