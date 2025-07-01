package use_case

import (
	"Integrador/src/sensor_temperatura/application/repositories"
	"Integrador/src/sensor_temperatura/domain"
	"Integrador/src/sensor_temperatura/domain/entities"
	"log"
)

type Create_TemperatureSensor struct {
	appointmentRepo     domain.ITemperatureSensor
	serviceNotification *repositories.ServiceNotification
}

func NewCreate_TemperatureSensore(appointmentRepo domain.ITemperatureSensor, serviceNotification *repositories.ServiceNotification) *Create_TemperatureSensor {
	return &Create_TemperatureSensor{
		appointmentRepo:     appointmentRepo,
		serviceNotification: serviceNotification,
	}
}

func (c *Create_TemperatureSensor) Execute(appointment entities.TemperatureSensor) (entities.TemperatureSensor, error) {
	created, err := c.appointmentRepo.Save(appointment)
	if err != nil {
		log.Printf("Error saving data: %v", err)
		return entities.TemperatureSensor{}, err
	}
	err = c.serviceNotification.NotifyAppointmentCreated(created)
	if err != nil {
		log.Printf("Error notifying about created appointment: %v", err)
		return entities.TemperatureSensor{}, err
	}

	return created, nil
}
