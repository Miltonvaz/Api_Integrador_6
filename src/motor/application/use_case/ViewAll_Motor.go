package use_case

import (
	"Integrador/src/motor/domain"
	"Integrador/src/motor/domain/entities"
)

type Get_All_Motor struct {
	db domain.MotorFermentadorRepository
}

func NewGet_All_Motor(db domain.MotorFermentadorRepository) *Get_All_Motor {
	return &Get_All_Motor{db: db}
}

func (gt *Get_All_Motor) Execute(userID int) ([]entities.MotorFermentador, error) {
	return gt.db.GetAll(userID)
}
