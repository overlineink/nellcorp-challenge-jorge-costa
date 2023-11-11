package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/infra/http"
)

func main() {
	app := fiber.New()

	app.Post("/account/register", http.RegisterAccountHandler)

	app.Listen(":3000")
}
