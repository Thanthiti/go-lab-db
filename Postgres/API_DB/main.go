package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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



func getTaskHandleID(c *fiber.Ctx) error {
	taskID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	task, err := GetTaskID(taskID)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)

	}

	return c.JSON(task)
}
func getTasksHandle(c *fiber.Ctx) error {

	task, err := GetAllTask()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)

	}

	return c.JSON(task)
}
func PostTaskHandle(c *fiber.Ctx) error {
	task := new(Task)
	if err := c.BodyParser(task); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err := CreateTask(task)
	if  err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	

	return c.JSON(task)
}
func PutTaskHandle(c *fiber.Ctx) error {

	taskID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	task := new(Task)
	if err := c.BodyParser(task); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	UpdateTask, err := UpdateTask(taskID,task) 
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)

	}
	return c.JSON(UpdateTask)
}
func DeleteTaskHandle(c *fiber.Ctx) error {
	taskID , err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)

	}
	err = DeleteTaskID(taskID)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
		
	}
	
	return c.SendStatus(fiber.StatusNoContent)
}

func JoinTaskUserHandle(c *fiber.Ctx) error {
    join, err := JoinTaskUser()
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(join)
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
	app.Use(checkMiddleware)
	app.Get("/task/:id", getTaskHandleID)
	app.Get("/tasks/", getTasksHandle)
	app.Get("/taske/join", JoinTaskUserHandle)
	app.Post("/task/", PostTaskHandle)
	app.Put("/task/:id", PutTaskHandle)
	app.Delete("/task/:id", DeleteTaskHandle)
	app.Listen(":8080")
}
