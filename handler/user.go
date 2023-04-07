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
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("account failed to create", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	NewUser, err := h.userService.RegisterUser(input)
	if err != nil {
		// error handling
		response := helper.APIResponse("account failed to create", 400, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(NewUser, "token")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {}
