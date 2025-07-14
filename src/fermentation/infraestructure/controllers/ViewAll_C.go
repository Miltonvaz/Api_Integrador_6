package controllers

import (
	"Integrador/src/fermentation/application/use_case"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetAll_Fermentation_C struct {
	useCase *use_case.Get_All_Fermentation
}

func NewGetAll_Fermentation_C(useCase *use_case.Get_All_Fermentation) *GetAll_Fermentation_C {
	return &GetAll_Fermentation_C{useCase: useCase}
}

func (gt *GetAll_Fermentation_C) Execute(ctx *gin.Context) {
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
