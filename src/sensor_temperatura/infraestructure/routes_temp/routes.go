package routes_temp

import (
	"Integrador/src/core/security"
	"Integrador/src/sensor_temperatura/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	createController *controllers.Create_TemperatureSensor_C,
	getByIDController *controllers.GetById_TemperatureSensor_C,
	getAllController *controllers.GetAll_TemperatureSensor_C,
	deleteController *controllers.DeleteTemperatureSensorr_C,
) {
	api := router.Group("/api/sensor-temperatura")
	{
		api.POST("/create", createController.Execute)
		api.GET("/measurement/:id/:userID", security.JWTMiddleware(), getByIDController.Execute)
		api.GET("/all/:userID", security.JWTMiddleware(), getAllController.Execute)
		api.DELETE("/delete/:id/:userID", security.JWTMiddleware(), deleteController.Execute)
	}
}
