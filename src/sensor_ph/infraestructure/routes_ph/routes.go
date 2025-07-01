package routes_ph

import (
	"Integrador/src/core/security"
	"Integrador/src/sensor_ph/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	createController *controllers.Create_PhSensor_C,
	getByIDController *controllers.GetById_PhSensor_C,
	getAllController *controllers.GetAll_PhSensor_C,
	deleteController *controllers.DeletePhSensor_C,
) {
	api := router.Group("/api/sensor-ph")
	{
		api.POST("/create", createController.Execute)
		api.GET("/measurement/:id/:userID", security.JWTMiddleware(), getByIDController.Execute)
		api.GET("/all/:userID", security.JWTMiddleware(), getAllController.Execute)
		api.DELETE("/delete/:id/:userID", security.JWTMiddleware(), deleteController.Execute)
	}
}
