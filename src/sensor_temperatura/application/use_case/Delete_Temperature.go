package use_case

import "Integrador/src/sensor_temperatura/domain"

type Delete_TemperatureSensor struct {
	db domain.ITemperatureSensor
}

func NewDelete_TemperatureSensor(db domain.ITemperatureSensor) *Delete_TemperatureSensor {
	return &Delete_TemperatureSensor{db: db}
}

func (dt *Delete_TemperatureSensor) Execute(id int, userID int) error {
	return dt.db.Delete(id, userID)
}
