package use_case

import "Integrador/src/sensor_alcohol/domain"

type Delete_AlcoholSensor struct {
	db domain.IAlcoholSensor
}

func NewDelete_AlcoholSensor(db domain.IAlcoholSensor) *Delete_AlcoholSensor {
	return &Delete_AlcoholSensor{db: db}
}

func (dt *Delete_AlcoholSensor) Execute(id int, userID int) error {
	return dt.db.Delete(id, userID)
}
