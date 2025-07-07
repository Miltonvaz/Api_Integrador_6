package dependencies_c

import (
	"Integrador/src/sensor_conductividad/application/repositories"
	"Integrador/src/sensor_conductividad/application/use_case"
	"Integrador/src/sensor_conductividad/infraestructure/adapters"
	"Integrador/src/sensor_conductividad/infraestructure/controllers"
	"log"

	"Integrador/src/core"
)

func Init(pool *core.Conn_MySQL) (
	createController *controllers.Create_ConductividtySensor_C,
	getByIDController *controllers.GetById_ConductividtySensor_C,
	getAllController *controllers.GetAll_ConductividtySensor_C,
	deleteController *controllers.DeleteConductividtySensor_C,
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

	createUseCase := use_case.NewCreate_ConductividtySensor(repository, serviceNotification)
	getByIDUseCase := use_case.NewGetByID_ConductividtySensor(repository)
	getAllUseCase := use_case.NewGet_All_ConductividtySensor(repository)
	deleteUseCase := use_case.NewDelete_ConductividtySensor(repository)

	createController = controllers.NewCreate_ConductividtySensor_C(createUseCase)
	getByIDController = controllers.NewGetById_ConductividtySensor_C(getByIDUseCase)
	getAllController = controllers.NewGetAll_ConductividtySensor_C(getAllUseCase)
	deleteController = controllers.NewDeleteConductividtySensor_C(deleteUseCase)

	return createController,
		getByIDController,
		getAllController,
		deleteController,
		serviceNotification,
		nil
}
