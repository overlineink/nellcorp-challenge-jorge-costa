package usecase_test

import (
	"github.com/overlineink/nellcorp-challenge-jorge-costa/application/usecase"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database/repositories"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_RefundMoney(t *testing.T) {
	db := database.ConnectDB()
	transactionRepository := repositories.TransactionRepositoryDb{Db: db}
	accountRepository := repositories.AccountRepositoryDb{Db: db}
	RefundMoney := usecase.RefundMoney{
		TransactionRepository: &transactionRepository,
	}
	RegisterAccount := usecase.RegisterAccount{AccountRepository: &accountRepository}

	account1, _ := RegisterAccount.Execute("Jorge Costa", 500000)
	account2, _ := RegisterAccount.Execute("Assis Ngolo", 50000)

	depositMoney := usecase.TransferMoney{
		AccountRepository:     &accountRepository,
		TransactionRepository: &transactionRepository,
	}
	transaction, _ := depositMoney.Execute(account1.ID, account2.ID, 500000, "its a gift")

	err := RefundMoney.Execute(transaction.ID, "it was a mistake")
	require.Nil(t, err)
}
