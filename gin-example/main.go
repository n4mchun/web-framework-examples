package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = map[string]User{}

func main() {
	r := gin.Default()

	setupRoutes(r)

	r.Run(":8080")
}

func setupRoutes(r *gin.Engine) {
	r.POST("/users", createUser)
	r.GET("/users", getAllUsers)
	r.GET("/users/:id", getUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)
}

func createUser(c *gin.Context) {
	u := new(User)
	if err := c.ShouldBindJSON(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	if _, exists := users[u.ID]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}

	users[u.ID] = *u
	c.JSON(http.StatusOK, u)
}

func getAllUsers(c *gin.Context) {
	result := []User{}
	for _, u := range users {
		result = append(result, u)
	}
	c.JSON(http.StatusOK, result)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	u, exists := users[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, u)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	_, exists := users[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	u := new(User)
	if err := c.ShouldBindJSON(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	u.ID = id

	users[id] = *u
	c.JSON(http.StatusOK, u)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	_, exists := users[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	delete(users, id)
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
