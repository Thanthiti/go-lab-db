package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Task struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}


// === Handlers ===

func createTaskHandler(c *fiber.Ctx) error {
	var t Task
	if err := c.BodyParser(&t); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse json"})
	}

	res, err := db.Exec(`INSERT INTO tasks(task) VALUES (?)`, t.Task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot insert"})
	}

	id, _ := res.LastInsertId()
	t.ID = int(id)
	return c.Status(fiber.StatusCreated).JSON(t)
}

func getAllTasksHandler(c *fiber.Ctx) error {
	rows, err := db.Query(`SELECT id, task FROM tasks`)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot query"})
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		rows.Scan(&t.ID, &t.Task)
		tasks = append(tasks, t)
	}

	return c.JSON(tasks)
}

func getTaskByIDHandler(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var t Task
	err := db.QueryRow(`SELECT id, task FROM tasks WHERE id = ?`, id).Scan(&t.ID, &t.Task)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found"})
	}
	return c.JSON(t)
}

func updateTaskHandler(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var t Task
	if err := c.BodyParser(&t); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse json"})
	}

	res, err := db.Exec(`UPDATE tasks SET task = ? WHERE id = ?`, t.Task, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot update"})
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found"})
	}

	t.ID = id
	return c.JSON(t)
}

func deleteTaskHandler(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	res, err := db.Exec(`DELETE FROM tasks WHERE id = ?`, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot delete"})
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found"})
	}

	return c.JSON(fiber.Map{"message": fmt.Sprintf("task %d deleted", id)})
}

func main() {
	// Connect SQLite
	var err error
	db, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table if not exists
	createTable := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT NOT NULL
	);
	`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	// Routes
	app.Post("/tasks", createTaskHandler)
	app.Get("/tasks", getAllTasksHandler)
	app.Get("/tasks/:id", getTaskByIDHandler)
	app.Put("/tasks/:id", updateTaskHandler)
	app.Delete("/tasks/:id", deleteTaskHandler)

	log.Fatal(app.Listen(":8080"))
}
