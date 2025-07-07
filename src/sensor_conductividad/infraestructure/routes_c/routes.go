package routes_c

import (
	"Integrador/src/core/security"
	"Integrador/src/sensor_conductividad/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	createController *controllers.Create_ConductividtySensor_C,
	getByIDController *controllers.GetById_ConductividtySensor_C,
	getAllController *controllers.GetAll_ConductividtySensor_C,
	deleteController *controllers.DeleteConductividtySensor_C,
) {
	api := router.Group("/api/sensor-conductividad")
	{
		api.POST("/create", createController.Execute)
		api.GET("/measurement/:id/:userID", security.JWTMiddleware(), getByIDController.Execute)
		api.GET("/all/:userID", security.JWTMiddleware(), getAllController.Execute)
		api.DELETE("/delete/:id/:userID", security.JWTMiddleware(), deleteController.Execute)
	}
}
