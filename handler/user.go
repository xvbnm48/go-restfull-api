package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/go-restfull-api/helper"
	"github.com/xvbnm48/go-restfull-api/user"
	"net/http"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}
func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// error handling
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user, err := h.userService.RegisterUser(input)
	if err != nil {
		// error handling
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)
}
