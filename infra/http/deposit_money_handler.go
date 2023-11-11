package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/application/usecase"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database/repositories"
)

type DepositMoneyDto struct {
	AccountId   string
	Amount      float64
	Description string
}

func DepositMoneyHandler(c *fiber.Ctx) error {
	body := new(DepositMoneyDto)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("malformed body")
	}

	db := database.ConnectDB()
	accountRepository := repositories.AccountRepositoryDb{Db: db}
	transactionRepository := repositories.TransactionRepositoryDb{Db: db}
	DepositMoney := usecase.DepositMoney{
		AccountRepository:     &accountRepository,
		TransactionRepository: &transactionRepository,
	}

	err := DepositMoney.Execute(body.AccountId, body.Amount, body.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())

	}

	c.Status(fiber.StatusOK).SendString("Transaction were successful")
	return nil
}
