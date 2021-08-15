package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang-todo-list/models"
	"golang-todo-list/repositories"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
)

func NewTodoController(db *sql.DB) *TodoController {
	return &TodoController{DB: db}
}

type TodoController struct {
	DB *sql.DB
}

func (tc *TodoController) Index(c *gin.Context) {
	todos, err := repositories.GetAll(tc.DB)

	if err != nil {
		log.Println("err", err)
		c.String(http.StatusInternalServerError, "")

		return
	}

	c.JSON(http.StatusOK, todos)
}

func (tc *TodoController) Create(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
	}

	if err := validator.Validate(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	_, err := repositories.Create(tc.DB, todo.Name)

	if err != nil {
		log.Println("err", err)
		c.String(http.StatusInternalServerError, "")

		return
	}

	c.String(http.StatusCreated, "")
}

func (tc *TodoController) Delete(c *gin.Context) {
	_, err := repositories.Delete(tc.DB, c.Param("id"))

	if err != nil {
		log.Println("err", err)
		c.String(http.StatusInternalServerError, "")

		return
	}

	c.String(http.StatusOK, "")
}
