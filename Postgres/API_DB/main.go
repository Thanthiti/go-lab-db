package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Task struct {
	ID   int
	Task string
}

func CreateTask(task *Task) error {
	_, err := db.Exec(`INSERT INTO public."Task"("Task") VALUES ($1);`, task.Task)

	return err

}

func GetTaskID(id int) (Task, error) {
	var t Task
	row := db.QueryRow(`SELECT "ID","Task" FROM public."Task" WHERE "ID"=$1`, id)

	if err := row.Scan(&t.ID, &t.Task); err != nil {
		return Task{}, err
	}
	return t, nil
}

func GetAllTask() ([]Task, error) {
	var tasks []Task
	row, err := db.Query(`SELECT "ID","Task" FROM public."Task"`)

	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	for row.Next() {
		var task Task
		if err := row.Scan(&task.ID, &task.Task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := row.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func UpdateTask(id int, t *Task) (Task, error) {
	var task Task
	row := db.QueryRow(`UPDATE public."Task" SET  "Task"=$1 WHERE "ID" = $2 RETURNING "ID","Task"`, t.Task, id)
	if err := row.Scan(&task.ID, &task.Task); err != nil {
		return Task{}, err
	}
	return task, nil
}

func DeleteTaskID(id int) error {
	res, err := db.Exec(`DELETE FROM "Task" WHERE  "ID"=$1`, id)
	if err != nil {
		return err
	}
	// check if the id not exits in db, if so delete it
	rowAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	// if not then return message
	if rowAffected == 0 {
		return fmt.Errorf("task with id %d not found", id)
	}
	return err
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

	// if err = CreateTask(&Task{Task: "Test Task"}); err != nil{
	// 	log.Fatal(err)
	// }

	// task,err := GetAllTask()
	// if  err != nil{
	// 	log.Fatal(err)
	// }
	// fmt.Print(task)

	// task,err := GetTaskID(1)
	// if  err != nil{
	// 	log.Fatal(err)
	// }
	// fmt.Print(task)

	// update,err := UpdateTask(4,&Task{Task: "Update Task Test"})
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println(update)

	// err = DeleteTaskID(5)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("Delete Successful")
	


	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})
	app.Listen(":8080")
}
