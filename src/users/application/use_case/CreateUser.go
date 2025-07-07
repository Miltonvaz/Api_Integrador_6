package use_case

import (
	"Integrador/src/core/security"
	"Integrador/src/users/application/repositories"
	"Integrador/src/users/domain"
	"Integrador/src/users/domain/entities"
	"log"
)

type CreateUser struct {
	db                  domain.IUser
	serviceNotification *repositories.ServiceNotification
}

func NewCreateUser(db domain.IUser, serviceNotification *repositories.ServiceNotification) *CreateUser {
	return &CreateUser{
		db:                  db,
		serviceNotification: serviceNotification,
	}
}

func (cc *CreateUser) Execute(client entities.User) error {
	hashedPassword, err := security.HashPassword(client.Password)
	if err != nil {
		return err
	}
	client.Password = hashedPassword

	if err := cc.db.Save(client); err != nil {
		return err
	}

	err = cc.serviceNotification.NotifyAppointmentCreated(client)
	if err != nil {
		log.Printf("Error notificando la creación del usuario: %v", err)
	}

	return nil
}
