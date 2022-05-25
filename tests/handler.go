package tests

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"ipanda.baac.tech/golib/zlogwrap"
)

func HandlerContext(c *fiber.Ctx) error {
	logger := zlogwrap.New(zlogwrap.Config{
		ServiceName: "HandlerContext",
		Context:     c,
	})
	body := fiber.Map{}
	json.Unmarshal(c.Request().Body(), &body)
	logger.SetField("body", body).Info("logging")
	return c.SendStatus(fiber.StatusNoContent)
}
