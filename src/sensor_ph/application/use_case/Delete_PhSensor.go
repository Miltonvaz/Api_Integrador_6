package use_case

import "Integrador/src/sensor_ph/domain"

type Delete_PhSensor struct {
	db domain.IPhSensor
}

func NewDelete_PhSensor(db domain.IPhSensor) *Delete_PhSensor {
	return &Delete_PhSensor{db: db}
}

func (dt *Delete_PhSensor) Execute(id int, userID int) error {
	return dt.db.Delete(id, userID)
}
