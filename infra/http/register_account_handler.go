package http

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/application/usecase"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/database/repositories"
)

type RegisterAccountDto struct {
	FullName string  `json:"full_name"`
	Amount   float64 `json:"amount"`
}

func RegisterAccountHandler(c *fiber.Ctx) error {
	db := database.ConnectDB()
	accountRepository := repositories.AccountRepositoryDb{Db: db}
	RegisterAccount := usecase.RegisterAccount{AccountRepository: &accountRepository}

	body := new(RegisterAccountDto)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("malformed body")
	}

	accountRaw, err := RegisterAccount.Execute(body.FullName, body.Amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())

	}
	account, _ := json.Marshal(accountRaw)

	c.Status(fiber.StatusCreated).Write(account)

	return nil
}
