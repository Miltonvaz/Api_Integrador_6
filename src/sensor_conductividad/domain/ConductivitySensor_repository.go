package domain

import "Integrador/src/sensor_conductividad/domain/entities"

type IConductividtySensor interface {
	Save(sensor entities.ConductivitySensor) (entities.ConductivitySensor, error)
	GetByID(id, userID int) (entities.ConductivitySensor, error)
	GetAll(userID int) ([]entities.ConductivitySensor, error)
	Delete(id, userID int) error
}
