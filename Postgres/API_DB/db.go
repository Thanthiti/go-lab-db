package main

import (
	"fmt"
	"log"
)

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

func JoinTaskUser() ([]TaskResponse,error){

	var join []TaskResponse
	row,err := db.Query(`SELECT "ID", "Task",u."Name"
	FROM "Task" as t  join "User" as u on t."User_ID" = u."UserID"
	order by "ID"`)	
	if err != nil{
		log.Fatal(err)
	}
	for row.Next(){
		var res TaskResponse
		if err := row.Scan(&res.ID,&res.Task,&res.UserName); err != nil {
			return nil,err
	}
		join = append(join, res)
	}
	fmt.Println(join[0].ID)
	if err := row.Err(); err != nil {
		return nil, err
	}
	return join,nil
}
