package Models

import (
	"todo/backend/Utils"
)

type Todo struct {
	Id                uint            `json:"id" gorm:"primaryKey; autoIncrement"`
	Activity_group_id int             `json:"activity_group_id" form:"activity_group_id"`
	Title             string          `json:"title"`
	Is_active         bool            `json:"is_active" gorm:"default:1"`
	Priority          string          `json:"priority" gorm:"default:'very-high'"`
	Created_at        Utils.JSONTime  `json:"created_at"`
	Updated_at        Utils.JSONTime  `json:"updated_at"`
	Deleted_at        *Utils.JSONTime `json:"deleted_at"`
}

type Todo_Created struct {
	Created_at        Utils.JSONTime `json:"created_at"`
	Updated_at        Utils.JSONTime `json:"updated_at"`
	Id                uint           `json:"id"`
	Title             string         `json:"title"`
	Activity_group_id int            `json:"activity_group_id"`
	Is_active         bool           `json:"is_active"`
	Priority          string         `json:"priority"`
}
