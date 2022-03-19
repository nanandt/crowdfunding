package handler

import (
	"crowdfunding/helper"
	"crowdfunding/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler{
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context){
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors":errors}
		response := helper.APIResponse("Failed to registered", http.StatusUnprocessableEntity,"error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Failed to registered", http.StatusBadRequest,"error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "ascbakjcbakjcabsc")

	response := helper.APIResponse("Account has been registered", http.StatusOK,"success", formatter)

	c.JSON(http.StatusOK, response)
}
