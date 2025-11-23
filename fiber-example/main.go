package main

import "github.com/gofiber/fiber/v2"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = map[string]User{}

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":8080")
}

func setupRoutes(app *fiber.App) {
	app.Post("/users", createUser)
	app.Get("/users", getAllUsers)
	app.Get("/users/:id", getUser)
	app.Put("/users/:id", updateUser)
	app.Delete("/users/:id", deleteUser)
}

func createUser(c *fiber.Ctx) error {
	u := new(User)
	if err := c.BodyParser(u); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	if _, exists := users[u.ID]; exists {
		return c.Status(409).JSON(fiber.Map{"error": "user already exists"})
	}

	users[u.ID] = *u
	return c.JSON(u)
}

func getAllUsers(c *fiber.Ctx) error {
	result := []User{}
	for _, u := range users {
		result = append(result, u)
	}
	return c.JSON(result)
}

func getUser(c *fiber.Ctx) error {
	id := c.Params("id")
	u, exists := users[id]
	if !exists {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(u)
}

func updateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	_, exists := users[id]
	if !exists {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}

	u := new(User)
	if err := c.BodyParser(u); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	u.ID = id

	users[id] = *u
	return c.JSON(u)
}

func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	_, exists := users[id]
	if !exists {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}

	delete(users, id)
	return c.JSON(fiber.Map{"status": "deleted"})
}