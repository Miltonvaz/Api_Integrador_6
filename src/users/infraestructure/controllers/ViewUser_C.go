package controllers

import (
	"Integrador/src/users/application/use_case"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ViewUserController struct {
	useCase use_case.ViewUser
}

func NewViewUserController(useCase use_case.ViewUser) *ViewUserController {
	return &ViewUserController{useCase: useCase}
}

func (cc_c *ViewUserController) Execute(c *gin.Context) {

	users, err := cc_c.useCase.Execute()
	if err != nil {
		fmt.Printf("Error retrieving users: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Unable to retrieve users: %v", err)})
		return
	}

	for i := range users {
		users[i].Password = ""
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
