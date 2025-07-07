package use_case

import "Integrador/src/sensor_conductividad/domain"

type GetByID_ConductividtySensor struct {
	db domain.IConductividtySensor
}

func NewGetByID_ConductividtySensor(db domain.IConductividtySensor) *GetByID_ConductividtySensor {
	return &GetByID_ConductividtySensor{db: db}
}

func (u *GetByID_ConductividtySensor) Execute(id, userID int) (interface{}, error) {
	return u.db.GetByID(id, userID)
}
