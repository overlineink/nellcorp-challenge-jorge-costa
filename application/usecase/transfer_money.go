package usecase

import (
	"github.com/overlineink/nellcorp-challenge-jorge-costa/application/repositories"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/domain/entities"
	"sync"
)

type TransferMoney struct {
	accountRepository     repositories.AccountRepository
	transactionRepository repositories.TransactionRepository
}

func (u *TransferMoney) Execute(
	accountId string,
	payeeId string,
	amount float64,
	description string,
	processTransactionChan chan<- *entities.Transaction,
	wg *sync.WaitGroup,
) error {
	defer wg.Done()
	account, err := u.accountRepository.FindAccountById(accountId)
	if err != nil {
		return err
	}
	payee, err := u.accountRepository.FindAccountById(payeeId)
	if err != nil {
		return err
	}

	transaction, err := entities.NewTransaction(entities.MoneyTransfer, amount, payee, account, description)
	if err != nil {
		return err
	}

	account.AddTransaction(transaction)
	err = u.accountRepository.Save(account)
	if err != nil {
		return err
	}
	payee.AddTransaction(transaction)
	err = u.accountRepository.Save(payee)
	if err != nil {
		return err
	}

	processTransactionChan <- transaction

	return nil
}
