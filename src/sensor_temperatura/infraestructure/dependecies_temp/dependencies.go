package dependecies_temp

import (
	"Integrador/src/core"
	"Integrador/src/sensor_temperatura/application/repositories"
	"Integrador/src/sensor_temperatura/application/use_case"
	"Integrador/src/sensor_temperatura/infraestructure/adapters"
	"Integrador/src/sensor_temperatura/infraestructure/controllers"
	"log"
)

func Init(pool *core.Conn_MySQL) (
	createController *controllers.Create_TemperatureSensor_C,
	getByIDController *controllers.GetById_TemperatureSensor_C,
	getAllController *controllers.GetAll_TemperatureSensor_C,
	deleteController *controllers.DeleteTemperatureSensorr_C,
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

	createUseCase := use_case.NewCreate_TemperatureSensore(repository, serviceNotification)
	getByIDUseCase := use_case.NewGetByID_TemperatureSensor(repository)
	getAllUseCase := use_case.NewGet_All_TemperatureSensor(repository)
	deleteUseCase := use_case.NewDelete_TemperatureSensor(repository)

	createController = controllers.NewCreate_TemperatureSensor_C(createUseCase)
	getByIDController = controllers.NewGetById_TemperatureSensor_C(getByIDUseCase)
	getAllController = controllers.NewGetAll_TemperatureSensor_C(getAllUseCase)
	deleteController = controllers.NewDeleteTemperatureSensor_C(deleteUseCase)

	return createController,
		getByIDController,
		getAllController,
		deleteController,
		serviceNotification,
		nil
}
