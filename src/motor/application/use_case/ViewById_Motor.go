package use_case

import "Integrador/src/motor/domain"

type GetByID_Motor struct {
	db domain.MotorFermentadorRepository
}

func NewGetByID_Motor(db domain.MotorFermentadorRepository) *GetByID_Motor {
	return &GetByID_Motor{db: db}
}

func (u *GetByID_Motor) Execute(id, userID int) (interface{}, error) {
	return u.db.GetByID(id, userID)
}
