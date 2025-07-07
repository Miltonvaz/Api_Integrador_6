package routes_t

import (
	"Integrador/src/core/security"
	"Integrador/src/sensor_turbuidez/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	createController *controllers.Create_TurbiditySensor_C,
	getByIDController *controllers.GetById_TurbiditySensor_C,
	getAllController *controllers.GetAll_TurbiditySensor_C,
	deleteController *controllers.DeleteTurbiditySensor_C,
) {
	api := router.Group("/api/sensor-turbuidez")
	{
		api.POST("/create", createController.Execute)
		api.GET("/measurement/:id/:userID", security.JWTMiddleware(), getByIDController.Execute)
		api.GET("/all/:userID", security.JWTMiddleware(), getAllController.Execute)
		api.DELETE("/delete/:id/:userID", security.JWTMiddleware(), deleteController.Execute)
	}
}
