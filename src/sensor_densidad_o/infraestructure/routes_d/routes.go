package routes_d

import (
	"Integrador/src/core/security"
	"Integrador/src/sensor_densidad_o/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	createController *controllers.Create_DensitySensor_C,
	getByIDController *controllers.GetById_DensitySensor_C,
	getAllController *controllers.GetAll_DensitySensor_C,
	deleteController *controllers.DeleteDensitySensor_C,
) {
	api := router.Group("/api/sensor-densidad")
	{
		api.POST("/create", createController.Execute)
		api.GET("/measurement/:id/:userID", security.JWTMiddleware(), getByIDController.Execute)
		api.GET("/all/:userID", security.JWTMiddleware(), getAllController.Execute)
		api.DELETE("/delete/:id/:userID", security.JWTMiddleware(), deleteController.Execute)
	}
}
