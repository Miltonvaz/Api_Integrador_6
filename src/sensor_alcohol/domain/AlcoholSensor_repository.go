package domain

import "Integrador/src/sensor_alcohol/domain/entities"

type IAlcoholSensor interface {
	Save(sensor entities.AlcoholSensor) (entities.AlcoholSensor, error)
	GetByID(id, userID int) (entities.AlcoholSensor, error)
	GetAll(userID int) ([]entities.AlcoholSensor, error)
	Delete(id, userID int) error
}
