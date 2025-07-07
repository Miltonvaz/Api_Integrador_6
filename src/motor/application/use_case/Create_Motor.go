package use_case

import (
	"Integrador/src/motor/application/repositories"
	"Integrador/src/motor/domain"
	"Integrador/src/motor/domain/entities"
	"log"
)

type Create_Motor struct {
	appointmentRepo     domain.MotorFermentadorRepository
	serviceNotification *repositories.ServiceNotification
}

func NewCreate_Motor(appointmentRepo domain.MotorFermentadorRepository, serviceNotification *repositories.ServiceNotification) *Create_Motor {
	return &Create_Motor{
		appointmentRepo:     appointmentRepo,
		serviceNotification: serviceNotification,
	}
}

func (c *Create_Motor) Execute(appointment entities.MotorFermentador) (entities.MotorFermentador, error) {
	created, err := c.appointmentRepo.Save(appointment)
	if err != nil {
		log.Printf("Error saving data: %v", err)
		return entities.MotorFermentador{}, err
	}
	err = c.serviceNotification.NotifyAppointmentCreated(created)
	if err != nil {
		log.Printf("Error notifying about created appointment: %v", err)
		return entities.MotorFermentador{}, err
	}

	return created, nil
}
