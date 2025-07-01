package use_case

import (
	"Integrador/src/sensor_densidad_o/application/repositories"
	"Integrador/src/sensor_densidad_o/domain"
	"Integrador/src/sensor_densidad_o/domain/entities"
	"log"
)

type Create_DensitySensor struct {
	appointmentRepo     domain.IDensitySensor
	serviceNotification *repositories.ServiceNotification
}

func NewCreate_DensitySensor(appointmentRepo domain.IDensitySensor, serviceNotification *repositories.ServiceNotification) *Create_DensitySensor {
	return &Create_DensitySensor{
		appointmentRepo:     appointmentRepo,
		serviceNotification: serviceNotification,
	}
}

func (c *Create_DensitySensor) Execute(appointment entities.DensitySensor) (entities.DensitySensor, error) {
	created, err := c.appointmentRepo.Save(appointment)
	if err != nil {
		log.Printf("Error saving data: %v", err)
		return entities.DensitySensor{}, err
	}
	err = c.serviceNotification.NotifyAppointmentCreated(created)
	if err != nil {
		log.Printf("Error notifying about created appointment: %v", err)
		return entities.DensitySensor{}, err
	}

	return created, nil
}
