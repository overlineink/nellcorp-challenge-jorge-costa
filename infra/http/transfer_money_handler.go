package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/application/usecase"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database/repositories"
)

type TransferMoneyDto struct {
	AccountId   string  `json:"account_id"`
	PayeeId     string  `json:"payee_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

func TransferMoneyHandler(c *fiber.Ctx) error {
	body := new(TransferMoneyDto)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("malformed body")
	}

	db := database.ConnectDB()
	accountRepository := repositories.AccountRepositoryDb{Db: db}
	transactionRepository := repositories.TransactionRepositoryDb{Db: db}
	TransferMoney := usecase.TransferMoney{
		AccountRepository:     &accountRepository,
		TransactionRepository: &transactionRepository,
	}

	_, err := TransferMoney.Execute(body.AccountId, body.PayeeId, body.Amount, body.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())

	}

	c.Status(fiber.StatusOK).SendString("Transaction were successful")
	return nil
}
