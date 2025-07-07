package use_case

import (
	"Integrador/src/sensor_temperatura/domain"
	"Integrador/src/sensor_temperatura/domain/entities"
)

type Get_All_TemperatureSensor struct {
	db domain.ITemperatureSensor
}

func NewGet_All_TemperatureSensor(db domain.ITemperatureSensor) *Get_All_TemperatureSensor {
	return &Get_All_TemperatureSensor{db: db}
}

func (gt *Get_All_TemperatureSensor) Execute(userID int) ([]entities.TemperatureSensor, error) {
	return gt.db.GetAll(userID)
}
