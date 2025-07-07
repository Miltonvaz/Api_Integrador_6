package use_case

import (
	"Integrador/src/sensor_alcohol/domain"
	"Integrador/src/sensor_alcohol/domain/entities"
)

type Get_All_AlcoholSensor struct {
	db domain.IAlcoholSensor
}

func NewGet_All_AlcoholSensor(db domain.IAlcoholSensor) *Get_All_AlcoholSensor {
	return &Get_All_AlcoholSensor{db: db}
}

func (gt *Get_All_AlcoholSensor) Execute(userID int) ([]entities.AlcoholSensor, error) {
	return gt.db.GetAll(userID)
}
