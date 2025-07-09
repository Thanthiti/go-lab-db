package model

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string	`json:"name"`
	Author      string	`json:"author"`
	Description string	`json:"description"`
	Price       int		`json:"price"`
}

func CreateBook(db *gorm.DB ,book *Book) error{
	result := db.Create(book)
	if result.Error != nil{
		return result.Error
	}
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

func GetBooks(db *gorm.DB) *[]Book{
	var book []Book
	result := db.Find(&book)
	if result.Error != nil{
		log.Fatalf("Error get book %v", result.Error)
	}
	
	fmt.Println("Get book successful")
	
	return  &book
}

func UpdateBook(db *gorm.DB ,book *Book){
	result := db.Save(&book)
	if result.Error != nil{
		log.Fatalf("Error Update book %v", result.Error)
	}
	
	fmt.Println("Update book successful")
}

func DeleteBook(db *gorm.DB ,id int){
	var book Book
	result := db.Delete(&book,id)
	if result.Error != nil{
		log.Fatalf("Error delete book %v", result.Error)
	}

	fmt.Println("delete book successful")
}

func SearchBook(db *gorm.DB, bookName string) *Book{
	var book Book
	result := db.Where("name = ?",bookName).First(&book)
	if result.Error != nil{
		log.Fatalf("Error delete book %v", result.Error)
	}
	
	fmt.Println("Search book successful")
	return  &book
}

