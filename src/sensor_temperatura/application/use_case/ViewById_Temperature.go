package use_case

import "Integrador/src/sensor_temperatura/domain"

type GetByID_TemperatureSensor struct {
	db domain.ITemperatureSensor
}

func NewGetByID_TemperatureSensor(db domain.ITemperatureSensor) *GetByID_TemperatureSensor {
	return &GetByID_TemperatureSensor{db: db}
}

func (u *GetByID_TemperatureSensor) Execute(id, userID int) (interface{}, error) {
	return u.db.GetByID(id, userID)
}
