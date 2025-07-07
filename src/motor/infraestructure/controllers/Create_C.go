package controllers

import (
	"Integrador/src/motor/application/use_case"
	"Integrador/src/motor/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Create_Motor_C struct {
	UseCase *use_case.Create_Motor
}

func NewCreate_Motor_C(useCase *use_case.Create_Motor) *Create_Motor_C {
	return &Create_Motor_C{UseCase: useCase}
}

func (c *Create_Motor_C) Execute(ctx *gin.Context) {
	var sensor entities.MotorFermentador

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
