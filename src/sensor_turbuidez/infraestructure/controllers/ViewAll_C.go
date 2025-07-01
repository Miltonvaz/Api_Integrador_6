package controllers

import (
	"Integrador/src/sensor_turbuidez/application/use_case"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetAll_TurbiditySensor_C struct {
	useCase *use_case.Get_All_TurbiditySensor
}

func NewGetAll_TurbiditySensor_C(useCase *use_case.Get_All_TurbiditySensor) *GetAll_TurbiditySensor_C {
	return &GetAll_TurbiditySensor_C{useCase: useCase}
}

func (gt *GetAll_TurbiditySensor_C) Execute(ctx *gin.Context) {
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
