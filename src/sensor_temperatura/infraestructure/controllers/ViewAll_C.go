package controllers

import (
	"Integrador/src/sensor_temperatura/application/use_case"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetAll_TemperatureSensor_C struct {
	useCase *use_case.Get_All_TemperatureSensor
}

func NewGetAll_TemperatureSensor_C(useCase *use_case.Get_All_TemperatureSensor) *GetAll_TemperatureSensor_C {
	return &GetAll_TemperatureSensor_C{useCase: useCase}
}

func (gt *GetAll_TemperatureSensor_C) Execute(ctx *gin.Context) {
	userIDStr := ctx.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userID inválido"})
		return
	}

	data, err := gt.useCase.Execute(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los registros", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
