package controllers

import (
	"Integrador/src/sensor_densidad_o/application/use_case"
	"Integrador/src/sensor_densidad_o/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Create_DensitySensor_C struct {
	UseCase *use_case.Create_DensitySensor
}

func NewCreate_DensitySensor_C(useCase *use_case.Create_DensitySensor) *Create_DensitySensor_C {
	return &Create_DensitySensor_C{UseCase: useCase}
}

func (c *Create_DensitySensor_C) Execute(ctx *gin.Context) {
	var sensor entities.DensitySensor

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
