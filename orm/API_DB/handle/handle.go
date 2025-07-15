package handle

import (
	"myModule/orm/API_DB/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetBooksHandler godoc
// @Summary Get all books
// @Description Get all books
// @Tags book
// @Accept json
// @Produce json
// @Success 200 {array} model.BookResponse
// @Failure 400 {string} string "Bad Request"
// @Router /books [get]
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

// GetBookIDHandler godoc
// @Summary Get book by ID
// @Description Get a single book by ID
// @Tags book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} model.BookResponse
// @Failure 400 {string} string "Bad Request"
// @Router /book/{id} [get]
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

// PostBookHandler godoc
// @Summary Create a new book
// @Description Create a new book record
// @Tags book
// @Accept json
// @Produce json
// @Param book body model.Book true "Book info"
// @Success 200 {object} model.BookResponse
// @Failure 400 {string} string "Bad Request"
// @Router /book [post]
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

// PutBookHandler godoc
// @Summary Update book by ID
// @Description Update a book record by ID
// @Tags book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body model.Book true "Book info"
// @Success 200 {object} model.BookResponse
// @Failure 400 {string} string "Bad Request"
// @Router /book/{id} [put]
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

// DeleteBookHandler godoc
// @Summary Delete book by ID
// @Description Delete a book record by ID
// @Tags book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Bad Request"
// @Router /book/{id} [delete]
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


// GetBookByNameHandler godoc
// @Summary Search book by name
// @Description Search books by name
// @Tags book
// @Accept json
// @Produce json
// @Param book body model.Book true "Book name"
// @Success 200 {array} model.BookResponse
// @Failure 400 {string} string "Bad Request"
// @Router /books/search [post]
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
