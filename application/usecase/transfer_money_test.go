package usecase_test

import (
	"github.com/overlineink/nellcorp-challenge-jorge-costa/application/usecase"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database/repositories"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_TransferMoney(t *testing.T) {
	db := database.ConnectDB()
	accountRepository := repositories.AccountRepositoryDb{Db: db}
	transactionRepository := repositories.TransactionRepositoryDb{Db: db}

	RegisterAccount := usecase.RegisterAccount{AccountRepository: &accountRepository}
	account1, _ := RegisterAccount.Execute("Assis Ngolo", 500000)
	account2, _ := RegisterAccount.Execute("Jorge Costa", 50000)

	depositMoney := usecase.TransferMoney{
		AccountRepository:     &accountRepository,
		TransactionRepository: &transactionRepository,
	}
	_, err := depositMoney.Execute(account1.ID, account2.ID, 500000, "my description")
	require.Nil(t, err)
}
