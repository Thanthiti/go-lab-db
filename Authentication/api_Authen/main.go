package main

import (
	"fmt"
	"log"
	"myModule/Authentication/api_Authen/handle"
	"myModule/Authentication/api_Authen/model"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Check Connected database
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	
	db.AutoMigrate(&model.Book{}, &model.User{})
	
	// Route Book
	app.Use("/books",handle.AuthRequire)
	app.Get("/books", handle.GetBooksHandler(db))
	app.Get("/book/:id", handle.GetBookIDHandler(db))
	app.Post("/book/", handle.PostBookHandler(db))
	app.Put("/book/:id", handle.PutBookHandler(db))
	app.Delete("/book/:id", handle.DeleteBookHandler(db))
	app.Post("/books/search", handle.GetBookByNameHandler(db))

	// Route User
	app.Post("/register/", func(c *fiber.Ctx) error {
		user := new(model.User)
		if err := c.BodyParser(user); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		err = model.CreateUser(db, user)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(fiber.Map{
			"message": "Register success",
		})
	})

	app.Post("/login/", func(c *fiber.Ctx) error {
		user := new(model.User)
		if err := c.BodyParser(user); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		token, err := model.LoginUser(db, user)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		c.Cookie(&fiber.Cookie{
			Name : "jwt",
			Value : token,
			Expires: time.Now().Add(time.Hour * 72),
			HTTPOnly: true,
		})
		return c.JSON(fiber.Map{
			"message":"login Succress",
		})
	})

	app.Listen(":8080")

}
