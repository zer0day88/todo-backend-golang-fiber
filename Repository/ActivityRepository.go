package Repository

import (
	_ "github.com/go-sql-driver/mysql"
	"todo/backend/Config"
	"todo/backend/Models"
)

func GetAllActivities(activity *[]Models.Activity) (err error) {
	if err = Config.DB.Find(activity).Error; err != nil {
		return err
	}
	return nil
}

func CreateActivity(activity *Models.Activity) (err error) {
	if err = Config.DB.Create(activity).Error; err != nil {
		return err
	}
	return nil
}

func GetActivityById(activity *Models.Activity, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(activity).Error; err != nil {
		return err
	}
	return nil
}

func UpdateActivity(activity *Models.Activity, id string) (err error) {
	if err = Config.DB.Model(activity).Where("id = ?", id).Update("title", activity.Title).Error; err != nil {
		return err
	}
	return nil
}

func DeleteActivity(activity *Models.Activity, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Delete(activity).Error; err != nil {
		return err
	}
	return nil
}
