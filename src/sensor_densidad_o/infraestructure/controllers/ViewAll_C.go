package controllers

import (
	"Integrador/src/sensor_densidad_o/application/use_case"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetAll_DensitySensor_C struct {
	useCase *use_case.Get_All_DensitySensor
}

func NewGetAll_DensitySensor_C(useCase *use_case.Get_All_DensitySensor) *GetAll_DensitySensor_C {
	return &GetAll_DensitySensor_C{useCase: useCase}
}

func (gt *GetAll_DensitySensor_C) Execute(ctx *gin.Context) {
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
