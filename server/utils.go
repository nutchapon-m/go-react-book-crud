package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func HandleError(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case AppError:
		fmt.Fprintln(c, e)
		return c.SendStatus(e.Code)
	case error:
		fmt.Fprintln(c, e)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return nil
}
