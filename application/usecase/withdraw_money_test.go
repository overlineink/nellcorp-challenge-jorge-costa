package usecase_test

import (
	"github.com/overlineink/nellcorp-challenge-jorge-costa/application/usecase"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database/repositories"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_WithdrawMoney(t *testing.T) {
	db := database.ConnectDB()
	accountRepository := repositories.AccountRepositoryDb{Db: db}
	transactionRepository := repositories.TransactionRepositoryDb{Db: db}

	RegisterAccount := usecase.RegisterAccount{AccountRepository: &accountRepository}
	account, _ := RegisterAccount.Execute("Jorge Costa", 500000)

	WithdrawMoney := usecase.WithdrawMoney{
		AccountRepository:     &accountRepository,
		TransactionRepository: &transactionRepository,
	}
	err := WithdrawMoney.Execute(account.ID, 200000, "beer with friends")
	require.Nil(t, err)
}
