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
	RefundMoney := usecase.RefundMoney{
		TransactionRepository: &transactionRepository,
	}
	err := RefundMoney.Execute("ea94b58f-637c-44f5-a17b-f6d8231c30d2", "it was a mistake")
	require.Nil(t, err)
}
