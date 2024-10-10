package main

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func NewCorsMiddleWare() func(*fiber.Ctx) error {
	return cors.New(cors.Config{
		Next:             nil,
		AllowOriginsFunc: nil,
		AllowOrigins: strings.Join([]string{
			"http://localhost:51160",
			"http://localhost:8000",
			"http://localhost:5173",
		}, ","),
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		ExposeHeaders:    "",
		MaxAge:           0,
	})
}

func NewLoggerMiddleWare() func(*fiber.Ctx) error {
	return logger.New(logger.Config{
		Next:          nil,
		Done:          nil,
		Format:        "[ ${time} ] | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
		TimeFormat:    "2006-01-02 15:04:05",
		TimeZone:      "Local",
		TimeInterval:  500 * time.Millisecond,
		Output:        os.Stdout,
		DisableColors: false,
	})
}

func AuthMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("_token")

	secretKey := viper.GetString("SERVER_PRIVATE_TOKEN")

	_, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return c.SendStatus(401)
	}

	return c.Next()
}
