package entities_test

import (
	"github.com/overlineink/nellcorp-challenge-jorge-costa/domain/entities"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_NewAccount(t *testing.T) {
	account, err := entities.NewAccount("Assis Ngolo", 0)

	require.Nil(t, err)
	require.Equal(t, 0.00, account.Balance)
	require.Equal(t, "Assis Ngolo", account.FullName)
	require.NotEmpty(t, account.Number)
}

func Test_FullNameValidation(t *testing.T) {
	_, err := entities.NewAccount("Assis", 0)
	require.NotNil(t, err)

	_, err = entities.NewAccount("Assis Ngolo", 0)
	require.Nil(t, err)
}

func Test_CreditAccount(t *testing.T) {
	account, _ := entities.NewAccount("Assis Ngolo", 0)
	require.Equal(t, float64(0), account.Balance)
	account.Credit(10)
	require.Equal(t, float64(10), account.Balance)
}

func Test_DebitAccount(t *testing.T) {
	account, _ := entities.NewAccount("Assis Ngolo", 150)
	account.Debit(10)
	require.Equal(t, float64(140), account.Balance)
}
