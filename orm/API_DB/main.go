package main

import (
	"fmt"
	"log"
	"myModule/orm/API_DB/handle"
	"myModule/orm/ConnectDB/model"
	"os"

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

	

	
	
	app.Get("/books" ,handle.GetBooksHandler(db)) 
	
	app.Listen(":8080")
	
}
