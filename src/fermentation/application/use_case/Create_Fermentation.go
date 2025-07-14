package use_case

import (
	"Integrador/src/fermentation/application/repositories"
	"Integrador/src/fermentation/domain"
	"Integrador/src/fermentation/domain/entities"
	"log"
)

type Create_Fermentation struct {
	appointmentRepo     domain.FermentationRepository
	serviceNotification *repositories.ServiceNotification
}

func NewCreate_Fermentation(appointmentRepo domain.FermentationRepository, serviceNotification *repositories.ServiceNotification) *Create_Fermentation {
	return &Create_Fermentation{
		appointmentRepo:     appointmentRepo,
		serviceNotification: serviceNotification,
	}
}

func (c *Create_Fermentation) Execute(appointment entities.Fermentation) (entities.Fermentation, error) {
	created, err := c.appointmentRepo.Save(appointment)
	if err != nil {
		log.Printf("Error saving data: %v", err)
		return entities.Fermentation{}, err
	}
	err = c.serviceNotification.NotifyAppointmentCreated(created)
	if err != nil {
		log.Printf("Error notifying about created appointment: %v", err)
		return entities.Fermentation{}, err
	}

	return created, nil
}
