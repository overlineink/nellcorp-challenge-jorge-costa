package http

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/application/usecase"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database/repositories"
)

func GetAccountBalanceHandler(c *fiber.Ctx) error {
	accountId := c.Params("id")
	if accountId == "" {
		c.Status(fiber.StatusBadRequest)
		return fmt.Errorf("please provide an account ID")
	}

	db := database.ConnectDB()
	accountRepository := repositories.AccountRepositoryDb{Db: db}
	GetAccountBalance := usecase.GetAccountBalance{
		AccountRepository: &accountRepository,
	}

	balanceRaw, err := GetAccountBalance.Execute(accountId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	balance, _ := json.Marshal(balanceRaw)

	c.Status(fiber.StatusOK).Send(balance)
	return nil
}
