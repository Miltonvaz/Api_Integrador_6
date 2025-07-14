package controllers

import (
	"Integrador/src/fermentation/application/use_case"
	"Integrador/src/fermentation/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateFermentationC struct {
	useCase *use_case.UpdateFermentation
}

func NewUpdateFermentationC(useCase *use_case.UpdateFermentation) *UpdateFermentationC {
	return &UpdateFermentationC{useCase: useCase}
}

func (c *UpdateFermentationC) Execute(ctx *gin.Context) {

	userIDStr := ctx.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userID inválido"})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de fermentación inválido"})
		return
	}

	var updatedFermentation entities.Fermentation
	if err := ctx.ShouldBindJSON(&updatedFermentation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido", "details": err.Error()})
		return
	}

	updatedFermentation.ID = id
	updatedFermentation.OperatorID = userID

	data, err := c.useCase.Execute(updatedFermentation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el registro", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}
