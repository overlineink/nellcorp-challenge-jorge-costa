package usecase

import (
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database/repositories"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_TransferMoney(t *testing.T) {
	db := database.ConnectDB()
	accountRepository := repositories.AccountRepositoryDb{Db: db}
	transactionRepository := repositories.TransactionRepositoryDb{Db: db}
	depositMoney := TransferMoney{
		AccountRepository:     &accountRepository,
		TransactionRepository: &transactionRepository,
	}
	err := depositMoney.Execute("c3db21a0-ddd8-4eab-81c3-0a78ed51106a", "f1cf2d8d-bdbb-4d6a-b333-753a45f451e3", 1000, "a simple gift")
	require.Nil(t, err)
}
