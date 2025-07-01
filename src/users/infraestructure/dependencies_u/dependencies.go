package dependencies_u

import (
	"Integrador/src/core"
	"Integrador/src/users/application/repositories"
	"Integrador/src/users/application/use_case"
	"Integrador/src/users/infraestructure/adapters"
	"Integrador/src/users/infraestructure/controllers"
	"log"
)

func Init(pool *core.Conn_MySQL) (
	*controllers.CreateUserController,
	*controllers.ViewUserController,
	*controllers.EditUserController,
	*controllers.DeleteUserController,
	*controllers.ViewUserByIdController,
	*controllers.AuthController,
	error,
) {

	ps := adapters.NewMySQL(pool.DB)

	rabbitMQAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error initializing RabbitMQ: %v", err)
	}

	serviceNotification := repositories.NewServiceNotification(rabbitMQAdapter)

	createClient := use_case.NewCreateUser(ps, serviceNotification)
	viewClient := use_case.NewListUser(ps)
	editClient := use_case.NewEditUser(ps)
	deleteClient := use_case.NewDeleteUser(ps)
	viewClientById := use_case.NewUserById(ps)
	authService := use_case.NewAuthService(ps)

	authController := controllers.NewAuthController(authService)
	createClientController := controllers.NewCreateUserController(*createClient)
	viewClientController := controllers.NewViewUserController(*viewClient)
	editClientController := controllers.NewEditUserController(*editClient)
	deleteClientController := controllers.NewDeleteUserController(*deleteClient)
	viewClientByIdController := controllers.NewViewUserByIdController(*viewClientById)

	return createClientController, viewClientController, editClientController, deleteClientController, viewClientByIdController, authController, nil
}
