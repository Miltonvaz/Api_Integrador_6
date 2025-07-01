package dependencies_t

import (
	"Integrador/src/core"
	"Integrador/src/sensor_turbuidez/application/repositories"
	"Integrador/src/sensor_turbuidez/application/use_case"
	"Integrador/src/sensor_turbuidez/infraestructure/adapters"
	"Integrador/src/sensor_turbuidez/infraestructure/controllers"
	"log"
)

func Init(pool *core.Conn_MySQL) (
	createController *controllers.Create_TurbiditySensor_C,
	getByIDController *controllers.GetById_TurbiditySensor_C,
	getAllController *controllers.GetAll_TurbiditySensor_C,
	deleteController *controllers.DeleteTurbiditySensor_C,
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

	createUseCase := use_case.NewCreate_TurbiditySensor(repository, serviceNotification)
	getByIDUseCase := use_case.NewGetByID_TTurbiditySensor(repository)
	getAllUseCase := use_case.NewGet_All_TurbiditySensor(repository)
	deleteUseCase := use_case.NewDelete_TurbiditySensor(repository)

	createController = controllers.NewCreate_TurbiditySensor_C(createUseCase)
	getByIDController = controllers.NewGetById_TurbiditySensor_C(getByIDUseCase)
	getAllController = controllers.NewGetAll_TurbiditySensor_C(getAllUseCase)
	deleteController = controllers.NewDeleteTurbiditySensor_C(deleteUseCase)

	return createController,
		getByIDController,
		getAllController,
		deleteController,
		serviceNotification,
		nil
}
