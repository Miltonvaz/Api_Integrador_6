package controllers

import (
	"Integrador/src/sensor_ph/application/use_case"
	entities "Integrador/src/sensor_ph/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Create_PhSensor_C struct {
	UseCase *use_case.Create_PhSensor
}

func NewCreate_PhSensor_C(useCase *use_case.Create_PhSensor) *Create_PhSensor_C {
	return &Create_PhSensor_C{UseCase: useCase}
}

func (c *Create_PhSensor_C) Execute(ctx *gin.Context) {
	var sensor entities.PhSensor

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
