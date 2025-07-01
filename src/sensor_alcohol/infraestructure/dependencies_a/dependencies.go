package dependencies_a

import (
	"log"

	"Integrador/src/core"
	"Integrador/src/sensor_alcohol/application/repositories"
	"Integrador/src/sensor_alcohol/application/use_case"
	"Integrador/src/sensor_alcohol/infraestructure/adapters"
	"Integrador/src/sensor_alcohol/infraestructure/controllers"
)

func Init(pool *core.Conn_MySQL) (
	createController *controllers.Create_AlcoholSensor_C,
	getByIDController *controllers.GetById_AlcoholSensor_C,
	getAllController *controllers.GetAll_AlcoholSensor_C,
	deleteController *controllers.DeleteAlcoholSensor_C,
	serviceNotification *repositories.ServiceNotification,
	err error,
) {
	repository := adapters.NewMySQL(pool.DB)

	rabbitMQAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error initializing RabbitMQ: %v", err)
		return nil, nil, nil, nil, nil, err
	}

	serviceNotification = repositories.NewServiceNotification(rabbitMQAdapter)

	createUseCase := use_case.NewCreate_AlcoholSensor(repository, serviceNotification)
	getByIDUseCase := use_case.NewGetByID_AlcoholSensor(repository)
	getAllUseCase := use_case.NewGet_All_AlcoholSensor(repository)
	deleteUseCase := use_case.NewDelete_AlcoholSensor(repository)

	createController = controllers.NewCreate_AlcoholSensor_C(createUseCase)
	getByIDController = controllers.NewGetById_AlcoholSensor_C(getByIDUseCase)
	getAllController = controllers.NewGetAll_AlcoholSensor_C(getAllUseCase)
	deleteController = controllers.NewDeleteAlcoholSensor_C(deleteUseCase)

	return createController,
		getByIDController,
		getAllController,
		deleteController,
		serviceNotification,
		nil
}
