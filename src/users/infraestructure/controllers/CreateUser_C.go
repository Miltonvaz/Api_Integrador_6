package controllers

import (
	"Integrador/src/users/application/use_case"
	"Integrador/src/users/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserController struct {
	useCase use_case.CreateUser
}

func NewCreateUserController(useCase use_case.CreateUser) *CreateUserController {
	return &CreateUserController{useCase: useCase}
}

func (cc_c *CreateUserController) Execute(c *gin.Context) {
	var input entities.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := cc_c.useCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	input.Password = ""

	c.JSON(http.StatusCreated, gin.H{"Client": input})
}
