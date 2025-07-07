package use_case

import (
	"Integrador/src/sensor_alcohol/application/repositories"
	"Integrador/src/sensor_alcohol/domain"
	"Integrador/src/sensor_alcohol/domain/entities"
	"log"
)

type Create_AlcoholSensor struct {
	appointmentRepo     domain.IAlcoholSensor
	serviceNotification *repositories.ServiceNotification
}

func NewCreate_AlcoholSensor(appointmentRepo domain.IAlcoholSensor, serviceNotification *repositories.ServiceNotification) *Create_AlcoholSensor {
	return &Create_AlcoholSensor{
		appointmentRepo:     appointmentRepo,
		serviceNotification: serviceNotification,
	}
}

func (c *Create_AlcoholSensor) Execute(appointment entities.AlcoholSensor) (entities.AlcoholSensor, error) {
	created, err := c.appointmentRepo.Save(appointment)
	if err != nil {
		log.Printf("Error saving data: %v", err)
		return entities.AlcoholSensor{}, err
	}
	err = c.serviceNotification.NotifyAppointmentCreated(created)
	if err != nil {
		log.Printf("Error notifying about created appointment: %v", err)
		return entities.AlcoholSensor{}, err
	}

	return created, nil
}
