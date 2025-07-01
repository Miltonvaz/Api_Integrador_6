package entities

import "time"

type ConductivitySensor struct {
	MeasurementID int       `json:"measurement_id"`
	UserID        int       `json:"user_id"`
	Timestamp     time.Time `json:"timestamp"`
	Conductivity  float64   `json:"conductivity"`
}
