package controllers

import (
	"Integrador/src/sensor_conductividad/application/use_case"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetAll_ConductividtySensor_C struct {
	useCase *use_case.Get_All_ConductividtySensor
}

func NewGetAll_ConductividtySensor_C(useCase *use_case.Get_All_ConductividtySensor) *GetAll_ConductividtySensor_C {
	return &GetAll_ConductividtySensor_C{useCase: useCase}
}

func (gt *GetAll_ConductividtySensor_C) Execute(ctx *gin.Context) {
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
