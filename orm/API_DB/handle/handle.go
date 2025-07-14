package handle

import (
	"myModule/orm/API_DB/model"
	"strconv"

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
	return func(c *fiber.Ctx) error {

		bookID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		book := model.GetBook(db, bookID)
		return c.JSON(book)
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
		return c.JSON(fiber.Map{
			"message": "Create Book Successful",
		})
	}
}
func PutBookHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bookID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
			
		}
		UpdateBook := new(model.Book)
		if err := c.BodyParser(UpdateBook); err != nil{
			return c.SendStatus(fiber.StatusBadRequest) 	
			
		}
		UpdateBook.ID = uint(bookID)
		err = model.UpdateBook(db,UpdateBook)
		if err != nil{

			return c.SendStatus(fiber.StatusBadRequest) 	
		}
		return c.JSON(fiber.Map{
			"message": "Update Book Successful",
		})
	}
}
func DeleteBookHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bookID, err := strconv.Atoi(c.Params("id"))
		if err != nil{
		return c.SendStatus(fiber.StatusBadRequest) 	
	}
	err = model.DeleteBook(db,bookID)
	if err != nil{
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
		if err := c.BodyParser(NameBook); err != nil{
			return c.SendStatus(fiber.StatusBadRequest) 	
			
		}
		book := model.SearchBook(db,NameBook.Name) 
		return c.JSON(book)
	}
}
