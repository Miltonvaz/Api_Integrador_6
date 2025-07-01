package use_case

import (
	"Integrador/src/sensor_ph/domain"
	entities "Integrador/src/sensor_ph/domain/entities"
)

type Get_All_PhSensor struct {
	db domain.IPhSensor
}

func NewGet_All_PhSensor(db domain.IPhSensor) *Get_All_PhSensor {
	return &Get_All_PhSensor{db: db}
}

func (gt *Get_All_PhSensor) Execute(userID int) ([]entities.PhSensor, error) {
	return gt.db.GetAll(userID)
}
