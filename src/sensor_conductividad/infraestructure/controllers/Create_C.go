package controllers

import (
	"Integrador/src/sensor_conductividad/application/use_case"
	"Integrador/src/sensor_conductividad/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Create_ConductividtySensor_C struct {
	UseCase *use_case.Create_ConductividtySensor
}

func NewCreate_ConductividtySensor_C(useCase *use_case.Create_ConductividtySensor) *Create_ConductividtySensor_C {
	return &Create_ConductividtySensor_C{UseCase: useCase}
}

func (c *Create_ConductividtySensor_C) Execute(ctx *gin.Context) {
	var sensor entities.ConductivitySensor

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
