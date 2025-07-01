package controllers

import (
	"Integrador/src/sensor_temperatura/application/use_case"
	"Integrador/src/sensor_temperatura/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Create_TemperatureSensor_C struct {
	UseCase *use_case.Create_TemperatureSensor
}

func NewCreate_TemperatureSensor_C(useCase *use_case.Create_TemperatureSensor) *Create_TemperatureSensor_C {
	return &Create_TemperatureSensor_C{UseCase: useCase}
}

func (c *Create_TemperatureSensor_C) Execute(ctx *gin.Context) {
	var sensor entities.TemperatureSensor

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
