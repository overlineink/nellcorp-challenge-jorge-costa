package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/http"
)

func main() {
	app := fiber.New()

	app.Post("/account/register", http.RegisterAccountHandler)
	app.Get("/account/:id/balance", http.GetAccountBalanceHandler)
	app.Post("/account/balance", http.DepositMoneyHandler)

	app.Listen(":3000")
}
