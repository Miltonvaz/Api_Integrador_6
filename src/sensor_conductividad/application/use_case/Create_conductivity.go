package use_case

import (
	"Integrador/src/sensor_conductividad/application/repositories"
	"Integrador/src/sensor_conductividad/domain"
	"Integrador/src/sensor_conductividad/domain/entities"
	"log"
)

type Create_ConductividtySensor struct {
	appointmentRepo     domain.IConductividtySensor
	serviceNotification *repositories.ServiceNotification
}

func NewCreate_ConductividtySensor(appointmentRepo domain.IConductividtySensor, serviceNotification *repositories.ServiceNotification) *Create_ConductividtySensor {
	return &Create_ConductividtySensor{
		appointmentRepo:     appointmentRepo,
		serviceNotification: serviceNotification,
	}
}

func (c *Create_ConductividtySensor) Execute(appointment entities.ConductivitySensor) (entities.ConductivitySensor, error) {
	created, err := c.appointmentRepo.Save(appointment)
	if err != nil {
		log.Printf("Error saving data: %v", err)
		return entities.ConductivitySensor{}, err
	}
	err = c.serviceNotification.NotifyAppointmentCreated(created)
	if err != nil {
		log.Printf("Error notifying about created appointment: %v", err)
		return entities.ConductivitySensor{}, err
	}

	return created, nil
}
