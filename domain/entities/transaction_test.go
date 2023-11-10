package entities_test

import (
	"github.com/overlineink/nellcorp-challenge-jorge-costa/domain/entities"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_NewTransaction(t *testing.T) {
	payee, _ := entities.NewAccount("Jane Doe", 10)

	transaction1, err1 := entities.NewTransaction(entities.MoneyDeposit, 10, payee, nil, "to Jane")
	require.Nil(t, err1)
	err1 = transaction1.Commit()
	require.Nil(t, err1)
	require.Equal(t, entities.TransactionCompleted, transaction1.Status)
	require.Equal(t, float64(20), payee.Balance)

	account, _ := entities.NewAccount("Jane Doe", 10)
	transaction2, err2 := entities.NewTransaction(entities.MoneyTransfer, 10, payee, account, "to Jane")
	require.Nil(t, err2)
	err2 = transaction2.Commit()
	require.Nil(t, err2)
	require.Equal(t, float64(0), account.Balance)
}

func Test_NewTransactionWithoutAccount(t *testing.T) {
	payee, _ := entities.NewAccount("Jane Doe", 10.00)
	account, _ := entities.NewAccount("John Doe", 10.00)

	_, err := entities.NewTransaction(entities.MoneyTransfer, 10, payee, nil, "school driving tax")
	require.NotNil(t, err)

	_, err = entities.NewTransaction(entities.MoneyTransfer, 10, payee, payee, "school driving tax")
	require.NotNil(t, err)

	_, err = entities.NewTransaction(entities.MoneyTransfer, 100, payee, account, "school driving tax")
	require.NotNil(t, err)
}

func Test_NewWithdrawTransaction(t *testing.T) {
	account, _ := entities.NewAccount("John Doe", 10.00)
	_, err := entities.NewTransaction(entities.MoneyWithdraw, 10, nil, account, "withdrawal")
	require.Nil(t, err)
}

func Test_MoneyRefunding(t *testing.T) {
	payee, _ := entities.NewAccount("Jane Doe", 10.00)
	account, _ := entities.NewAccount("John Doe", 10.00)
	moneyTransfer, err := entities.NewTransaction(entities.MoneyTransfer, 10, payee, account, "transfer")
	require.Nil(t, err)
	err = moneyTransfer.Commit()
	require.Nil(t, err)
	require.Equal(t, float64(0), account.Balance)
	require.Equal(t, float64(20), payee.Balance)
	moneyRefunding, errRefunding := entities.NewTransaction(
		entities.MoneyRefund,
		moneyTransfer.Amount,
		moneyTransfer.Account,
		moneyTransfer.Payee,
		"cancel",
	)
	require.Nil(t, errRefunding)
	moneyRefunding.CancelTransaction = moneyTransfer
	err = moneyRefunding.Commit()
	require.Nil(t, err)
	require.Equal(t, float64(10), payee.Balance)
	require.Equal(t, float64(10), account.Balance)
	require.Equal(t, entities.TransactionError, moneyRefunding.CancelTransaction.Status)
}
