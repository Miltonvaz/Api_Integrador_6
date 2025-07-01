package use_case

import (
	"Integrador/src/sensor_ph/application/repositories"
	"Integrador/src/sensor_ph/domain"
	entities "Integrador/src/sensor_ph/domain/entities"
	"log"
)

type Create_PhSensor struct {
	appointmentRepo     domain.IPhSensor
	serviceNotification *repositories.ServiceNotification
}

func NewCreate_PhSensor(appointmentRepo domain.IPhSensor, serviceNotification *repositories.ServiceNotification) *Create_PhSensor {
	return &Create_PhSensor{
		appointmentRepo:     appointmentRepo,
		serviceNotification: serviceNotification,
	}
}

func (c *Create_PhSensor) Execute(appointment entities.PhSensor) (entities.PhSensor, error) {
	created, err := c.appointmentRepo.Save(appointment)
	if err != nil {
		log.Printf("Error saving data: %v", err)
		return entities.PhSensor{}, err
	}
	err = c.serviceNotification.NotifyAppointmentCreated(created)
	if err != nil {
		log.Printf("Error notifying about created appointment: %v", err)
		return entities.PhSensor{}, err
	}

	return created, nil
}
