package routes

import (
	"github.com/gin-gonic/gin"
	"golang-todo-list/controllers"
)

func SetupRoutes(r *gin.Engine, c *controllers.TodoController) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("/api/v1/todos", c.Index)
	r.POST("/api/v1/todos", c.Create)
	r.DELETE("/api/v1/todos/:id", c.Delete)
}
