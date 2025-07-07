package domain

import "Integrador/src/sensor_turbuidez/domain/entities"

type TurbiditySensor interface {
	Save(sensor entities.TurbiditySensor) (entities.TurbiditySensor, error)
	GetByID(id, userID int) (entities.TurbiditySensor, error)
	GetAll(userID int) ([]entities.TurbiditySensor, error)
	Delete(id, userID int) error
}
