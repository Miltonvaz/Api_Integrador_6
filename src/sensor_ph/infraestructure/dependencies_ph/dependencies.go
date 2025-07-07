package dependencies_ph

import (
	"Integrador/src/core"
	"Integrador/src/sensor_ph/application/repositories"
	"Integrador/src/sensor_ph/application/use_case"
	"Integrador/src/sensor_ph/infraestructure/adapters"
	"Integrador/src/sensor_ph/infraestructure/controllers"
	"log"
)

func Init(pool *core.Conn_MySQL) (
	createController *controllers.Create_PhSensor_C,
	getByIDController *controllers.GetById_PhSensor_C,
	getAllController *controllers.GetAll_PhSensor_C,
	deleteController *controllers.DeletePhSensor_C,
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

	createUseCase := use_case.NewCreate_PhSensor(repository, serviceNotification)
	getByIDUseCase := use_case.NewGetByID_PhSensor(repository)
	getAllUseCase := use_case.NewGet_All_PhSensor(repository)
	deleteUseCase := use_case.NewDelete_PhSensor(repository)

	createController = controllers.NewCreate_PhSensor_C(createUseCase)
	getByIDController = controllers.NewGetById_PhSensor_C(getByIDUseCase)
	getAllController = controllers.NewGetAll_PhSensor_C(getAllUseCase)
	deleteController = controllers.NewDeletePhSensor_C(deleteUseCase)

	return createController,
		getByIDController,
		getAllController,
		deleteController,
		serviceNotification,
		nil
}
