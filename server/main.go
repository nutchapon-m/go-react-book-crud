package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	initConfig()
	initTimeZone()
	logInit()
	initDB()

	// Migrate Model
	db.AutoMigrate(&Book{})
}

func main() {
	app := fiber.New()

	// Middleware
	app.Use(NewLoggerMiddleWare())
	app.Use(NewCorsMiddleWare())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// CRUD
	// Get Books
	app.Get("/books", func(c *fiber.Ctx) error {
		books := []Book{}
		tx := db.Find(&books)
		if tx.Error != nil {
			Error(tx.Error)
			return HandleError(c, tx.Error)
		}

		resBook := []BookReqRes{}

		for _, v := range books {
			resBook = append(resBook, BookReqRes{
				ID:     v.ID,
				Title:  v.Title,
				Author: v.Author,
			})
		}
		c.SendStatus(200)
		return c.JSON(resBook)
	})
	// Get Book
	app.Get("/book", func(c *fiber.Ctx) error {
		id := c.QueryInt("id")
		books := Book{}
		tx := db.Where("id = ?", id).First(&books)
		if tx.Error != nil {
			Error(tx.Error)
			return HandleError(c, tx.Error)
		}

		resBook := BookReqRes{
			ID:     books.ID,
			Title:  books.Title,
			Author: books.Author,
		}
		c.SendStatus(200)
		return c.JSON(resBook)
	})
	// Create Book
	app.Post("/book", func(c *fiber.Ctx) error {
		reqBook := BookReqRes{}
		if err := c.BodyParser(&reqBook); err != nil {
			Error(err)
			return HandleError(c, err)
		}

		book := Book{
			Title:  reqBook.Title,
			Author: reqBook.Author,
		}

		if err := db.Create(&book).Error; err != nil {
			Error(err)
			return HandleError(c, err)
		}
		return c.SendStatus(201)
	})
	// Update Book
	app.Patch("/book", func(c *fiber.Ctx) error {
		reqBook := BookReqRes{}
		if err := c.BodyParser(&reqBook); err != nil {
			Error(err)
			return HandleError(c, err)
		}

		book := Book{
			Title:  reqBook.Title,
			Author: reqBook.Author,
		}

		if err := db.Model(&Book{}).Where("id = ?", reqBook.ID).Updates(book).Error; err != nil {
			Error(err)
			return HandleError(c, err)
		}
		return c.SendStatus(204)
	})
	// Delete Book
	app.Delete("/book", func(c *fiber.Ctx) error {
		id := c.QueryInt("id")

		if err := db.Delete(&Book{}, id).Error; err != nil {
			Error(err)
			return HandleError(c, err)
		}
		return c.SendStatus(200)
	})

	app.Listen(fmt.Sprintf(":%v", viper.GetString("server.port")))
}
