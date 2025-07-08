package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get task by ID
// @Description Get task detail by ID
// @Tags task
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} Task
// @Failure 400 {string} string "Bad Request"
// @Router /task/{id} [get]
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

// @Summary Get all tasks
// @Description Get all tasks in DB
// @Tags task
// @Accept json
// @Produce json
// @Success 200 {array} Task
// @Failure 400 {string} string "Bad Request"
// @Router /tasks/ [get]
func getTasksHandle(c *fiber.Ctx) error {
    task, err := GetAllTask()
    if err != nil {
        return c.SendStatus(fiber.StatusBadRequest)
    }
    return c.JSON(task)
}

// @Summary Create new task
// @Description Create a new task
// @Tags task
// @Accept json
// @Produce json
// @Param task body Task true "Task data"
// @Success 200 {object} Task
// @Failure 400 {string} string "Bad Request"
// @Router /task/ [post]
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

// @Summary Update Task by ID
// @Description Update Task by ID 
// @Tags task
// @Accept json
// @Produce json
// @Param task body Task true "Task data"
// @Success 200 {array} TaskResponse
// @Failure 400 {string} string "Bad Request"
// @Router /task/{id} [Put]
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

// @Summary Delete Task by ID
// @Description Delete Task by ID 
// @Tags task
// @Accept json
// @Produce json
// @Param task body Task true "Task data"
// @Success 200 {array} TaskResponse
// @Failure 400 {string} string "Bad Request"
// @Router /task/{id} [Delete]
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
// @Summary Get Task join User on id 
// @Description Get task join user in DB 
// @Tags TaskResponse
// @Accept json
// @Produce json
// @Param task body TaskResponse true "TaskResponse data"
// @Success 200 {array} TaskResponse
// @Failure 400 {string} string "Bad Request"
// @Router /task/join [get]
func JoinTaskUserHandle(c *fiber.Ctx) error {
    join, err := JoinTaskUser()
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(join)
}