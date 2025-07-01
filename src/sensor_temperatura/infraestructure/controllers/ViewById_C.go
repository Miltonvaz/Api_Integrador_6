package controllers

import (
	"Integrador/src/sensor_temperatura/application/use_case"
	"Integrador/src/sensor_temperatura/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetById_TemperatureSensor_C struct {
	useCase *use_case.GetByID_TemperatureSensor
}

func NewGetById_TemperatureSensor_C(useCase *use_case.GetByID_TemperatureSensor) *GetById_TemperatureSensor_C {
	return &GetById_TemperatureSensor_C{useCase: useCase}
}
func (gh *GetById_TemperatureSensor_C) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	userIDStr := ctx.Param("userID")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID no válido"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userID inválido"})
		return
	}

	measurement, err := gh.useCase.Execute(id, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la temperatura", "details": err.Error()})
		return
	}
	if (measurement == entities.TemperatureSensor{}) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No se encontró el registro con el ID proporcionado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": measurement,
	})
}
