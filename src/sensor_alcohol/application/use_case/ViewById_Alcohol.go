package use_case

import "Integrador/src/sensor_alcohol/domain"

type GetByID_AlcoholSensor struct {
	db domain.IAlcoholSensor
}

func NewGetByID_AlcoholSensor(db domain.IAlcoholSensor) *GetByID_AlcoholSensor {
	return &GetByID_AlcoholSensor{db: db}
}

func (u *GetByID_AlcoholSensor) Execute(id, userID int) (interface{}, error) {
	return u.db.GetByID(id, userID)
}
