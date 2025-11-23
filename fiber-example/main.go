package main

import "github.com/gofiber/fiber/v2"

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

var users = map[string]User{}

func main() {
	app := fiber.New()

	// Create
	app.Post("/users", func(c *fiber.Ctx) error {
		u := new(User)
		if err := c.BodyParser(u); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
		}

		users[u.ID] = *u
		return c.JSON(u)
	})

	// Read All
	app.Get("/users", func(c *fiber.Ctx) error {
		result := []User{}
		for _, u := range users {
			result = append(result, u)
		}
		return c.JSON(result)
	})

	// Read User
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		u, ok := users[id]
		if !ok {
			return c.Status(404).JSON(fiber.Map{"error": "not found"})
		}
		return c.JSON(u)
	})

	// Update
	app.Put("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		_, ok := users[id]
		if !ok {
			return c.Status(404).JSON(fiber.Map{"error": "not found"})
		}

		u := new(User)
		if err := c.BodyParser(u); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
		}
		u.ID = id

		users[id] = *u
		return c.JSON(u)
	})

	// Delete
	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		_, ok := users[id]
		if !ok {
			return c.Status(404).JSON(fiber.Map{"error": "not found"})
		}

		delete(users, id)
		return c.JSON(fiber.Map{"status": "deleted"})
	})

	app.Listen(":8080")
}