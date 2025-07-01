package controllers

import (
	"Integrador/src/sensor_turbuidez/application/use_case"
	"Integrador/src/sensor_turbuidez/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Create_TurbiditySensor_C struct {
	UseCase *use_case.Create_TurbiditySensor
}

func NewCreate_TurbiditySensor_C(useCase *use_case.Create_TurbiditySensor) *Create_TurbiditySensor_C {
	return &Create_TurbiditySensor_C{UseCase: useCase}
}

func (c *Create_TurbiditySensor_C) Execute(ctx *gin.Context) {
	var sensor entities.TurbiditySensor

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
