package use_case

import "Integrador/src/sensor_densidad_o/domain"

type Delete_DensitySensor struct {
	db domain.IDensitySensor
}

func NewDelete_DensitySensor(db domain.IDensitySensor) *Delete_DensitySensor {
	return &Delete_DensitySensor{db: db}
}

func (dt *Delete_DensitySensor) Execute(id int, userID int) error {
	return dt.db.Delete(id, userID)
}
