package main

import (
	"fmt"
	"os"

	"github.com/buildingwatsize/zlogwrap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/zerolog"
)

func main() {

	// Default
	logger1 := zlogwrap.New()
	logger1.Debug()
	logger1.Debug("Debug Log")
	logger1.Info("Info Log")
	logger1.Error("Error Log")

	// a possibilities use please go to func `seeMoreUse()` in `more_usage.go`
	seeMoreUsage()

	// 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀
	// 🚀 USAGE IN FIBER APPLICATION 🚀
	// 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀

	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(requestid.New(requestid.Config{
		// zlogwrap using custom header 'Transaction-Id' instead of 'X-Request-Id'
		// and TransactionIDTag also available to use (it's constant)
		Header: zlogwrap.TransactionIDTag,
	}))

	app.Get("/", Handler) // GET http://localhost:8000/

	fmt.Println("Listening on http://localhost:8000")
	fmt.Println("Try to send a request :D")
	app.Listen(":8000")
}

func Handler(c *fiber.Ctx) error {

	// With Custom Logger and Context
	serviceName := "[Custom Logger]"
	myLogger := zerolog.New(os.Stdout).With().
		Str("foo", "bar").
		Float64("money", 10.99).
		Logger()
	logger4 := zlogwrap.New(zlogwrap.Config{
		ServiceName: serviceName,
		Logger:      myLogger,
		Context:     c,
	})
	logger4.Debug()
	logger4.Debug("Debug Log")
	logger4.Info("Info Log")
	logger4.Error("Error Log")

	return c.SendString("Watch your app logs!")
}
