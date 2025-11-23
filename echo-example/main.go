package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = map[string]User{}

func main() {
	e := echo.New()

	setupRoutes(e)

	e.Start(":8080")
}

func setupRoutes(e *echo.Echo) {
	e.POST("/users", createUser)
	e.GET("/users", getAllUsers)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
}

func createUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body"})
	}

	if _, exists := users[u.ID]; exists {
		return c.JSON(http.StatusConflict, map[string]string{"error": "user already exists"})
	}

	users[u.ID] = *u
	return c.JSON(http.StatusOK, u)
}

func getAllUsers(c echo.Context) error {
	result := []User{}
	for _, u := range users {
		result = append(result, u)
	}
	return c.JSON(http.StatusOK, result)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	u, exists := users[id]
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
	}
	return c.JSON(http.StatusOK, u)
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	_, exists := users[id]
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
	}

	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body"})
	}
	u.ID = id

	users[id] = *u
	return c.JSON(http.StatusOK, u)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	_, exists := users[id]
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
	}

	delete(users, id)
	return c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
