package handlers

import (
	"database/sql"
	"strconv"

	"github.com/facundocarballo/microservices-todo-list/api/utils"
	"github.com/facundocarballo/microservices-todo-list/domain"
	domain_task "github.com/facundocarballo/microservices-todo-list/domain/task"
	infrastructure "github.com/facundocarballo/microservices-todo-list/infrastructure/task"
	"github.com/gofiber/fiber/v2"
)

// TODO: Test this methods with POSTMAN

func CreateTask(
	c *fiber.Ctx,
	db *sql.DB,
	taskRepository infrastructure.TaskRepository,
) error {
	task := domain_task.BodyToTask(c.Body())
	if task == nil {
		return c.Status(fiber.StatusBadRequest).SendString(utils.TASK_BAD_REQUEST_MSG)
	}

	task, err := taskRepository.Create(task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Cannot create the task, check the body request.")
	}

	c.Set(utils.CONTENT_TYPE, utils.APP_JSON)
	return c.Status(fiber.StatusOK).JSON(task)
}

func CompleteTask(
	c *fiber.Ctx,
	db *sql.DB,
	taskRepository infrastructure.TaskRepository,
) error {
	task := domain_task.BodyToTask(c.Body())
	if task == nil {
		return c.Status(fiber.StatusBadRequest).SendString(utils.TASK_BAD_REQUEST_MSG)
	}

	taskCompleted, err := taskRepository.Complete(task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Cannot complete the task, check the body request.")
	}

	c.Set(utils.CONTENT_TYPE, utils.APP_JSON)
	return c.Status(fiber.StatusOK).JSON(taskCompleted)
}

func DeleteTask(
	c *fiber.Ctx,
	db *sql.DB,
	taskRepository infrastructure.TaskRepository,
) error {
	task := domain_task.BodyToTask(c.Body())
	if task == nil {
		return c.Status(fiber.StatusBadRequest).SendString(utils.TASK_BAD_REQUEST_MSG)
	}

	taskDeleted, err := taskRepository.Delete(task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Cannot delete the task, check the body request.")
	}

	c.Set(utils.CONTENT_TYPE, utils.APP_JSON)
	return c.Status(fiber.StatusOK).JSON(taskDeleted)
}

func GetTasksFromUser(
	c *fiber.Ctx,
	db *sql.DB,
	taskRepository infrastructure.TaskRepository,
) error {
	userIdString := c.Query("user_id", "")
	if userIdString == "" {
		return c.Status(fiber.StatusBadRequest).SendString("'user_id' parameter not found.")
	}
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("'user_id' not a number.")
	}

	tasks, err := taskRepository.Get(&domain.User{Id: userId})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Cannot get all the task for this user.")
	}

	c.Set(utils.CONTENT_TYPE, utils.APP_JSON)
	return c.Status(fiber.StatusOK).JSON(tasks)
}
