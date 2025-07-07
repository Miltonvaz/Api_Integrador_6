package routes_m

import (
	"Integrador/src/core/security"
	"Integrador/src/motor/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	createController *controllers.Create_Motor_C,
	getByIDController *controllers.GetById_Motor_C,
	getAllController *controllers.GetAll_Motor_C,
	deleteController *controllers.DeleteMotor_C,
) {
	api := router.Group("/api/motor")
	{
		api.POST("/create", createController.Execute)
		api.GET("/measurement/:id/:userID", security.JWTMiddleware(), getByIDController.Execute)
		api.GET("/all/:userID", security.JWTMiddleware(), getAllController.Execute)
		api.DELETE("/delete/:id/:userID", security.JWTMiddleware(), deleteController.Execute)
	}
}
