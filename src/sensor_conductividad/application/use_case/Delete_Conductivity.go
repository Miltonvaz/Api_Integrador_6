package use_case

import "Integrador/src/sensor_conductividad/domain"

type Delete_ConductividtySensor struct {
	db domain.IConductividtySensor
}

func NewDelete_ConductividtySensor(db domain.IConductividtySensor) *Delete_ConductividtySensor {
	return &Delete_ConductividtySensor{db: db}
}

func (dt *Delete_ConductividtySensor) Execute(id int, userID int) error {
	return dt.db.Delete(id, userID)
}
