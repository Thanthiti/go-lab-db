package main

import (
	"fmt"
	"log"
	"myModule/orm/API_DB/handle"
	"myModule/orm/ConnectDB/model"
	"os"
	
	_ "myModule/orm/API_DB/docs"
	fiberSwagger "github.com/gofiber/swagger"

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

	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{	})

	// Check Connected database
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	db.AutoMigrate(&model.Book{})

	app.Get("/swagger/*", fiberSwagger.HandlerDefault)
	
	app.Get("/books" ,handle.GetBooksHandler(db)) 
	app.Get("/book/:id" ,handle.GetBookIDHandler(db)) 
	app.Post("/book" ,handle.PostBookHandler(db)) 
	app.Put("/book/:id" ,handle.PutBookHandler(db)) 
	app.Delete("/book/:id" ,handle.DeleteBookHandler(db)) 
	app.Post("/books/search" ,handle.GetBookByNameHandler(db)) 
	
	app.Listen(":8080")
	
}
