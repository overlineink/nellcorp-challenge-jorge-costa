package entities_test

import (
	"github.com/overlineink/nellcorp-challenge-jorge-costa/domain/entities"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_NewAccount(t *testing.T) {
	account, err := entities.NewAccount("Assis Ngolo")

	require.Nil(t, err)
	require.Equal(t, 0.00, account.Balance)
	require.NotEmpty(t, account.Number)
}

func Test_FullNameValidation(t *testing.T) {
	_, err := entities.NewAccount("Assis")
	require.NotNil(t, err)

	_, err = entities.NewAccount("Assis Ngolo")
	require.Nil(t, err)
}
