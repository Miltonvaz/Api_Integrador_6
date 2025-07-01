package use_case

import (
	"Integrador/src/sensor_turbuidez/application/repositories"
	"Integrador/src/sensor_turbuidez/domain"
	"Integrador/src/sensor_turbuidez/domain/entities"

	"log"
)

type Create_TurbiditySensor struct {
	appointmentRepo     domain.TurbiditySensor
	serviceNotification *repositories.ServiceNotification
}

func NewCreate_TurbiditySensor(appointmentRepo domain.TurbiditySensor, serviceNotification *repositories.ServiceNotification) *Create_TurbiditySensor {
	return &Create_TurbiditySensor{
		appointmentRepo:     appointmentRepo,
		serviceNotification: serviceNotification,
	}
}

func (c *Create_TurbiditySensor) Execute(appointment entities.TurbiditySensor) (entities.TurbiditySensor, error) {
	created, err := c.appointmentRepo.Save(appointment)
	if err != nil {
		log.Printf("Error saving data: %v", err)
		return entities.TurbiditySensor{}, err
	}
	err = c.serviceNotification.NotifyAppointmentCreated(created)
	if err != nil {
		log.Printf("Error notifying about created appointment: %v", err)
		return entities.TurbiditySensor{}, err
	}

	return created, nil
}
