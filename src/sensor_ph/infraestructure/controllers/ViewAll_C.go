package controllers

import (
	"Integrador/src/sensor_ph/application/use_case"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetAll_PhSensor_C struct {
	useCase *use_case.Get_All_PhSensor
}

func NewGetAll_PhSensor_C(useCase *use_case.Get_All_PhSensor) *GetAll_PhSensor_C {
	return &GetAll_PhSensor_C{useCase: useCase}
}

func (gt *GetAll_PhSensor_C) Execute(ctx *gin.Context) {
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
