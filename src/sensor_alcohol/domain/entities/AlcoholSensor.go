package entities

import "time"

type AlcoholSensor struct {
	MeasurementID        int       `json:"measurement_id"`
	UserID               int       `json:"user_id"`
	Timestamp            time.Time `json:"timestamp"`
	AlcoholConcentration float64   `json:"alcohol_concentration"`
}
