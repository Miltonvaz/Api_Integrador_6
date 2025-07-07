package dependencies_d

import (
	"Integrador/src/sensor_densidad_o/application/repositories"
	"Integrador/src/sensor_densidad_o/application/use_case"
	"Integrador/src/sensor_densidad_o/infraestructure/adapters"
	"Integrador/src/sensor_densidad_o/infraestructure/controllers"
	"log"

	"Integrador/src/core"
)

func Init(pool *core.Conn_MySQL) (
	createController *controllers.Create_DensitySensor_C,
	getByIDController *controllers.GetById_DensitySensor_C,
	getAllController *controllers.GetAll_DensitySensor_C,
	deleteController *controllers.DeleteDensitySensor_C,
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

	createUseCase := use_case.NewCreate_DensitySensor(repository, serviceNotification)
	getByIDUseCase := use_case.NewGetByID_DensitySensor(repository)
	getAllUseCase := use_case.NewGet_All_DensitySensor(repository)
	deleteUseCase := use_case.NewDelete_DensitySensor(repository)

	createController = controllers.NewCreate_DensitySensor_C(createUseCase)
	getByIDController = controllers.NewGetById_DensitySensor_C(getByIDUseCase)
	getAllController = controllers.NewGetAll_DensitySensor_C(getAllUseCase)
	deleteController = controllers.NewDeleteDensitySensor_C(deleteUseCase)

	return createController,
		getByIDController,
		getAllController,
		deleteController,
		serviceNotification,
		nil
}
