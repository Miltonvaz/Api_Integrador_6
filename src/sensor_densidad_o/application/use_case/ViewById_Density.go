package use_case

import "Integrador/src/sensor_densidad_o/domain"

type GetByID_DensitySensor struct {
	db domain.IDensitySensor
}

func NewGetByID_DensitySensor(db domain.IDensitySensor) *GetByID_DensitySensor {
	return &GetByID_DensitySensor{db: db}
}

func (u *GetByID_DensitySensor) Execute(id, userID int) (interface{}, error) {
	return u.db.GetByID(id, userID)
}
