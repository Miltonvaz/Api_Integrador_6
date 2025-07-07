package use_case

import (
	"Integrador/src/sensor_turbuidez/domain"
	"Integrador/src/sensor_turbuidez/domain/entities"
)

type Get_All_TurbiditySensor struct {
	db domain.TurbiditySensor
}

func NewGet_All_TurbiditySensor(db domain.TurbiditySensor) *Get_All_TurbiditySensor {
	return &Get_All_TurbiditySensor{db: db}
}

func (gt *Get_All_TurbiditySensor) Execute(userID int) ([]entities.TurbiditySensor, error) {
	return gt.db.GetAll(userID)
}
