package usecase

import (
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database/repositories"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_DepositMoney(t *testing.T) {
	db := database.ConnectDB()
	accountRepository := repositories.AccountRepositoryDb{Db: db}
	transactionRepository := repositories.TransactionRepositoryDb{Db: db}
	depositMoney := DepositMoney{
		AccountRepository:     &accountRepository,
		TransactionRepository: &transactionRepository,
	}
	err := depositMoney.Execute("c3db21a0-ddd8-4eab-81c3-0a78ed51106a", 900, "to my friend Joshua")
	require.Nil(t, err)
}
