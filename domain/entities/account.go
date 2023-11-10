package entities

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/core/domain"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"time"
)

type Account struct {
	domain.Entity `valid:"required"`
	FullName      string         `json:"full_name" valid:"notnull"`
	Number        string         `json:"number" valid:"notnull"`
	Balance       float64        `json:"balance" valid:"notnull"`
	Transactions  []*Transaction `json:"transactions" valid:"-"`
}

func (a *Account) isValid() error {
	if _, err := govalidator.ValidateStruct(a); err != nil {
		return err
	}

	if len(a.FullName) <= 5 {
		return errors.New("Please provide a full name")
	}

	return nil
}

func (a *Account) GenerateAccountNumber() {
	startRange := 10000
	offset := 999999
	a.Number = fmt.Sprint(offset + rand.Intn(offset-startRange+1))
}

func (a *Account) Credit(amount float64) error {
	a.Balance += amount
	a.UpdatedAt = time.Now()

	return nil
}

func (a *Account) Debit(amount float64) error {
	a.Balance -= amount
	a.UpdatedAt = time.Now()
	return nil
}

func (a *Account) AddTransaction(transaction *Transaction) {
	a.Transactions = append(a.Transactions, transaction)
}

func NewAccount(fullName string, balance float64) (*Account, error) {
	account := Account{
		FullName: fullName,
		Balance:  balance,
	}

	account.GenerateAccountNumber()
	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	if err := account.isValid(); err != nil {
		return nil, err
	}

	return &account, nil
}
