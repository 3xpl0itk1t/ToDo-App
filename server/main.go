package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	todos := []Todo{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}
		todo.ID = len(todos) + 1

		todos = append(todos, *todo)

		return c.JSON(todos)

	})

	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(401).SendString(err.Error())
		}
		for i, t := range todos {
			if t.ID == id {
				todos[i].Done = true
				return c.JSON(todos)
			}
		}
		return c.JSON(todos)

	})

	app.Get("/api/todos/", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	port := os.Getenv("PORT")
	app.Listen(":" + port)

}
