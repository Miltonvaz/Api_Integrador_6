package use_case

import "Integrador/src/sensor_ph/domain"

type GetByID_PhSensor struct {
	db domain.IPhSensor
}

func NewGetByID_PhSensor(db domain.IPhSensor) *GetByID_PhSensor {
	return &GetByID_PhSensor{db: db}
}

func (u *GetByID_PhSensor) Execute(id, userID int) (interface{}, error) {
	return u.db.GetByID(id, userID)
}
