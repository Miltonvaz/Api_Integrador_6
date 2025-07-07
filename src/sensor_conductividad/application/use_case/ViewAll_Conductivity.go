package use_case

import (
	"Integrador/src/sensor_conductividad/domain"
	"Integrador/src/sensor_conductividad/domain/entities"
)

type Get_All_ConductividtySensor struct {
	db domain.IConductividtySensor
}

func NewGet_All_ConductividtySensor(db domain.IConductividtySensor) *Get_All_ConductividtySensor {
	return &Get_All_ConductividtySensor{db: db}
}

func (gt *Get_All_ConductividtySensor) Execute(userID int) ([]entities.ConductivitySensor, error) {
	return gt.db.GetAll(userID)
}
