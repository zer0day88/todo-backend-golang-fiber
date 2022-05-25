package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"todo/backend/Config"
	"todo/backend/Controllers"
	"todo/backend/Models"
)

var err error

func SetupRoutes(app *fiber.App) {

	act := app.Group("/activity-groups")

	act.Get("", Controllers.GetActivities)
	act.Get(":id", Controllers.GetActivityById)
	act.Post("", Controllers.CreateActivity)
	act.Delete(":id", Controllers.DeleteActivity)
	act.Patch(":id", Controllers.UpdateActivity)

	todo := app.Group("/todo-items")
	todo.Get("", Controllers.GetTodos)
	todo.Get(":id", Controllers.GetTodoById)
	todo.Post("", Controllers.CreateTodo)
	todo.Delete(":id", Controllers.DeleteTodo)
	todo.Patch(":id", Controllers.UpdateTodo)
}
func main() {
	godotenv.Load()

	Config.DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: Config.DbURL(Config.BuildDBConfig()),
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	sqlDB, err := Config.DB.DB()
	sqlDB.SetMaxIdleConns(1000)
	sqlDB.SetMaxOpenConns(1000)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		fmt.Println("Status:", err)
	}
	Config.DB.Set("gorm:table_options", "ENGINE=MyISAM").AutoMigrate(&Models.Activity{}, &Models.Todo{})

	app := fiber.New(fiber.Config{
		ETag:                      true,
		DisableDefaultContentType: true,
		DisableHeaderNormalizing:  true,
		DisableStartupMessage:     false,
		StreamRequestBody:         true,
		DisableDefaultDate:        true,
	})

	app.Use(etag.New())

	app.Use(cache.New(cache.Config{
		KeyGenerator: func(c *fiber.Ctx) string {

			return c.Request().URI().String()
		},
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	SetupRoutes(app)

	app.Listen(":3030")
}
