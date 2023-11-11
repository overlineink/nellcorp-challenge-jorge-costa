package usecase_test

import (
	"github.com/overlineink/nellcorp-challenge-jorge-costa/application/usecase"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database/repositories"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_GetAccountBalance(t *testing.T) {
	db := database.ConnectDB()
	accountRepository := repositories.AccountRepositoryDb{Db: db}
	defaultBalance := 500000.00

	RegisterAccount := usecase.RegisterAccount{AccountRepository: &accountRepository}
	account, _ := RegisterAccount.Execute("Jorge Costa", defaultBalance)
	GetAccountBalance := usecase.GetAccountBalance{
		AccountRepository: &accountRepository,
	}
	balance, err := GetAccountBalance.Execute(account.ID)

	require.Nil(t, err)
	require.Equal(t, defaultBalance, balance)
}
