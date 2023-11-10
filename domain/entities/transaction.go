package entities

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/core/domain"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/exp/slices"
	"time"
)

const (
	TransactionPending   string = "pending"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
)

const (
	MoneyTransfer = "transfer"
	MoneyWithdraw = "withdraw"
	MoneyDeposit  = "deposit"
	MoneyRefund   = "refund"
)

type Transaction struct {
	domain.Entity     `valid:"required"`
	Type              string       `json:"type" valid:"notnull"`
	Amount            float64      `json:"amount" valid:"notnull"`
	Status            string       `json:"status" valid:"notnull"`
	Payee             *Account     `json:"payee" valid:"-"`
	Account           *Account     `json:"account" valid:"-"`
	Description       string       `json:"description" valid:"-"`
	CancelTransaction *Transaction `json:"cancel_transaction" valid:"-"`
	CancelDescription string       `json:"cancel_description" valid:"-"`
}

func (t *Transaction) isAccountRequired(operation string) bool {
	allowedOperations := []string{MoneyTransfer, MoneyRefund, MoneyWithdraw}
	return slices.Contains(allowedOperations, operation)
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		return err
	}

	if t.Amount <= 0.00 {
		return errors.New("transaction failed! This is not a valid amount for the transaction")
	}

	if (t.Account == nil) && t.isAccountRequired(t.Type) {
		return errors.New("account not provided")
	}

	if (t.Account != nil) && (!t.isAccountRequired(t.Type)) {
		return errors.New("account not allowed to perform this operation")
	}

	if (t.Type == MoneyWithdraw) && (t.Payee != nil) {
		return errors.New("payee not allowed to perform this operation")
	}

	if (t.Account != nil) && (t.Payee != nil) && (t.Account.ID == t.Payee.ID) && (t.isAccountRequired(t.Type)) {
		return errors.New("the source and destination account cannot be the same")
	}

	if (t.Account != nil) && (t.Type == MoneyTransfer) && (t.Account.Balance < t.Amount) {
		return errors.New("your account does not have enough balance to complete the transaction")
	}

	//if (t.Type == MoneyRefund) && (t.CancelTransaction.Status != TransactionCompleted) {
	//	return errors.New("the given transaction was not processed yet")
	//}

	return nil
}

func (t *Transaction) Commit() error {
	switch t.Type {
	case MoneyTransfer:
		if err := t.Account.Debit(t.Amount); err != nil {
			return err
		}
		if err := t.Payee.Credit(t.Amount); err != nil {
			return err
		}
	case MoneyWithdraw:
		if err := t.Account.Debit(t.Amount); err != nil {
			return err
		}
	case MoneyDeposit:
		if err := t.Payee.Credit(t.Amount); err != nil {
			return nil
		}
	case MoneyRefund:
		if err := t.refund(); err != nil {
			return err
		}
	default:
		return errors.New("unknown operation")
	}
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Transaction) refund() error {
	if err := t.Account.Debit(t.Amount); err != nil {
		return err
	}
	if err := t.Payee.Credit(t.Amount); err != nil {
		return err
	}
	t.CancelTransaction.CancelDescription = t.Description
	t.CancelTransaction.Status = TransactionError
	t.CancelTransaction.UpdatedAt = time.Now()
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()

	if err := t.CancelTransaction.isValid(); err != nil {
		return err
	}

	return nil
}

func NewTransaction(Type string, amount float64, payee *Account, account *Account, description string) (*Transaction, error) {
	transaction := Transaction{
		Type:        Type,
		Amount:      amount,
		Status:      TransactionPending,
		Payee:       payee,
		Account:     account,
		Description: description,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
