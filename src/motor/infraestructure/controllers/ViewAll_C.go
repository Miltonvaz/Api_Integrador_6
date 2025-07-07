package controllers

import (
	"Integrador/src/motor/application/use_case"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetAll_Motor_C struct {
	useCase *use_case.Get_All_Motor
}

func NewGetAll_Motor_C(useCase *use_case.Get_All_Motor) *GetAll_Motor_C {
	return &GetAll_Motor_C{useCase: useCase}
}

func (gt *GetAll_Motor_C) Execute(ctx *gin.Context) {
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
