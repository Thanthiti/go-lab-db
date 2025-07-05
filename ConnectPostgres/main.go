package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	database = "TodoList"
	username = "myUser"
	password = "12345678"
)
var db *sql.DB
 
type Task struct{
	id int
	task string
}

func CreateTask(task *Task) error{
	_,err := db.Exec(`INSERT INTO public."Task"("Task") VALUES ($1);`,task.task)
	return err

}
func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)

	fmt.Println(psqlInfo)

	sdb, err := sql.Open("postgres", psqlInfo) // 
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sdb)

	db = sdb

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	if err = CreateTask(&Task{task: "Test Task"}); err != nil{
		log.Fatal(err)
	}
	
}
