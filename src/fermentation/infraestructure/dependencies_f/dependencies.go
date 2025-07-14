package dependencies_f

import (
	"log"

	"Integrador/src/core"
	"Integrador/src/fermentation/application/repositories"
	"Integrador/src/fermentation/application/use_case"
	"Integrador/src/fermentation/infraestructure/adapters"
	"Integrador/src/fermentation/infraestructure/controllers"
)

func Init(
	pool *core.Conn_MySQL,
) (
	createController *controllers.Create_Fermentation_C,
	getByIDController *controllers.GetById_Fermentation_C,
	getAllController *controllers.GetAll_Fermentation_C,
	deleteController *controllers.DeleteFermentation_C,
	updateController *controllers.UpdateFermentationC,
	serviceNotification *repositories.ServiceNotification,
	err error,
) {
	repository := adapters.NewMySQL(pool.DB)

	rabbitMQAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error initializing RabbitMQ: %v", err)
		return nil, nil, nil, nil, nil, nil, err
	}

	serviceNotification = repositories.NewServiceNotification(rabbitMQAdapter)

	createUseCase := use_case.NewCreate_Fermentation(repository, serviceNotification)
	getByIDUseCase := use_case.NewGetByID_Fermentation(repository)
	getAllUseCase := use_case.NewGet_All_Fermentation(repository)
	deleteUseCase := use_case.NewDelete_Fermentation(repository)
	updateUseCase := use_case.NewUpdateFermentation(repository)

	createController = controllers.NewCreate_Fermentation_C(createUseCase)
	getByIDController = controllers.NewGetById_Fermentation_C(getByIDUseCase)
	getAllController = controllers.NewGetAll_Fermentation_C(getAllUseCase)
	deleteController = controllers.NewDeleteFermentation_C(deleteUseCase)
	updateController = controllers.NewUpdateFermentationC(updateUseCase)

	return createController,
		getByIDController,
		getAllController,
		deleteController,
		updateController,
		serviceNotification,
		nil
}
