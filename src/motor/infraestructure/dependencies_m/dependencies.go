package dependencies_m

import (
	"log"

	"Integrador/src/core"
	"Integrador/src/motor/application/repositories"
	"Integrador/src/motor/application/use_case"
	"Integrador/src/motor/infraestructure/adapters"
	"Integrador/src/motor/infraestructure/controllers"
)

func Init(pool *core.Conn_MySQL) (
	createController *controllers.Create_Motor_C,
	getByIDController *controllers.GetById_Motor_C,
	getAllController *controllers.GetAll_Motor_C,
	deleteController *controllers.DeleteMotor_C,
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

	createUseCase := use_case.NewCreate_Motor(repository, serviceNotification)
	getByIDUseCase := use_case.NewGetByID_Motor(repository)
	getAllUseCase := use_case.NewGet_All_Motor(repository)
	deleteUseCase := use_case.NewDelete_Motor(repository)

	createController = controllers.NewCreate_Motor_C(createUseCase)
	getByIDController = controllers.NewGetById_Motor_C(getByIDUseCase)
	getAllController = controllers.NewGetAll_Motor_C(getAllUseCase)
	deleteController = controllers.NewDeleteMotor_C(deleteUseCase)

	return createController,
		getByIDController,
		getAllController,
		deleteController,
		serviceNotification,
		nil
}
