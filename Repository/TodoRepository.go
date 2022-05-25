package Repository

import (
	"todo/backend/Config"
	"todo/backend/Models"
)

func GetAllTodos(todo *[]Models.Todo) (err error) {
	if err = Config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetAllTodosByActivityId(todo *[]Models.Todo, activity_group_id string) (err error) {
	if err = Config.DB.Where("activity_group_id = ?", activity_group_id).Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateTodo(todo *Models.Todo) (err error) {
	if err = Config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetTodoById(todo *Models.Todo, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodo(todo *Models.Todo, id string) (err error) {

	if todo.Title != "" {
		if err = Config.DB.Model(todo).Where("id = ?", id).Update("title", todo.Title).Error; err != nil {
			return err
		}

	} else {
		if err = Config.DB.Model(todo).Where("id = ?", id).Update("is_active", todo.Is_active).Error; err != nil {
			return err
		}
	}

	return nil

}

func DeleteTodo(todo *Models.Todo, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Delete(todo).Error; err != nil {
		return err
	}
	return nil
}
