package routes_f

import (
	"Integrador/src/core/security"
	"Integrador/src/fermentation/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterFermentationRoutes(
	router *gin.Engine,
	createController *controllers.Create_Fermentation_C,
	getByIDController *controllers.GetById_Fermentation_C,
	getAllController *controllers.GetAll_Fermentation_C,
	deleteController *controllers.DeleteFermentation_C,
	updateController *controllers.UpdateFermentationC,
) {
	api := router.Group("/api/fermentation")
	{
		api.POST("/create", createController.Execute)
		api.GET("/record/:id/:userID", security.JWTMiddleware(), getByIDController.Execute)
		api.GET("/all/:userID", security.JWTMiddleware(), getAllController.Execute)
		api.DELETE("/delete/:id/:userID", security.JWTMiddleware(), deleteController.Execute)
		api.PUT("/update/:id/:userID", security.JWTMiddleware(), updateController.Execute)
	}
}
