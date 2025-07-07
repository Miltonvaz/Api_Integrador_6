package use_case

import "Integrador/src/motor/domain"

type Delete_Motor struct {
	db domain.MotorFermentadorRepository
}

func NewDelete_Motor(db domain.MotorFermentadorRepository) *Delete_Motor {
	return &Delete_Motor{db: db}
}

func (dt *Delete_Motor) Execute(id int, userID int) error {
	return dt.db.Delete(id, userID)
}
