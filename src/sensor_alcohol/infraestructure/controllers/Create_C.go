package controllers

import (
	"Integrador/src/sensor_alcohol/application/use_case"
	"Integrador/src/sensor_alcohol/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Create_AlcoholSensor_C struct {
	UseCase *use_case.Create_AlcoholSensor
}

func NewCreate_AlcoholSensor_C(useCase *use_case.Create_AlcoholSensor) *Create_AlcoholSensor_C {
	return &Create_AlcoholSensor_C{UseCase: useCase}
}

func (c *Create_AlcoholSensor_C) Execute(ctx *gin.Context) {
	var sensor entities.AlcoholSensor

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
