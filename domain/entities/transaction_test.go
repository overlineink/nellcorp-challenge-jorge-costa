package entities_test

import (
	"github.com/overlineink/nellcorp-challenge-jorge-costa/domain/entities"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_NewTransaction(t *testing.T) {
	payee, _ := entities.NewAccount("Jane Doe", 10)

	transaction1, err1 := entities.NewTransaction(entities.MoneyDeposit, 10, payee, nil, "to Jane")
	transaction1.Commit()
	require.Nil(t, err1)
	require.Equal(t, entities.TransactionConfirmed, transaction1.Status)
	require.Equal(t, float64(20), payee.Balance)

	account, _ := entities.NewAccount("Jane Doe", 10)
	transaction2, err2 := entities.NewTransaction(entities.MoneyTransfer, 10, payee, account, "to Jane")
	transaction2.Commit()
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
