package use_case

import "Integrador/src/sensor_turbuidez/domain"

type GetByID_TurbiditySensor struct {
	db domain.TurbiditySensor
}

func NewGetByID_TTurbiditySensor(db domain.TurbiditySensor) *GetByID_TurbiditySensor {
	return &GetByID_TurbiditySensor{db: db}
}

func (u *GetByID_TurbiditySensor) Execute(id, userID int) (interface{}, error) {
	return u.db.GetByID(id, userID)
}
