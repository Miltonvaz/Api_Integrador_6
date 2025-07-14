package main

import (
	"log"
	"os"
	"time"

	"Integrador/src/core"
	"Integrador/src/fermentation/infraestructure/dependencies_f"
	"Integrador/src/fermentation/infraestructure/routes_f"
	"Integrador/src/motor/infraestructure/dependencies_m"
	"Integrador/src/motor/infraestructure/routes_m"
	"Integrador/src/sensor_alcohol/infraestructure/dependencies_a"
	"Integrador/src/sensor_alcohol/infraestructure/routes_a"
	"Integrador/src/sensor_conductividad/infraestructure/dependencies_c"
	"Integrador/src/sensor_conductividad/infraestructure/routes_c"
	"Integrador/src/sensor_densidad_o/infraestructure/dependencies_d"
	"Integrador/src/sensor_densidad_o/infraestructure/routes_d"
	"Integrador/src/sensor_ph/infraestructure/dependencies_ph"
	"Integrador/src/sensor_ph/infraestructure/routes_ph"
	"Integrador/src/sensor_temperatura/infraestructure/dependecies_temp"
	"Integrador/src/sensor_temperatura/infraestructure/routes_temp"
	"Integrador/src/sensor_turbuidez/infraestructure/dependencies_t"
	"Integrador/src/sensor_turbuidez/infraestructure/routes_t"
	"Integrador/src/users/infraestructure/dependencies_u"
	"Integrador/src/users/infraestructure/routes_u"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func startServer() {
	log.Println("Iniciando servidor...")

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.ExposeHeaders = []string{"Content-Length", "Authorization"}
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))

	if err := initializeDependencies(router); err != nil {
		log.Fatalf("Error al inicializar dependencias: %v", err)
		return
	}

	go func() {
		time.Sleep(2 * time.Minute)
		log.Println("Tiempo límite alcanzado. Cerrando servidor para reinicio...")
		os.Exit(0)
	}()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error en el servidor: %v", err)
	}
}

func initializeDependencies(router *gin.Engine) error {
	pool := core.GetDBPool()

	alcoholC, alcoholByID, alcoholAll, alcoholDel, _, err := dependencies_a.Init(pool)
	if err != nil {
		return err
	}
	routes_a.RegisterRoutes(router, alcoholC, alcoholByID, alcoholAll, alcoholDel)

	phC, phByID, phAll, phDel, _, err := dependencies_ph.Init(pool)
	if err != nil {
		return err
	}
	routes_ph.RegisterRoutes(router, phC, phByID, phAll, phDel)

	densidadC, densidadByID, densidadAll, densidadDel, _, err := dependencies_d.Init(pool)
	if err != nil {
		return err
	}
	routes_d.RegisterRoutes(router, densidadC, densidadByID, densidadAll, densidadDel)

	condC, condByID, condAll, condDel, _, err := dependencies_c.Init(pool)
	if err != nil {
		return err
	}
	routes_c.RegisterRoutes(router, condC, condByID, condAll, condDel)

	turbC, turbByID, turbAll, turbDel, _, err := dependencies_t.Init(pool)
	if err != nil {
		return err
	}
	routes_t.RegisterRoutes(router, turbC, turbByID, turbAll, turbDel)

	tempC, tempByID, tempAll, tempDel, _, err := dependecies_temp.Init(pool)
	if err != nil {
		return err
	}
	routes_temp.RegisterRoutes(router, tempC, tempByID, tempAll, tempDel)

	motC, motByID, motAll, motDel, _, err := dependencies_m.Init(pool)
	if err != nil {
		return err
	}
	routes_m.RegisterRoutes(router, motC, motByID, motAll, motDel)

	createUserC, viewUserC, editUserC, deleteUserC, viewByIdUserC, loginC, err := dependencies_u.Init(pool)
	if err != nil {
		return err
	}
	routes_u.RegisterClientRoutes(router, createUserC, viewUserC, editUserC, deleteUserC, viewByIdUserC, loginC)

	fermentCreateC, fermentByID, fermentAll, fermentDel, fermentUpdateC, _, err := dependencies_f.Init(pool)
	if err != nil {
		return err
	}
	routes_f.RegisterFermentationRoutes(router, fermentCreateC, fermentByID, fermentAll, fermentDel, fermentUpdateC)

	return nil
}

func main() {
	startServer()
}
