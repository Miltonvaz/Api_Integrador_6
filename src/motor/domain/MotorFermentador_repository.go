package domain

import "Integrador/src/motor/domain/entities"

type MotorFermentadorRepository interface {
	Save(sensor entities.MotorFermentador) (entities.MotorFermentador, error)
	GetByID(id, userID int) (entities.MotorFermentador, error)
	GetAll(userID int) ([]entities.MotorFermentador, error)
	Delete(id, userID int) error
}
