package domain

import "Integrador/src/sensor_densidad_o/domain/entities"

type IDensitySensor interface {
	Save(sensor entities.DensitySensor) (entities.DensitySensor, error)
	GetByID(id, userID int) (entities.DensitySensor, error)
	GetAll(userID int) ([]entities.DensitySensor, error)
	Delete(id, userID int) error
}
