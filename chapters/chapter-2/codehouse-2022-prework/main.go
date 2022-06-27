package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var nextId int = 0
var todos []Todo

func GetNextId() int {
	value := nextId
	nextId++
	return value
}

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"list": todos})
}

func PostTodo(c *gin.Context) {
	var item Todo
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ID = GetNextId()
	todos = append(todos, item)
	c.String(http.StatusCreated, c.FullPath()+"/"+strconv.Itoa(item.ID))
}

func DeleteTodo(c *gin.Context) {
	idString := c.Param("id")

	if id, err := strconv.Atoi(idString); err == nil {
		for index := range todos {
			if todos[index].ID == id {
				todos = append(todos[:index], todos[index+1:]...)
				c.Writer.WriteHeader(http.StatusNoContent)
				return
			}
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
}

func main() {
	todos = append(todos, Todo{ID: GetNextId(), Value: "CodeHouse", DueDate: "7/31/2022"})

	r := gin.Default()
	r.GET("/api/todos", GetTodos)
	r.POST("/api/todos", PostTodo)
	r.DELETE("/api/todos/:id", DeleteTodo)
	r.Run(":8090")
}
