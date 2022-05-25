package Controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"todo/backend/Models"
	"todo/backend/Repository"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []Models.Todo
	var resp Models.Response
	activity_group_id := c.Query("activity_group_id")
	if activity_group_id == "" {

		Repository.GetAllTodos(&todos)

		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = todos

		return c.Status(200).JSON(&fiber.Map{
			"status":  "Success",
			"message": "Success",
			"data":    todos,
		})

	} else {
		c.Response().Header.Add("refresh", "true")
		c.Response().Header.Del("Content-Length")
		Repository.GetAllTodosByActivityId(&todos, activity_group_id)

		if todos == nil {

			return c.Status(200).JSON(&fiber.Map{
				"status":  "Success",
				"message": "Success",
				"data":    make([]string, 0),
			})
		} else {
			return c.Status(200).JSON(&fiber.Map{
				"status":  "Success",
				"message": "Success",
				"data":    todos,
			})
		}
	}
}

func CreateTodo(c *fiber.Ctx) error {
	var resp Models.Response
	var todo_created Models.Todo_Created
	newTodo := new(Models.Todo)

	c.BodyParser(newTodo)

	if newTodo.Title == "" {
		resp.Status = "Bad Request"
		resp.Message = "title cannot be null"
		resp.Data = new(Models.EmptyObject)
		c.Status(400)
		return c.JSON(resp)
	}

	if newTodo.Activity_group_id == 0 {
		resp.Status = "Bad Request"
		resp.Message = "activity_group_id cannot be null"
		resp.Data = new(Models.EmptyObject)
		c.Status(400)
		return c.JSON(resp)
	}

	Repository.CreateTodo(newTodo)

	todo_created.Id = newTodo.Id
	todo_created.Activity_group_id = newTodo.Activity_group_id
	todo_created.Title = newTodo.Title
	todo_created.Priority = newTodo.Priority
	todo_created.Is_active = newTodo.Is_active
	todo_created.Created_at = newTodo.Created_at
	todo_created.Updated_at = newTodo.Updated_at

	return c.Status(201).JSON(&fiber.Map{
		"status":  "Success",
		"message": "Success",
		"data":    todo_created,
	})

}

func GetTodoById(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo Models.Todo
	var resp Models.Response
	err := Repository.GetTodoById(&todo, id)
	if err != nil {
		resp.Status = "Not Found"
		resp.Message = fmt.Sprintf("Todo with ID %s Not Found", id)
		resp.Data = new(Models.EmptyObject)
		c.Status(404)
		return c.JSON(resp)
	} else {

		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = todo
		c.Status(200)
		return c.JSON(resp)
	}
}

func UpdateTodo(c *fiber.Ctx) error {
	var todo_view Models.Todo
	var resp Models.Response

	id := c.Params("id")

	updateTodo := new(Models.Todo)

	c.BodyParser(updateTodo)

	err := Repository.GetTodoById(&todo_view, id)
	if err != nil {
		resp.Status = "Not Found"
		resp.Message = fmt.Sprintf("Todo with ID %s Not Found", id)
		resp.Data = new(Models.EmptyObject)
		c.Status(404)
		return c.JSON(resp)
	}

	Repository.UpdateTodo(updateTodo, id)

	_ = Repository.GetTodoById(&todo_view, id)

	resp.Status = "Success"
	resp.Message = "Success"
	resp.Data = todo_view
	c.Status(200)
	return c.JSON(resp)

}

func DeleteTodo(c *fiber.Ctx) error {
	var todo Models.Todo
	var resp Models.Response
	id := c.Params("id")

	err := Repository.GetTodoById(&todo, id)
	if err != nil {
		resp.Status = "Not Found"
		resp.Message = fmt.Sprintf("Todo with ID %s Not Found", id)
		resp.Data = new(Models.EmptyObject)
		c.Status(404)
		return c.JSON(resp)
	}

	err = Repository.DeleteTodo(&todo, id)
	if err != nil {
		resp.Status = "Not Found"
		resp.Message = fmt.Sprintf("Todo with ID %s Not Found", id)
		resp.Data = new(Models.EmptyObject)
		c.Status(404)
		return c.JSON(resp)
	} else {
		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = new(Models.EmptyObject)
		c.Status(200)
		return c.JSON(resp)
	}
}
