package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/smirnoffkin/todo-app/pkg/models"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

type getAllListsResponse struct {
	Data []models.TodoList `json:"data"`
}

type statusResponse struct {
	Status string `json:"status"`
}
