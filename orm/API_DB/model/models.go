package model

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string	`json:"name"`
	Author      string	`json:"author"`
	Description string	`json:"description"`
	Price       int		`json:"price"`
}

type BookResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
}


func ToBookResponse(book *Book) *BookResponse {
	return &BookResponse{
		ID:          book.ID,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
		DeletedAt:   toTimePtr(book.DeletedAt),
		Name:        book.Name,
		Author:      book.Author,
		Description: book.Description,
		Price:       book.Price,
	}
}

func toTimePtr(deletedAt gorm.DeletedAt) *time.Time {
	if deletedAt.Valid {
		return &deletedAt.Time
	}
	return nil
}


func CreateBook(db *gorm.DB ,book *Book) error{
	result := db.Create(book)
	if result.Error != nil{
		return result.Error
	}

	fmt.Println("Create book successful")
	return nil
}

func GetBook(db *gorm.DB ,id int) *Book{
	var book Book
	result := db.First(&book,id)
	if result.Error != nil{
		log.Fatalf("Error get book %v", result.Error)
	}
	
	fmt.Println("Get book successful")
	
	return  &book
}
func GetBooks(db *gorm.DB) []Book{
	var book []Book
	result := db.Find(&book)
	if result.Error != nil{
		log.Fatalf("Error get book %v", result.Error)
	}
	
	fmt.Println("Get book successful")
	
	return  book
}

func UpdateBook(db *gorm.DB ,book *Book) error {
	result := db.Save(&book)
	if result.Error != nil{
		return  result.Error
	}
	
	fmt.Println("Update book successful")
	return nil
}

func DeleteBook(db *gorm.DB ,id int) error{
	var book Book
	result := db.Delete(&book,id)
	if result.Error != nil{
		return 	result.Error
	}
	
	fmt.Println("delete book successful")
	return  nil
}

func SearchBook(db *gorm.DB, bookName string) []Book{
	var book []Book
	result := db.Where("name = ?",bookName).Find(&book)
	if result.Error != nil{
		log.Fatalf("Error get book %v", result.Error)
	}
	
	fmt.Println("Search book successful")
	return  book
}

