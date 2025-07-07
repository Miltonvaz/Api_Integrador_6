package use_case

import (
	"Integrador/src/sensor_densidad_o/domain"
	"Integrador/src/sensor_densidad_o/domain/entities"
)

type Get_All_DensitySensor struct {
	db domain.IDensitySensor
}

func NewGet_All_DensitySensor(db domain.IDensitySensor) *Get_All_DensitySensor {
	return &Get_All_DensitySensor{db: db}
}

func (gt *Get_All_DensitySensor) Execute(userID int) ([]entities.DensitySensor, error) {
	return gt.db.GetAll(userID)
}
