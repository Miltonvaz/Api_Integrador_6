package controllers

import (
	"Integrador/src/fermentation/application/use_case"
	"Integrador/src/fermentation/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Create_Fermentation_C struct {
	UseCase *use_case.Create_Fermentation
}

func NewCreate_Fermentation_C(useCase *use_case.Create_Fermentation) *Create_Fermentation_C {
	return &Create_Fermentation_C{UseCase: useCase}
}

func (c *Create_Fermentation_C) Execute(ctx *gin.Context) {
	var sensor entities.Fermentation

	if err := ctx.ShouldBindJSON(&sensor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	createdSensor, err := c.UseCase.Execute(sensor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving data"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Data saved successfully", "data": createdSensor})
}
