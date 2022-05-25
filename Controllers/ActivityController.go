package Controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"todo/backend/Models"
	"todo/backend/Repository"
)

func GetActivities(c *fiber.Ctx) error {
	var activity []Models.Activity
	var activity_view []Models.Activity_model
	var resp Models.Response
	Repository.GetAllActivities(&activity)

	for _, s := range activity {
		var act Models.Activity_model
		act.Id = s.Id
		act.Email = s.Email
		act.Title = s.Title
		act.Created_at = s.Created_at.Format("2006-01-02T03:04:05.000Z")
		act.Updated_at = s.Updated_at.Format("2006-01-02T03:04:05.000Z")
		if s.Deleted_at == nil {
			act.Deleted_at = nil
		} else {
			var deleted_time = s.Deleted_at.Format("2006-01-02T03:04:05.000Z")
			act.Deleted_at = &deleted_time
		}
		activity_view = append(activity_view, act)
	}

	resp.Status = "Success"
	resp.Message = "Success"
	resp.Data = activity_view
	
	c.Status(200)
	return c.JSON(resp)

}

func CreateActivity(c *fiber.Ctx) error {
	var resp Models.Response
	var activity_view Models.Activity_Created

	newActivity := new(Models.Activity)

	c.BodyParser(newActivity)

	if newActivity.Title == "" {
		resp.Status = "Bad Request"
		resp.Message = "title cannot be null"
		resp.Data = new(Models.EmptyObject)
		c.Status(400)
		return c.JSON(resp)
	}

	Repository.CreateActivity(newActivity)

	activity_view.Id = newActivity.Id
	activity_view.Email = newActivity.Email
	activity_view.Title = newActivity.Title
	activity_view.Created_at = newActivity.Created_at
	activity_view.Updated_at = newActivity.Updated_at

	resp.Status = "Success"
	resp.Message = "Success"
	resp.Data = activity_view

	c.Status(201)
	return c.JSON(resp)

}

func GetActivityById(c *fiber.Ctx) error {
	id := c.Params("id")
	var activity Models.Activity
	var resp Models.Response
	err := Repository.GetActivityById(&activity, id)
	if err != nil {
		resp.Status = "Not Found"
		resp.Message = fmt.Sprintf("Activity with ID %s Not Found", id)
		resp.Data = new(Models.EmptyObject)
		c.Status(404)
		return c.JSON(resp)
	} else {

		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = activity

		c.Status(200)
		return c.JSON(resp)
	}
}

func UpdateActivity(c *fiber.Ctx) error {

	var act Models.Activity
	var resp Models.Response

	id := c.Params("id")

	newActivity := new(Models.Activity)

	c.BodyParser(newActivity)

	if newActivity.Title == "" {
		resp.Status = "Bad Request"
		resp.Message = "title cannot be null"
		resp.Data = new(Models.EmptyObject)

		c.Status(400)
		return c.JSON(resp)
	}

	err := Repository.GetActivityById(&act, id)
	if err != nil {
		resp.Status = "Not Found"
		resp.Message = fmt.Sprintf("Activity with ID %s Not Found", id)
		resp.Data = new(Models.EmptyObject)
		c.Status(404)
		return c.JSON(resp)
	}

	err = Repository.UpdateActivity(newActivity, id)

	_ = Repository.GetActivityById(&act, id)

	resp.Status = "Success"
	resp.Message = "Success"
	resp.Data = act

	c.Status(200)
	return c.JSON(resp)

}

func DeleteActivity(c *fiber.Ctx) error {
	var activity Models.Activity
	var resp Models.Response
	id := c.Params("id")

	err := Repository.GetActivityById(&activity, id)
	if err != nil {
		resp.Status = "Not Found"
		resp.Message = fmt.Sprintf("Activity with ID %s Not Found", id)
		resp.Data = new(Models.EmptyObject)
		c.Status(404)
		return c.JSON(resp)

	}

	Repository.DeleteActivity(&activity, id)

	resp.Status = "Success"
	resp.Message = "Success"
	resp.Data = new(Models.EmptyObject)
	c.Status(200)
	return c.JSON(resp)

}
