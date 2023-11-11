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
	WithdrawMoney := usecase.WithdrawMoney{
		AccountRepository:     &accountRepository,
		TransactionRepository: &transactionRepository,
	}
	err := WithdrawMoney.Execute("c3db21a0-ddd8-4eab-81c3-0a78ed51106a", 20, "beer with friends")
	require.Nil(t, err)
}
