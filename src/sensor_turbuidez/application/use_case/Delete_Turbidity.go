package use_case

import "Integrador/src/sensor_turbuidez/domain"

type Delete_TurbiditySensor struct {
	db domain.TurbiditySensor
}

func NewDelete_TurbiditySensor(db domain.TurbiditySensor) *Delete_TurbiditySensor {
	return &Delete_TurbiditySensor{db: db}
}

func (dt *Delete_TurbiditySensor) Execute(id int, userID int) error {
	return dt.db.Delete(id, userID)
}
