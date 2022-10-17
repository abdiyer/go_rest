package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json: "id"`
	Item      string `json: "item"`
	Completed bool   `json: "completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean room", Completed: false},
	{ID: "2", Item: "Make dishes", Completed: false},
	{ID: "3", Item: "Play games", Completed: true},
}

func getShits(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addShit(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)

	print(todos)

}

func getShitById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("shit not found")
}

func getShit(context *gin.Context) {
	id := context.Param("id")
	todo, err := getShitById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "shit not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func toggleShitStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getShitById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "shit not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/address", getShit)
	router.GET("/address/:id", getShit)
	router.PATCH("/todos/:id", toggleShitStatus)
	router.POST("/address", addShit)
	router.Run("localhost:7777")
}
