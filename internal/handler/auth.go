package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smirnoffkin/todo-app/pkg/models"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "User with this username or email is already exists")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	token, err := h.services.Authorization.CreateAccessToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Incorrect email or password")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token": token,
		"token_type":   "bearer",
	})
}
