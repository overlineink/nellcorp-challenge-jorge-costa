package usecase

import (
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database/repositories"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_RegisterAccount(t *testing.T) {
	db := database.ConnectDB()
	repo := repositories.AccountRepositoryDb{Db: db}
	registerAccount := RegisterAccount{AccountRepository: &repo}
	account, err := registerAccount.Execute("Joshua Nick", 100)
	require.Nil(t, err)
	require.Equal(t, "Joshua Nick", account.FullName)
}
