package usecase

import (
	"errors"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/application/repositories"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/domain/entities"
)

type RefundMoney struct {
	TransactionRepository repositories.TransactionRepository
}

func (u *RefundMoney) Execute(transactionId, cancelDescription string) error {
	cancelledTransaction, err := u.TransactionRepository.FindById(transactionId)
	if err != nil {
		return err
	}

	transaction, err := entities.NewTransaction(entities.MoneyRefund, cancelledTransaction.Amount, cancelledTransaction.Account, cancelledTransaction.Payee, cancelDescription)

	if err = u.TransactionRepository.Register(transaction); err != nil {
		return err
	}

	transaction.CancelTransaction = cancelledTransaction
	transaction.CancelTransactionID = cancelledTransaction.ID
	if err = transaction.Commit(); err != nil {
		return err
	}

	if err = u.TransactionRepository.Save(transaction); err != nil {
		return errors.New("error while saving transaction")
	}
	if err = u.TransactionRepository.Save(cancelledTransaction); err != nil {
		return errors.New("error while associating older transaction")
	}

	return nil
}
