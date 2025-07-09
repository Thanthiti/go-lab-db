package handle

import (
	"myModule/orm/API_DB/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetBooksHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		books := model.GetBooks(db)
		return c.JSON(books)
	}
}

func GetBookIDHandler(db *gorm.DB) fiber.Handler {
	return func (c *fiber.Ctx)  error{
		return c.JSON("asd")
	}
}
func PostBookHandler(db *gorm.DB) fiber.Handler {
	return func (c *fiber.Ctx)  error{
		return c.JSON("asd")
	}
}
func PutBookHandler(db *gorm.DB) fiber.Handler {
	return func (c *fiber.Ctx)  error{
		return c.JSON("asd")
	}
}
func DeleteBookHandler(db *gorm.DB) fiber.Handler {
	return func (c *fiber.Ctx)  error{
		return c.JSON("asd")
	}
}
func GetBookByNameHandler(db *gorm.DB) fiber.Handler {
	return func (c *fiber.Ctx)  error{
		return c.JSON("asd")
	}
}