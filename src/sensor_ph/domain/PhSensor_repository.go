package domain

import entities "Integrador/src/sensor_ph/domain/entities"

type IPhSensor interface {
	Save(sensor entities.PhSensor) (entities.PhSensor, error)
	GetByID(id, userID int) (entities.PhSensor, error)
	GetAll(userID int) ([]entities.PhSensor, error)
	Delete(id, userID int) error
}
