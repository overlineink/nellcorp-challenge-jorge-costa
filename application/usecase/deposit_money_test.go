package usecase_test

import (
	"github.com/overlineink/nellcorp-challenge-jorge-costa/application/usecase"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database/repositories"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_DepositMoney(t *testing.T) {
	db := database.ConnectDB()
	accountRepository := repositories.AccountRepositoryDb{Db: db}
	transactionRepository := repositories.TransactionRepositoryDb{Db: db}
	RegisterAccount := usecase.RegisterAccount{AccountRepository: &accountRepository}
	account, _ := RegisterAccount.Execute("Jorge Costa", 100)
	depositMoney := usecase.DepositMoney{
		AccountRepository:     &accountRepository,
		TransactionRepository: &transactionRepository,
	}
	err := depositMoney.Execute(account.ID, 900, "my description")
	require.Nil(t, err)
}
