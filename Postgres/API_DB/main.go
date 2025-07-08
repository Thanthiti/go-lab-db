package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "myModule/Postgres/API_DB/docs"

	fiberSwagger "github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB
// DB model
type Task struct {
	ID   int
	Task string
	UserID string
}

type TaskResponse struct {
    ID       int    `json:"id"`
    Task     string `json:"task"`
    UserName string `json:"user_name"`
}


func checkMiddleware(c *fiber.Ctx) error{
	start := time.Now()

	fmt.Printf("URL = %s , Method = %s , Time = %s\n",c.OriginalURL(),c.Method(),start)

	return  c.Next()

}


func main() {
	defer db.Close()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	sdb, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	db = sdb

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	// Swagger
	app.Get("/swagger/*", fiberSwagger.HandlerDefault)

	app.Use(checkMiddleware)
	app.Get("/task/join", JoinTaskUserHandle)
	app.Get("/tasks", getTasksHandle)
	app.Get("/task/:id", getTaskHandleID)
	app.Post("/task/", PostTaskHandle)
	app.Put("/task/:id", PutTaskHandle)
	app.Delete("/task/:id", DeleteTaskHandle)
	app.Listen(":8080")
}
