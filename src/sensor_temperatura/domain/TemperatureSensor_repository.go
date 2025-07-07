package domain

import "Integrador/src/sensor_temperatura/domain/entities"

type ITemperatureSensor interface {
	Save(sensor entities.TemperatureSensor) (entities.TemperatureSensor, error)
	GetByID(id, userID int) (entities.TemperatureSensor, error)
	GetAll(userID int) ([]entities.TemperatureSensor, error)
	Delete(id, userID int) error
}
