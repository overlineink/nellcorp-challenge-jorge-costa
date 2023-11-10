package usecase

import (
	"github.com/overlineink/nellcorp-challenge-jorge-costa/application/repositories"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/domain/entities"
	"sync"
)

type RefundMoney struct {
	transactionRepository repositories.TransactionRepository
}

func (u *RefundMoney) Execute(transactionId, cancelDescription string, processTransactionChan chan<- *entities.Transaction, wg *sync.WaitGroup) error {
	defer wg.Done()
	prevTransaction, err := u.transactionRepository.FindById(transactionId)
	if err != nil {
		return err
	}

	transaction, err := entities.NewTransaction(
		entities.MoneyRefund,
		prevTransaction.Amount,
		prevTransaction.Account,
		prevTransaction.Payee,
		cancelDescription,
	)

	processTransactionChan <- transaction

	return nil
}
