package handle

import (
	"fmt"
	"myModule/orm/API_DB/model"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func AuthRequire(c *fiber.Ctx) error{
	cookie := c.Cookies("jwt")
	jwtSecretKey := os.Getenv("JWT_SECRET")
	token ,err := jwt.ParseWithClaims(cookie , jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey),nil
	})
	if err != nil || !token.Valid{
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claim := token.Claims.(jwt.MapClaims)
	fmt.Println(claim)
	return c.Next()
}

func GetBooksHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		books := model.GetBooks(db)

		var res []model.BookResponse
		for _, book := range books {
			res = append(res, *model.ToBookResponse(&book))
		}
		return c.JSON(res)
	}
}

func GetBookIDHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bookID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		book := model.GetBook(db, bookID)
		return c.JSON(model.ToBookResponse(book))
	}
}


func PostBookHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		newBook := new(model.Book)
		if err := c.BodyParser(newBook); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		err := model.CreateBook(db, newBook)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(model.ToBookResponse(newBook))
	}
}

func PutBookHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bookID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		UpdateBook := new(model.Book)
		if err := c.BodyParser(UpdateBook); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		UpdateBook.ID = uint(bookID)
		err = model.UpdateBook(db, UpdateBook)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(model.ToBookResponse(UpdateBook))
	}
}

func DeleteBookHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bookID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		err = model.DeleteBook(db, bookID)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(fiber.Map{
			"message": "Delete Book Successful",
		})
	}
}


func GetBookByNameHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		NameBook := new(model.Book)
		if err := c.BodyParser(NameBook); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		books := model.SearchBook(db, NameBook.Name)

		var responses []model.BookResponse
		for _, b := range books {
			responses = append(responses, *model.ToBookResponse(&b))
		}

		return c.JSON(responses)
	}
}


