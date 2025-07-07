package controllers

import (
	"Integrador/src/motor/application/use_case"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteMotor_C struct {
	useCase *use_case.Delete_Motor
}

func NewDeleteMotor_C(useCase *use_case.Delete_Motor) *DeleteMotor_C {
	return &DeleteMotor_C{useCase: useCase}
}

func (ct *DeleteMotor_C) Execute(ctx *gin.Context) {
	idStr := ctx.Param("id")
	userIDStr := ctx.Param("userID")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userID inválido"})
		return
	}

	if err := ct.useCase.Execute(id, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar los datos"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Datos eliminados correctamente"})
}
