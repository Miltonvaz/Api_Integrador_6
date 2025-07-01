package routes_a

import (
	"Integrador/src/core/security"
	"Integrador/src/sensor_alcohol/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	createController *controllers.Create_AlcoholSensor_C,
	getByIDController *controllers.GetById_AlcoholSensor_C,
	getAllController *controllers.GetAll_AlcoholSensor_C,
	deleteController *controllers.DeleteAlcoholSensor_C,
) {
	api := router.Group("/api/sensor-alcohol")
	{
		api.POST("/create", createController.Execute)
		api.GET("/measurement/:id/:userID", security.JWTMiddleware(), getByIDController.Execute)
		api.GET("/all/:userID", security.JWTMiddleware(), getAllController.Execute)
		api.DELETE("/delete/:id/:userID", security.JWTMiddleware(), deleteController.Execute)
	}
}
